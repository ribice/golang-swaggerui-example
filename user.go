package model

import "time"

// User represents the object of individual and member of organization.
type User struct {
	ID        int    `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"-"`

	LastLogin time.Time `json:"last_login"`

	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	DeletedOn time.Time `json:"deleted_on"`
}

// NewUser instantiates a new user
func NewUser(id int, userName string) User {
	return User{
		ID:        id,
		URL:       "www.ribice.ba",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@mail.com",
		Username:  userName,
	}
}

// NewUsers instanties slice of new user models
func NewUsers(userNames ...string) []User {

	var u []User
	for i, v := range userNames {
		u = append(u, NewUser(i+1, v))
	}
	return u
}
