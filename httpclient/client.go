package httpclient

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get(ctx context.Context, url string) (resp *http.Response, err error) {
	return DefaultClient.Get(ctx, url)
}

func Post(ctx context.Context, url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return DefaultClient.Post(ctx, url, contentType, body)
}

func PostForm(ctx context.Context, url string, data url.Values) (resp *http.Response, err error) {
	return DefaultClient.PostForm(ctx, url, data)
}

func Head(ctx context.Context, url string) (resp *http.Response, err error) {
	return DefaultClient.Head(ctx, url)
}

var DefaultClient = Client{http.DefaultClient}

type Client struct {
	Doer
}

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func (c *Client) Get(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Post(ctx context.Context, url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

func (c *Client) PostForm(ctx context.Context, url string, data url.Values) (resp *http.Response, err error) {
	return c.Post(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (c *Client) Head(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
