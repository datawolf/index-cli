//
// repos.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package index

// RepositoriesService handles communication with the repository related
// methods of the rnd-dockerhub API
type RepositoriesService struct {
	client *Client
}

// Repository represents a rmd-dockerhub repository
type Repository struct {
	Description *string `json:"description,omitempty"`
	IsOfficial  *bool   `json:"is_official,omitempty"`
	IsTrusted   *bool   `json:"is_trusted,omitempty"`
	Name        *string `json:"name,omitempty"`
	StarCount   *int    `json:"star_count,omitempty"`
}

func (r Repository) String() string {
	return Stringify(r)
}
