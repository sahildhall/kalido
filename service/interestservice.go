package service

import (
	"context"
	"fmt"
	"kalido/api"
)

// mock database schema structure
type mockuser struct {
	key       int64
	name      string
	interests []string
}

func initInterestData() []mockuser {
	// simulated : user database
	users := []mockuser{
		{key: 1, name: "Sahil", interests: []string{"cycling", "football"}},
		{key: 2, name: "Vijay", interests: []string{"cricket", "tennis"}},
		{key: 3, name: "Abhishek", interests: []string{"cycling", "tennis"}},
		{key: 4, name: "Jacob", interests: []string{"chess", "cricket"}},
		{key: 5, name: "John", interests: []string{"football", "handball"}},
		{key: 6, name: "Pankaj", interests: []string{"handball", "chess"}},
	}
	return users
}

// this function takes in the key and the data to return the user instance from the data
func searchForUser(key int64, data []mockuser) (*mockuser, error) {
	for i := range data {
		if data[i].key == key {
			return &data[i], nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// rpc method definition
func (s *Service) GetSharedInterests(ctx context.Context, twoKeys *api.TwoUserKeys) (*api.Interests, error) {
	// inits
	user1 := twoKeys.GetUser1().Key
	user2 := twoKeys.GetUser2().Key
	data := initInterestData()
	commonInterests := []string{}
	noInterests := []string{"no common interests"}

	// making a temporary interest list of both the users
	user1Interests, user2Interests := []string{}, []string{}

	// Getting interests of requested user and user2
	for i := range data {
		if data[i].key == user1 {
			user1Interests = append(user1Interests, data[i].interests...)
		} else if data[i].key == user2 {
			user2Interests = append(user2Interests, data[i].interests...)
		} else {
			continue
		}
	}

	// Comparing both users interests
	for i := range user1Interests {
		for j := range user2Interests {
			if string(user1Interests[i]) == string(user2Interests[j]) {
				commonInterests = append(commonInterests, string(user1Interests[i]))
			}
		}
	}
	if len(commonInterests) == 0 {
		return &api.Interests{Interests: noInterests}, nil
	}
	return &api.Interests{Interests: commonInterests}, nil
}
