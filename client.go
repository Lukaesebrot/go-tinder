package tinder

import "github.com/valyala/fasthttp"

// Client represents an API client
type Client struct {
	token string
	http  *fasthttp.Client
}

// New creates a new client
func New(token string) *Client {
	return &Client{
		token: token,
		http: &fasthttp.Client{
			Name: "Tinder/7.5.3 (iPhone; iOS 10.3.2; Scale/2.00)",
		},
	}
}
