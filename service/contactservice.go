package service

import (
	"context"
	"fmt"
	"kalido/api"
)

// this struct implements all the service methods
type Service struct {
	api.UnimplementedUserServiceServer
	api.UnimplementedContactServiceServer
	api.UnimplementedInterestsServiceServer
	api.UnimplementedNetworkServiceServer
}

// basic user struct used as mock database schema
type user struct {
	key      int64
	name     string
	contacts []int64
}

// function to initialize the mocked contact database
func initMockContactsData() []user {
	users := []user{
		{key: 1, name: "Sahil", contacts: []int64{2, 3}},
		{key: 2, name: "Vijay", contacts: []int64{1, 3}},
		{key: 3, name: "Abhishek", contacts: []int64{4, 1}},
		{key: 4, name: "Jacob", contacts: []int64{5, 6}},
		{key: 5, name: "John", contacts: []int64{6, 1}},
		{key: 6, name: "Pankaj", contacts: []int64{3, 4}},
	}

	return users
}

// this function takes in the key and the data to return the user instance from the data
func search(key int64, data []user) (*user, error) {
	for i := range data {
		if data[i].key == key {
			return &data[i], nil
		}
	}
	return nil, fmt.Errorf("")
}

// rpc method definition
func (s *Service) GetCommonContacts(ctx context.Context, keys *api.TwoUserKeys) (*api.Contacts, error) {
	// initializing data
	mockedContactData := initMockContactsData()
	contactList := []string{}

	// getting data from request payload
	user1Key := keys.GetUser1()
	user2Key := keys.GetUser2()

	// making a contact list of both the users
	user1Contacts, user2Contacts := []int64{}, []int64{}

	// finding user1 and user2 contacts from mockedContactData and pushing them in tempArray
	for i := range mockedContactData {
		user := mockedContactData[i]
		if user.key == user1Key.Key {
			user1Contacts = append(user1Contacts, user.contacts...)
		} else if user.key == user2Key.Key {
			user2Contacts = append(user2Contacts, user.contacts...)
		} else {
			continue
		}
	}

	// Serach user2Contacts in user1Contacts(Requested user)
	for i := range user1Contacts {
		for j := range user2Contacts {
			if user1Contacts[i] == user2Contacts[j] {
				contact, err := search(user2Contacts[j], mockedContactData)
				if err != nil {
					return nil, fmt.Errorf("failed to get contact")
				}
				contactList = append(contactList, contact.name)
			}
		}
	}
	return &api.Contacts{Contacts: contactList}, nil
}
