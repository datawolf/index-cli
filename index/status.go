//
// status.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package index

import (
	"bytes"
)

type StatusService struct {
	client *Client
}

// Get get the status info of the rnd-dockerhub
func (s *StatusService) Get() (string, *Response, error) {
	req, err := s.client.NewRequest("GET", "index/_ping", nil)
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
