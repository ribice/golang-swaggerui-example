package model

import (
	"time"
)

// Repository represents a git repository.
type Repository struct {
	ID          int    `json:"id,omitempty"`
	Owner       User   `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Website     string `json:"website,omitempty"`

	NumStars  int `json:"num_stars,omitempty"`
	NumForks  int `json:"num_forks,omitempty"`
	NumIssues int `json:"num_issues,omitempty"`
	NumPulls  int `json:"num_pulls,omitempty"`

	IsPrivate bool `json:"is_private,omitempty"`
	IsMirror  bool `json:"is_mirror,omitempty"`
	IsFork    bool `json:"is_fork,omitempty"`

	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	DeletedOn time.Time `json:"deleted_on"`
}

// NewRepository istantiates a new repository
func NewRepository(id int, repoName, author string) Repository {
	return Repository{
		ID:          id,
		Name:        repoName,
		Description: "Description of this repository",
		Website:     "www.ribice.ba/glice",
		NumStars:    20,
		NumForks:    3,
		NumIssues:   15,
		NumPulls:    2,

		Owner: NewUser(1, author),
	}
}

// NewRepositories returns slice of repositories
func NewRepositories(author string) []Repository {
	return []Repository{NewRepository(1, "glice", author), NewRepository(2, "kiss", author)}
}
