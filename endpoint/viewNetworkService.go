package endpoint

import (
	"context"
	"fmt"
	"kalido/api"
	"log"
)

type Server struct {
	contactClient  api.ContactServiceClient
	interestClient api.InterestsServiceClient
	userClient     api.UserServiceClient
	networkClient  api.NetworkServiceClient
	api.UnimplementedViewNetworkServiceServer
}

// create a server connection
func NewServer(contactClient api.ContactServiceClient, interestClient api.InterestsServiceClient, userClient api.UserServiceClient, networkClient api.NetworkServiceClient) *Server {
	return &Server{
		contactClient:  contactClient,
		interestClient: interestClient,
		userClient:     userClient,
		networkClient:  networkClient,
	}
}

// rpc method definition
func (s *Server) ViewNetworkMembers(ctx context.Context, viewNetwok *api.UserViewingNetwork) (*api.NetworkMembersView, error) {
	// initializations
	allMemberDetails := []*api.MemberDetails{}

	// getting the user key and the network key from request payload
	userkey, networkey := viewNetwok.GetUser(), viewNetwok.GetNetwork()

	// calling network service to get network members
	networkMembersKeys, err := s.networkClient.GetNetworkMembers(ctx, networkey)

	// terminating in case of error, logging the error
	if err != nil {
		fmt.Println("viewnetworkserice.go:\t")
		log.Print("Invalid network, not found in DB\n")
		return nil, err
	}

	// getting the user data from the request payload by calling user service
	user, err := s.userClient.GetUser(ctx, userkey)

	// terminating in case of error, logging the error
	if err != nil {
		log.Print("didn't get result from user service\n")
		return nil, err
	}

	// getting common contacts between the requested user and every other user that are members of the same network
	for i := range networkMembersKeys.Users {

		keys := &api.TwoUserKeys{
			User1: &api.UserKey{Key: user.Key},
			User2: &api.UserKey{Key: networkMembersKeys.Users[i].Key},
		}

		// calling other services to enrich this service with data from the data received by other services
		contacts, contactServiceError := s.contactClient.GetCommonContacts(ctx, keys)

		userData, userServiceError := s.userClient.GetUser(ctx, networkMembersKeys.Users[i])
		commonInterests, interestServiceError := s.interestClient.GetSharedInterests(ctx, keys)

		// panicking if any of the services fail
		if contactServiceError != nil || userServiceError != nil || interestServiceError != nil {
			fmt.Println("one of the services has failed: panicking")
			log.Print("panicked")
			panic(err)
		} else {
			// build the response payload
			member := &api.MemberDetails{
				User:            userData,
				CommonContacts:  contacts,
				CommonInterests: commonInterests,
			}
			// creating the final resultant response payload
			allMemberDetails = append(allMemberDetails, member)
		}
	}
	return &api.NetworkMembersView{Members: allMemberDetails}, nil
}
