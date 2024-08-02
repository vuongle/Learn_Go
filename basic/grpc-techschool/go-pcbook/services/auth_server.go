package services

import (
	"context"

	"github.com/vuongle/grpc/jwt"
	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/storages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Define a auth server API for auth service
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	UserStore  storages.UserStore
	JwtManager *jwt.JWTManager
}

func NewAuthServer(
	userStore storages.UserStore,
	jwtManager *jwt.JWTManager) *AuthServer {
	return &AuthServer{
		UserStore:  userStore,
		JwtManager: jwtManager,
	}
}

// Implement unary RPC for login
func (auth *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	// find user in store
	user, err := auth.UserStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect user name or password")

	}

	// generate access token
	token, err := auth.JwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token: %v", err)
	}

	res := &pb.LoginResponse{AccessToken: token}

	return res, nil
}
