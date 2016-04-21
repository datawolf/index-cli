//
// repos.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package index

import (
	"bytes"
	"fmt"

	qs "github.com/google/go-querystring/query"
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

type RepoDesc struct {
	Description *string `json:"description,omitempty"`
}

func (s *RepositoriesService) GetRepoDesc(repo string) (*RepoDesc, *Response, error) {
	result := new(RepoDesc)
	u := fmt.Sprintf("/index/repositories/%s/description", repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return result, nil, err
	}
	resp, err := s.client.Do(req, result)
	return result, resp, err
}

func (s *RepositoriesService) SetRepoDesc(repo string, repoDesc *RepoDesc) (string, *Response, error) {
	u := fmt.Sprintf("/index/repositories/%s/description", repo)
	req, err := s.client.NewRequest("PUT", u, repoDesc)
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

func (s *RepositoriesService) DeleteRepo(repo string) (string, *Response, error) {
	u := fmt.Sprintf("/index/repositories/%s/entirety", repo)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return "", nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(req, buf)
	if err != nil {
		return "", resp, err
	}

	return buf.String(), resp, nil
}

func (s *RepositoriesService) DeleteTag(repo string, tag string) (string, *Response, error) {
	params, err := qs.Values(nil)
	if err != nil {
		return "", nil, err
	}
	params.Add("tag", tag)
	u := fmt.Sprintf("/index/repositories/%s/tag?%s", repo, params.Encode())
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return "", nil, err
	}

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(req, buf)
	if err != nil {
		return "", resp, err
	}

	return buf.String(), resp, nil
}

type UserRepo struct {
	NumberDL    *int    `json:"download_num,omitempty"`
	NumberImage *int    `json:"image_num,omitempty"`
	Property    *string `json:"property,omitempty"`
	RepoName    *string `json:"repo,omitempty"`
	Size        *int    `json:"size,omitempty"`
}

type UserRepoResult struct {
	RepoList []UserRepo `json:"repo_list,omitempty"`
}

func (u UserRepoResult) String() string {
	return Stringify(u)
}

func (u UserRepo) String() string {
	return Stringify(u)
}

func (s *RepositoriesService) GetUserRepo() (*UserRepoResult, *Response, error) {
	result := new(UserRepoResult)
	req, err := s.client.NewRequest("GET", "index/userrepo", nil)
	if err != nil {
		return result, nil, err
	}
	resp, err := s.client.Do(req, result)
	return result, resp, err
}
