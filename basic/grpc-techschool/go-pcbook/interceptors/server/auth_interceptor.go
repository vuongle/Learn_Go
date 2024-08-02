package interceptors_server

import (
	"context"
	"log"

	"github.com/vuongle/grpc/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Define an auth interceptor for server side
type AuthInterceptor struct {
	jwtManager      *jwt.JWTManager
	accessibleRoles map[string][]string //key: grpc method, value: a list of roles
}

func NewAuthInterceptor(
	jwtManager *jwt.JWTManager,
	accessibleRoles map[string][]string) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager:      jwtManager,
		accessibleRoles: accessibleRoles,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		log.Println("---> Unary interceptor: ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		log.Println("---> Stream interceptor: ", info.FullMethod)

		err := interceptor.authorize(ss.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, ss)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {

	// get role by grpc method
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		// if not found -> means that this grpc method can access by anyone
		return nil
	}

	// get metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	value := md["authorization"]
	if len(value) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// get token from metadata and verify it
	accessToken := value[0]
	userClaims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	// check role from token
	for _, role := range accessibleRoles {
		if role == userClaims.Role {

			// access granted
			return nil
		}
	}

	return status.Errorf(codes.PermissionDenied, "no permission to access this RPC")
}
