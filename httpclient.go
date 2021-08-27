package httpclient

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type client struct {
	http.Client
}

func New() *client {
	return &client{http.Client{}}
}
func setHeader(req *http.Request, header map[string]string) {
	if header == nil {
		return
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
}

func (c *client) Head(url string, header map[string]string) (*http.Response, error) {
	return c.HeadWithContext(context.Background(), url, header)
}
func (c *client) HeadWithContext(ctx context.Context, url string, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	setHeader(req, header)
	return c.Do(req)
}

func (c *client) Get(url string, header map[string]string) (*http.Response, error) {
	return c.GetWithContext(context.Background(), url, header)
}
func (c *client) GetWithContext(ctx context.Context, url string, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	setHeader(req, header)
	return c.Do(req)
}

func (c *client) Post(url string, header map[string]string, body io.Reader) (*http.Response, error) {
	return c.PostWithContext(context.Background(), url, header, body)
}
func (c *client) PostWithContext(ctx context.Context, url string, header map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}
	setHeader(req, header)
	return c.Do(req)
}
func (c *client) PostForm(url string, header map[string]string, data url.Values) (*http.Response, error) {
	return c.PostFormWithContext(context.Background(), url, header, data)
}
func (c *client) PostFormWithContext(ctx context.Context, url string, header map[string]string, data url.Values) (*http.Response, error) {
	if header == nil {
		header = make(map[string]string)
	}
	header["Content-Type"] = "application/x-www-form-urlencoded"
	return c.PostWithContext(ctx, url, header, strings.NewReader(data.Encode()))
}

func (c *client) Put(url string, header map[string]string, body io.Reader) (*http.Response, error) {
	return c.PutWithContext(context.Background(),url,header,body)
}
func (c *client) PutWithContext(ctx context.Context, url string, header map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "PUT", url, body)
	if err != nil {
		return nil, err
	}
	setHeader(req, header)
	return c.Do(req)
}

func (c *client) PutForm(url string, header map[string]string, data url.Values) (*http.Response, error) {
	return c.PutFormWithContext(context.Background(),url, header, data)
}
func (c *client) PutFormWithContext(ctx context.Context,url string, header map[string]string, data url.Values) (*http.Response, error) {
	if header == nil {
		header = make(map[string]string)
	}
	header["Content-Type"] = "application/x-www-form-urlencoded"
	return c.PutWithContext(ctx,url, header, strings.NewReader(data.Encode()))
}

func (c *client) Delete(url string, header map[string]string) (*http.Response, error) {
	return c.DeleteWithContext(context.Background(),url,header)
}
func (c *client) DeleteWithContext(ctx context.Context,url string, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx,"DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	setHeader(req, header)
	return c.Do(req)
}
