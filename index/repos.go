//
// repos.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package index

import (
	"fmt"
)

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

type Image struct {
	Tag  *string `json:"tag,omitempty"`
	Size *int    `json:"size,omitempty"`
}

func (i Image) String() string {
	return Stringify(i)
}

// Property represents a rund-dockerhub repo's property
type Property struct {
	NumberDL    *int    `json:"download_num,omitempty"`
	ImageList   []Image `json:"image_list,omitempty"`
	NumberImage *int    `json:"image_num,omitempty"`
	Property    *string `json:"property,omitempty"`
	RepoName    *string `json:"repo,omitempty"`
	Size        *int    `json:"size,omitempty"`
}

func (s *RepositoriesService) Get(repo string) (*Property, *Response, error) {
	result := new(Property)
	u := fmt.Sprintf("/index/repositories/%s/properties", repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return result, nil, err
	}
	resp, err := s.client.Do(req, result)
	return result, resp, err
}

func (s *RepositoriesService) Set(repo string, property *Property) (string, *Response, error) {
	u := fmt.Sprintf("/index/repositories/%s/properties", repo)
	req, err := s.client.NewRequest("PUT", u, property)
	if err != nil {
		return "", nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return "", nil, err
	}

	return "SUCCESS", resp, nil
}
