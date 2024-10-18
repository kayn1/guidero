package client

import (
	"context"
	"net/http"
	"net/url"

	"connectrpc.com/connect"
	v1 "github.com/kayn1/guidero/gen/proto/user/v1"
	"github.com/kayn1/guidero/gen/proto/user/v1/v1connect"
)

type GuideroClient struct {
	client     v1connect.UserServiceClient
	url        *url.URL
	httpClient *http.Client
}

type GuideroClientOption func(*GuideroClient)

func WithURL(url *url.URL) GuideroClientOption {
	return func(c *GuideroClient) {
		c.url = url
	}
}

func WithHTTPClient(httpClient *http.Client) GuideroClientOption {
	return func(c *GuideroClient) {
		c.httpClient = httpClient
	}
}

func newClientWithOptions(opts ...GuideroClientOption) (*GuideroClient, error) {
	client := &GuideroClient{}

	for _, opt := range opts {
		opt(client)
	}

	client.client = v1connect.NewUserServiceClient(client.httpClient, client.url.String())

	return client, nil
}

func NewClient(opts ...GuideroClientOption) (*GuideroClient, error) {
	defaultOptions := []GuideroClientOption{
		WithURL(&url.URL{
			Scheme: "http",
			Host:   "localhost:8080",
		}),
		WithHTTPClient(http.DefaultClient),
	}

	opts = append(defaultOptions, opts...)

	return newClientWithOptions(opts...)
}

func (c *GuideroClient) CreateUser(email, name string) (*v1.CreateResponse, error) {
	req := &v1.CreateRequest{
		Email: email,
		Name:  name,
	}

	connectReq := connect.NewRequest(req)
	resp, err := c.client.Create(context.Background(), connectReq)
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}
