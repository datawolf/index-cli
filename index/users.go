//
// users.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package index

import (
	"bytes"
)

// UsersService handles communication with the user related
// method of the Rnd-dockerhub API
//
// API docs: http://code.huawei.com/h00283522/europa/blob/master/docs/api.md
type UsersService struct {
	client *Client
}

type User struct {
	Username    *string `json:"name,omitempty"`
	Password    *string `json:"pwd,omitempty"`
	NewPassword *string `json:"pwdnew,omitempty"`
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone,omitempty"`
}

// Create create user via user's credential,
func (s *UsersService) Create(user *User) (string, *Response, error) {
	req, err := s.client.NewRequest("PUT", "/v1/user/create", user)
	if err != nil {
		return "", nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(req, buf)
	if err != nil {
		return "", nil, err
	}
	return buf.String(), resp, nil
}

// Update update user via user's credential,
func (s *UsersService) Update(user *User) (string, *Response, error) {
	req, err := s.client.NewRequest("PATCH", "/v1/user/update", user)
	if err != nil {
		return "", nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	buf := new(bytes.Buffer)
	resp, err := s.client.Do(req, buf)
	if err != nil {
		return "", nil, err
	}
	return buf.String(), resp, nil

}
