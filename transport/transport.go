package transport

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	FormatJson = "json"
	FormatXml = "xml"

	defaultURL = "https://api.novaposhta.ua"
	defaultVersion = "v2.0"
)

type Interface interface {
	Perform(*http.Request) (*http.Response, error)
}

type Config struct {
	Format string

	Transport http.RoundTripper
}

type Client struct {
	apiKey string
	format string
	url *url.URL

	transport http.RoundTripper
}

func New(cfg Config) *Client {

	if cfg.Transport == nil {
		cfg.Transport = http.DefaultTransport
	}

	// TODO only json supported now
	//if cfg.Format != FormatJson && cfg.Format != FormatXml {
	if cfg.Format != FormatJson {
		cfg.Format = FormatJson
	}

	u, _ := url.Parse(fmt.Sprintf("%s/%s/%s/", defaultURL, defaultVersion, cfg.Format))

	return &Client{
		transport: cfg.Transport,
		format: cfg.Format,
		url: u,
	}
}

func (c *Client) Perform(req *http.Request) (*http.Response, error) {

	baseUrl := c.getURL()

	if c.IsJsonFormat() {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", "text/xml")
	}

	req.URL.Scheme = baseUrl.Scheme
	req.URL.Host = baseUrl.Host

	if baseUrl.Path != "" {
		var b strings.Builder
		b.Grow(len(baseUrl.Path) + len(req.URL.Path))
		b.WriteString(baseUrl.Path)
		b.WriteString(req.URL.Path)
		req.URL.Path = b.String()
	}

	res, e := c.transport.RoundTrip(req)


	return res, e
}

func (c *Client) IsJsonFormat() bool {

	return c.format == FormatJson
}

func (c *Client) IsXmlFormat() bool {

	return c.format == FormatXml
}

func (c *Client) getURL() *url.URL {
	return c.url
}
