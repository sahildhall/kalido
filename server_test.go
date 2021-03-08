package main

import (
	"context"
	"fmt"
	"kalido/api"
	"kalido/endpoint"
	"kalido/service"
	"log"
	"testing"

	"google.golang.org/grpc"
)

// tests
func TestViewNetworkService(t *testing.T) {
	fmt.Printf("\n\nUnit test case for View Network Service:====\n\n")

	grpcConn, err := grpc.Dial(":3031", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()

	networkClient := api.NewNetworkServiceClient(grpcConn)
	userClient := api.NewUserServiceClient(grpcConn)
	contactClient := api.NewContactServiceClient(grpcConn)
	interestsClient := api.NewInterestsServiceClient(grpcConn)

	s := endpoint.NewServer(contactClient, interestsClient, userClient, networkClient)
	ctx := context.Background()

	userK := &api.UserKey{Key: 1}
	networkkey := &api.NetworkKey{Key: 101}

	netPayload := &api.UserViewingNetwork{
		User:    userK,
		Network: networkkey,
	}

	result, err := s.ViewNetworkMembers(ctx, netPayload)
	if err != nil {
		log.Print("failed to return members")
		t.Errorf("ViewNetworkService failed :\n%v", err)

	} else {
		log.Printf("ViewNetworkService passed successfully \n result:\t%v", result)
	}
}

func TestNetworkService(t *testing.T) {
	s := service.Service{}
	ctx := context.Background()

	networkkey := &api.NetworkKey{Key: 100}
	result, error := s.GetNetworkMembers(ctx, networkkey)
	if error != nil {
		t.Errorf("problem")
	}
	fmt.Println("Test For Network Service\n", result)
}

func TestUserService(t *testing.T) {
	s := service.Service{}
	ctx := context.Background()

	userKey := &api.UserKey{Key: 1}
	response, err := s.GetUser(ctx, userKey)
	if err != nil {
		t.Errorf("err")
	}
	fmt.Println("Test For User Service\n", response)
}

func TestInterestService(t *testing.T) {

	s := service.Service{}
	ctx := context.Background()

	keys := &api.TwoUserKeys{
		User1: &api.UserKey{Key: int64(0)},
		User2: &api.UserKey{Key: int64(1)},
	}
	interests, err := s.GetSharedInterests(ctx, keys)
	if err != nil {
		t.Errorf("failure retrieving interests from given user keys")
	}
	fmt.Println("Test For Interest Service\n", interests)
}

func TestContactService(t *testing.T) {
	s := service.Service{}
	ctx := context.Background()

	keys := &api.TwoUserKeys{
		User1: &api.UserKey{Key: int64(1)},
		User2: &api.UserKey{Key: int64(2)},
	}
	contacts, err := s.GetCommonContacts(ctx, keys)
	if err != nil {
		t.Errorf("failure contacts from given user keys")
	}
	fmt.Println("Test For Contact Service\n", contacts)
}
