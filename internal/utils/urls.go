package utils

import (
	"errors"
	"net/url"
	"strings"
)

func BuildURL(baseURL, path, query string) (*url.URL, error) {

	if baseURL == "" {
		return nil, errors.New("missing base url")
	}

	safeURL := baseURL

	if !strings.HasPrefix(baseURL, "http") {
		safeURL = "http://" + baseURL
	}

	if path != "" {
		safeURL = strings.Join([]string{
			safeURL,
			path,
		}, "")
	}

	if query != "" {
		safeURL = strings.Join([]string{
			safeURL,
			query,
		}, "?")
	}

	forwardURL, err := url.Parse(safeURL)

	return forwardURL, err
}
