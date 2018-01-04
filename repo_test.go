package model

import "testing"
import "reflect"

func TestNewRepository(t *testing.T) {
	r := Repository{
		ID:          1,
		Name:        "glice",
		Description: "Description of this repository",
		Website:     "www.ribice.ba/glice",
		NumStars:    20,
		NumForks:    3,
		NumIssues:   15,
		NumPulls:    2,
		Owner:       NewUser(1, "ribice"),
	}
	if !reflect.DeepEqual(r, NewRepository(1, "glice", "ribice")) {
		t.Error("expected equal structs")
	}
}

func TestNewRepositories(t *testing.T) {
	r := []Repository{
		{
			ID:          1,
			Name:        "glice",
			Description: "Description of this repository",
			Website:     "www.ribice.ba/glice",
			NumStars:    20,
			NumForks:    3,
			NumIssues:   15,
			NumPulls:    2,
			Owner:       NewUser(1, "ribice"),
		},
		{
			ID:          2,
			Name:        "kiss",
			Description: "Description of this repository",
			Website:     "www.ribice.ba/glice",
			NumStars:    20,
			NumForks:    3,
			NumIssues:   15,
			NumPulls:    2,
			Owner:       NewUser(1, "ribice"),
		},
	}
	if !reflect.DeepEqual(r, NewRepositories("ribice")) {
		t.Error("expected equal structs")
	}
}
