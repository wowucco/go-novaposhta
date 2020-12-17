package novaposhta

import (
	"errors"
	"github.com/wowucco/go-novaposhta/api"
	"github.com/wowucco/go-novaposhta/transport"
	"net/http"
)

type Config struct {
	ApiKey string
	Format string

	Transport http.RoundTripper
}

type Client struct {
	*api.API
	Transport transport.Interface
}

func NewClient(cfg Config) (*Client, error) {

	if cfg.ApiKey == "" {
		return nil, errors.New("cannot create client: api key is required")
	}

	t := transport.New(transport.Config{
		Transport: cfg.Transport,

		Format: cfg.Format,
	})

	return &Client{
		Transport: t,
		API: api.New(t, cfg.ApiKey),
	}, nil
}
