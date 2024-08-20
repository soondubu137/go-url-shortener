package utils

import (
	ierrors "go-url-shortener/errors"
	"net/url"
	"path"
)

// GetBaseURL returns the base URL of a given URL.
// For example, given the URL "https://example.com/1a2b3c?arg=val", the base URL is "1a2b3c".
func GetBaseURL(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", ierrors.ErrInvalidURL
	}
	return path.Base(parsed.Path), nil
}
