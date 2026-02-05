package http_internal

import (
	"errors"
	"net/url"
)

func validateURL(raw string) error {
	if raw == "" {
		return errors.New("empty url")
	}

	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return errors.New("invalid url format")
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("unsupported url scheme")
	}

	if u.Host == "" {
		return errors.New("missing host")
	}

	return nil
}
