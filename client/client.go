package client

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/insighted4/siconv/siconv"
	"github.com/insighted4/siconv/version"
	"gopkg.in/resty.v1"
)

const Prefix = "/api/v1/"

var (
	UserAgent   = fmt.Sprintf("SICONV API/Client %s", version.Version)
	ErrNotFound = errors.New(http.StatusText(http.StatusNotFound))
)

type Client struct {
	client *resty.Client
	prefix string
	token  string
}

var _ siconv.Service = (*Client)(nil)

func (s *Client) get(result interface{}, url string, params map[string]string) (int, error) {
	var apiError map[string]interface{}
	res, err := s.client.R().
		SetResult(result).
		SetHeader("Content-Type", "application/json").
		SetQueryParams(params).
		SetError(&apiError).
		Get(url)

	if err != nil {
		return 0, err
	}

	if apiError != nil {
		return 0, fmt.Errorf("unable get %s: %v", url, apiError["message"])
	}

	if res.StatusCode() == http.StatusNotFound {
		return 0, ErrNotFound
	}

	total := 0
	return total, nil

}

func (s *Client) post(body interface{}, url string) (string, error) {
	var apiError map[string]interface{}
	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetError(&apiError).
		SetBody(body).
		Post(url)
	if err != nil {
		return "", err
	}

	if apiError != nil {
		return "", fmt.Errorf("unable to post %s: %v", url, apiError["message"])
	}

	location := res.RawResponse.Header["Location"]
	if location == nil || len(location) == 0 {
		return "", fmt.Errorf("unable to read id")
	}

	return location[0], nil
}

func (s *Client) put(body interface{}, url string) error {
	var apiError map[string]interface{}
	res, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetError(&apiError).
		SetBody(body).
		Put(url)
	if err != nil {
		return err
	}

	if apiError != nil {
		return fmt.Errorf("unable to put %v: %v", url, apiError["message"])
	}

	if res.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("unexpected response from server: status %d", res.StatusCode())
	}

	return nil
}

func NewRequestMiddleware(token string) func(*resty.Client, *resty.Request) error {
	return func(c *resty.Client, req *resty.Request) error {
		req.SetHeader("Accept", "application/json").
			SetHeader("User-Agent", UserAgent)

		if token != "" {
			req.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
		}

		return nil
	}
}

// New returns s new client to different Atom APIs.
func New(host string, token string) *Client {
	transport := &http.Transport{
		MaxIdleConns:        20000,
		MaxIdleConnsPerHost: 1000, // see https://github.com/golang/go/issues/13801
		DisableKeepAlives:   false,
		DisableCompression:  true,
		// 5 minutes is typically above the maximum sane scrape interval. So we can
		// use keepalive for all configurations.
		IdleConnTimeout: 5 * time.Minute,
	}

	s := &Client{
		client: resty.New().
			OnBeforeRequest(NewRequestMiddleware(token)).
			SetHostURL(host).
			SetTransport(transport).
			SetTimeout(2 * time.Minute),
		prefix: Prefix,
		token:  token,
	}

	return s
}
