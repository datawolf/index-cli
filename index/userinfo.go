//
// userinfo.go
// Copyright (C) 2016 wanglong <wanglong@laoqinren.net>
//
// Distributed under terms of the MIT license.
//

package index

type UserInfoService struct {
	client *Client
}

// UserInfo represents a rnd-dockerhub userinfo
type UserInfo struct {
	NumberImage        *int    `json:"image_num,omitempty"`
	NumberImagePrivate *int    `json:"private_image_num,omitempty"`
	NumberImageProtect *int    `json:"protect_image_num,omitempty"`
	NumberImagePublic  *int    `json:"public_image_num,omitempty"`
	Namespace          *string `json:"namespace,omitempty"`
	Product            *string `json:"product,omitempty"`
	Quote              *int64  `json:"quota,omitempty"`
	UsedSpace          *int64  `json:"used_space,omitempty"`
	Username           *string `json:"username,omitempty"`
}

// Get fetch the current user infomation
func (s *UserInfoService) Get() (*UserInfo, *Response, error) {
	result := new(UserInfo)
	req, err := s.client.NewRequest("GET", "/index/userinfo", nil)
	if err != nil {
		return result, nil, err
	}

	resp, err := s.client.Do(req, result)
	return result, resp, err
}
