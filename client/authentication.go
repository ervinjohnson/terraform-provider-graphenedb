package client

import (
	"net/http"
)

type AuthenticatedClient struct {
	*Client
}

func (c *Client) Authenticate() (*AuthenticatedClient, error) {
	return &AuthenticatedClient{
		Client: c,
	}, nil
}

func (c *AuthenticatedClient) newAuthenticatedRequest(builder func() (*http.Request, error)) (*http.Request, error) {
	req, err := builder()
	if err != nil {
		return nil, err
	}

	// if err = c.setAuthenticationCookie(req); err != nil {
	// 	return nil, err
	// }

	return req, nil
}

func (c *AuthenticatedClient) newAuthenticatedPostRequest(path string, body interface{}) (*http.Request, error) {
	return c.newAuthenticatedRequest(func() (*http.Request, error) { return c.newPostRequest(path, body) })
}
