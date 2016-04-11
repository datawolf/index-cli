//
// search.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package index

import (
	"fmt"
	qs "github.com/google/go-querystring/query"
)

// SearchService provides access to the search related functions in the rnd-dockerhub API.
//
// rnd-dockerhub API docks: http://code.huawei.com/docker-incubator/index/blob/master/docs/api.md
type SearchService struct {
	client *Client
}

// SearchOptions specties optional parameters to the SearchService methods
type SearchOptions struct {
	// How to sort the search results.
	Sort string `url:"sort,omitempty"`
	// Sort order if sort parameter is provided. Possible values are asc, desc
	// Default is desc
	Order string `url:"order,omitempty"`
}

// RepositoriesSearchResult represents the result of a repositories search
type RepositoriesSearchResult struct {
	NumberResults *int         `json:"num_results,omitempty"`
	QueryString   *string      `json:"query,omitempty"`
	Repositories  []Repository `json:"results,omitempty"`
}

// Repositories searches repositories via various criteria.
func (s *SearchService) Repositories(query string, opt *SearchOptions) (*RepositoriesSearchResult, *Response, error) {
	result := new(RepositoriesSearchResult)
	resp, err := s.search(query, opt, result)
	return result, resp, err
}

func (s *SearchService) search(query string, opt *SearchOptions, result interface{}) (*Response, error) {
	params, err := qs.Values(opt)
	if err != nil {
		return nil, err
	}
	params.Add("q", query)
	u := fmt.Sprintf("/v1/search?%s", params.Encode())
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, result)
}
