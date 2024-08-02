package interceptors_client

import (
	"context"
	"log"
	"time"

	"github.com/vuongle/grpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Implement an auth interceptor for client side
type AuthInterceptor struct {
	authClient  *client.AuthClient
	authMethods map[string]bool
	accessToken string
}

func MewAuthInterceptor(
	authClient *client.AuthClient,
	authMethods map[string]bool,
	refreshDuration time.Duration,
) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		authClient:  authClient,
		authMethods: authMethods,
	}

	// schedule the time to refresh the token
	err := interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	// always the refresh token for the 1st time so that the token is ready to be used
	err := interceptor.refreshTokenWithoutSchedule()
	if err != nil {
		return err
	}

	// start a goroutine to refresh the token based on the duration
	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)

			err := interceptor.refreshTokenWithoutSchedule()
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()

	return nil
}

func (interceptor *AuthInterceptor) refreshTokenWithoutSchedule() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}

	// update the new token
	interceptor.accessToken = accessToken

	log.Printf("token refreshed: %v", accessToken)
	return nil
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		log.Printf("Unary client interceptor: %s", method)

		// check if the grpc method needs authentication
		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption) (grpc.ClientStream, error) {
		log.Printf("Stream client interceptor: %s", method)

		// check if the grpc method needs authentication
		if interceptor.authMethods[method] {
			return streamer(interceptor.attachToken(ctx), desc, cc, method, opts...)
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	// add the token to the context
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}
