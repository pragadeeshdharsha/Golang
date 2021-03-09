package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int    `json : Id`
	FirstName string `json : First Name`
	LastName  string `json : Last Name`
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUsers(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User should not have an ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {

	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, errors.New("User not found")
}

func UpdateUser(u User) (User, error) {
	for i, us := range users {
		if us.ID == u.ID {
			users[i] = &u
			return *users[i], nil
		}
	}

	return User{}, errors.New("User not found for updation")
}

func RemoveUser(id int) error {

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID %v not found", id)
}
