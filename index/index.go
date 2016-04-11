//
// index.go
// Copyright (C) 2016 datawolf <datawolf@datawolf-Lenovo-G460>
//
// Distributed under terms of the MIT license.
//

package index

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	libraryVersion = "0.0.1"
	defaultBaseURL = "http://rnd-dockerhub.huawei.com"
	userAgent      = "rnd-dockerhub/" + libraryVersion
)

// A Client manages communication with the rnd-dockerhub API.
type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client
	// Bse URL for API request.
	BaseURL *url.URL
	// User agent used when communicating with the API
	UserAgent string

	// Servcie used for talking to different parts of the rnd-dockerhub API
	Repositories  *RepositoriesService
	Organizations *OrganizationsService
	Search        *SearchService
}

// addOptions adds the parameters in opt as URL query  parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, nil
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns a new rnd-dockerhub  API client. if a nil httpclient is provided,
// http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}

	c.Repositories = &RepositoriesService{client: c}
	c.Organizations = &OrganizationsService{client: c}
	c.Search = &SearchService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "xxxx")
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Response is a rnd-dockerhub API response. This wraps  the standard http.Response
// returnef from rnd-dockerhub.
type Response struct {
	*http.Response
}

func newResponse(res *http.Response) *Response {
	response := &Response{Response: res}

	return response
}

// CheckResponse check the API response for errors
// A response is considered an error if it has astatus code outside the 200 range
func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return errors.New("error")
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		//  Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		// even thougth  there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, nil
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore  EOF errors caused by  empty response body
			}
		}
	}

	return response, err
}

// BasicAuthTransport is an http.RoundTripper that authenticates all requests
// using HTTP Basic Authentication with the provided username and password.
type BasicAuthTransport struct {
	Username string
	Password string

	// Transport is the underlying HTTP transport to use when making requests.
	//It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req)
	req.SetBasicAuth(t.Username, t.Password)

	return t.transport().RoundTrip(req)
}

// Client returns an *http.Client that makes request that are authenticated
// using HTTP Basic Authentication
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BasicAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request. The clone is a
// shadow copy of the struct and its Header map
func cloneRequest(r *http.Request) *http.Request {
	// shadow copy of the struct
	r2 := new(http.Request)
	*r2 = *r

	// deep coy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, v := range r.Header {
		r2.Header[k] = append([]string(nil), v...)
	}

	return r2
}
