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
		{key: 1, name: "user1"},
		{key: 2, name: "user2"},
		{key: 3, name: "user3"},
		{key: 4, name: "user4"},
		{key: 5, name: "user5"},
		{key: 6, name: "user6"},
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
