package service

import (
	"context"
	"fmt"
	"kalido/api"
)

type network struct {
	networkKey int64
	keys       []*api.UserKey
}

// initialize database
func initializeMockDatabase() []network {
	// this is a simulated schema of user-network relationship
	// the idea is to store every corresponding array of member keys with their network key
	data := []network{
		{
			networkKey: 100,
			keys: []*api.UserKey{
				{Key: 1},
				{Key: 2},
				{Key: 3}},
		},
		{
			networkKey: 101,
			keys: []*api.UserKey{
				{Key: 4},
				{Key: 5},
				{Key: 6}},
		},
	}
	return data
}

// rpc method definition
func (s *Service) GetNetworkMembers(ctx context.Context, networkKey *api.NetworkKey) (*api.UserKeys, error) {

	// mockedNetworkData
	mockedNetworkData := initializeMockDatabase()
	for i := range mockedNetworkData {
		network := mockedNetworkData[i]

		// Check requested network in database, if found return keys (Array of userIds) else continue
		if network.networkKey == networkKey.Key {
			return &api.UserKeys{Users: network.keys}, nil
		} else {
			continue
		}
	}
	// if nothing found return with an error
	return nil, fmt.Errorf("network not found\n network key %d maybe wrong", networkKey.Key)
}
