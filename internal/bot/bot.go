package bot

import (
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"pomogator/lib/e"
)

type Client struct {
	host     string
	basePath string
	Client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		Client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) (Updates, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	var err error
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (c *Client) SendMessage() {

}
