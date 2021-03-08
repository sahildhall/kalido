package main

import (
	"fmt"
	"net"

	"kalido/api"
	"kalido/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// listen
	listener, err := net.Listen("tcp", ":3031")
	if err != nil {
		panic(err)
	}

	serverObject := grpc.NewServer()
	fmt.Println("server")
	// register services
	api.RegisterContactServiceServer(serverObject, &service.Service{})
	api.RegisterUserServiceServer(serverObject, &service.Service{})
	api.RegisterInterestsServiceServer(serverObject, &service.Service{})
	api.RegisterNetworkServiceServer(serverObject, &service.Service{})

	reflection.Register(serverObject)

	// Serve
	if e := serverObject.Serve(listener); e != nil {
		panic(e)
	}
}
