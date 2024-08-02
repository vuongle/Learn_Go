package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vuongle/grpc/helpers"
	interceptors_server "github.com/vuongle/grpc/interceptors/server"
	"github.com/vuongle/grpc/jwt"
	"github.com/vuongle/grpc/pb"
	"github.com/vuongle/grpc/services"
	"github.com/vuongle/grpc/storages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// get port from command line
	port := flag.Int("port", 0, "the server port")
	serverType := flag.String("type", "grpc", "the server type")
	flag.Parse()

	// set up laptop service:
	// create laptop service's server API and its dependencies ( 3 storages )
	laptopStore := storages.NewInMemoryLaptopStore()
	imageStore := storages.NewDiskImageStore("img")
	ratingStore := storages.NewInMemoryRatingStore()
	laptopServer := services.NewLaptopServer(laptopStore, imageStore, ratingStore)

	// set up auth service:
	// create auth service's server API and its dependencies
	// seed some sample users
	userStore := storages.NewInMemoryUserStore()
	jwtManager := jwt.NewJWTManager(helpers.SecretKey, helpers.TokenDuration)
	authServer := services.NewAuthServer(userStore, jwtManager)

	err := services.SeedUsers(userStore)
	if err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	if *serverType == "grpc" {
		err = runGRPCServer(authServer, laptopServer, jwtManager, listener)
	} else {
		err = runRESTServer(authServer, laptopServer, listener)
	}

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

func runGRPCServer(
	authServer pb.AuthServiceServer,
	laptopServer pb.LaptopServiceServer,
	jwtManager *jwt.JWTManager,
	listener net.Listener,
) error {

	// create grpc server, its interceptors and register grpc server APIs
	interceptor := interceptors_server.NewAuthInterceptor(jwtManager, helpers.AccessibleRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	reflection.Register(grpcServer)

	// start grpc server
	log.Printf("Start GRPC server on %s", listener.Addr().String())
	return grpcServer.Serve(listener)
}

func runRESTServer(
	authServer pb.AuthServiceServer,
	laptopServer pb.LaptopServiceServer,
	listener net.Listener,
) error {
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// register http handler
	err := pb.RegisterAuthServiceHandlerServer(ctx, mux, authServer)
	if err != nil {
		return err
	}
	err = pb.RegisterLaptopServiceHandlerServer(ctx, mux, laptopServer)
	if err != nil {
		return err
	}

	log.Printf("Start REST server on %s", listener.Addr().String())
	return http.Serve(listener, mux)
}
