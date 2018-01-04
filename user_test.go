package model

import "testing"
import "reflect"

func TestNewUser(t *testing.T) {
	u := User{
		ID:        1,
		URL:       "www.ribice.ba",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@mail.com",
		Username:  "user",
	}
	if !reflect.DeepEqual(u, NewUser(1, "user")) {
		t.Error("expected equal structs")
	}
}

func TestNewUsers(t *testing.T) {
	u := []User{
		{
			ID:        1,
			URL:       "www.ribice.ba",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@mail.com",
			Username:  "user1",
		},
		{
			ID:        2,
			URL:       "www.ribice.ba",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@mail.com",
			Username:  "user2",
		},
	}
	if !reflect.DeepEqual(u, NewUsers("user1", "user2")) {
		t.Error("expected equal structs")
	}
}
