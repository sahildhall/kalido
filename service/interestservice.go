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
		{key: 1, name: "userA", interests: []string{"cycling", "football"}},
		{key: 2, name: "userB", interests: []string{"cricket", "tennis"}},
		{key: 3, name: "userC", interests: []string{"cycling", "tennis"}},
		{key: 4, name: "userD", interests: []string{"chess", "cricket"}},
		{key: 5, name: "userE", interests: []string{"football", "handball"}},
		{key: 6, name: "userF", interests: []string{"handball", "chess"}},
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
	tempList1, tempList2 := []string{}, []string{}

	// time complexity (worst case, best case) for searching : O(n), O(1)
	for i := range data {
		if data[i].key == user1 {
			tempList1 = append(tempList1, data[i].interests...)
		} else if data[i].key == user2 {
			tempList2 = append(tempList2, data[i].interests...)
		} else {
			continue
		}
	}

	// time complexity to find common interests : O(n1 + n2)
	for i := range tempList1 {
		for j := range tempList2 {
			if string(tempList1[i]) == string(tempList2[j]) {
				commonInterests = append(commonInterests, string(tempList1[i]))
			}
		}
	}
	if len(commonInterests) == 0 {
		return &api.Interests{Interests: noInterests}, nil
	}
	return &api.Interests{Interests: commonInterests}, nil
}
