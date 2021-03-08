package service

import (
	"context"
	"fmt"
	"kalido/api"
)

// mock database schema structure
type userData struct {
	key  int64
	name string
}

// function to initialize the mocked database
func initMockUserData() []userData {
	users := []userData{
		{key: 1, name: "Sahil"},
		{key: 2, name: "Vijay"},
		{key: 3, name: "Abhishek"},
		{key: 4, name: "Jacob"},
		{key: 5, name: "John"},
		{key: 6, name: "Pankaj"},
	}

	return users
}

// utility funtion that gets user record provided the unique user key and the data
func getUserRecord(userKey int64, data []userData) (*userData, error) {
	for i := range data {
		if data[i].key == userKey {
			return &data[i], nil
		}
	}
	return nil, fmt.Errorf("did not find user key: %d, in the database", userKey)
}

// rpc method definition
func (s *Service) GetUser(ctx context.Context, userKey *api.UserKey) (*api.User, error) {
	mockedUserData := initMockUserData()
	userResult, err := getUserRecord(userKey.Key, mockedUserData)

	// If not found return error
	if err != nil {
		return nil, err
	}

	// If found return username and id
	return &api.User{Key: userResult.key, Name: userResult.name}, nil
}
