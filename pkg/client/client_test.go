package client

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
}

func (suite *ClientTestSuite) SetupTest() {
}

func (suite *ClientTestSuite) TestNewClient() {
	tests := []struct {
		name           string
		options        []GuideroClientOption
		expectedScheme string
		expectedHost   string
		expectedClient *http.Client
	}{
		{
			name:           "DefaultOptions",
			options:        nil,
			expectedScheme: "http",
			expectedHost:   "localhost:8080",
			expectedClient: http.DefaultClient,
		},
		{
			name: "CustomURL",
			options: []GuideroClientOption{
				WithURL(&url.URL{
					Scheme: "http",
					Host:   "customhost:9090",
				}),
			},
			expectedScheme: "http",
			expectedHost:   "customhost:9090",
			expectedClient: http.DefaultClient,
		},
		{
			name: "CustomHTTPClient",
			options: []GuideroClientOption{
				WithHTTPClient(&http.Client{}),
			},
			expectedScheme: "http",
			expectedHost:   "localhost:8080",
			expectedClient: &http.Client{},
		},
		{
			name: "CustomOptions",
			options: []GuideroClientOption{
				WithURL(&url.URL{
					Scheme: "http",
					Host:   "customhost:9090",
				}),
				WithHTTPClient(&http.Client{}),
			},
			expectedScheme: "http",
			expectedHost:   "customhost:9090",
			expectedClient: &http.Client{},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			client, err := NewClient(tt.options...)
			assert.NoError(suite.T(), err)
			assert.NotNil(suite.T(), client)
			assert.Equal(suite.T(), tt.expectedScheme, client.url.Scheme)
			assert.Equal(suite.T(), tt.expectedHost, client.url.Host)
			assert.Equal(suite.T(), tt.expectedClient, client.httpClient)
		})
	}
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
