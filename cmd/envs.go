package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type Config struct {
	FbAccessToken string `env:"FB_ACCESS_TOKEN"`
	FbVerifySecret string `env:"FB_VERIFY_SECRET"`
	FbPageID string `env:"FB_PAGE_ID"`
}

func extractConfigFromEnv() (*Config, error) {
	config := new(Config)
	if err := env.Parse(config); err != nil {
		return nil, errors.Wrap(err, "environment variables failed at try to parsing")
	}
	return config, nil
}