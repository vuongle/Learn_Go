package interceptors_server

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("---> Unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func StreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	log.Println("---> Stream interceptor: ", info.FullMethod)
	return handler(srv, ss)
}
