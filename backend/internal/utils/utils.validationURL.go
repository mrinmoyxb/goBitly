package utils

import (
	"net/url"
	"strings"
)

func IsURLValidUtil(originalURL string) bool {
	if originalURL == "" {
		return false
	}

	if len(originalURL) < 3 || len(originalURL) > 2048 {
		return false
	}

	parsedURL, err := url.ParseRequestURI(originalURL)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	if parsedURL.Host == "" {
		return false
	}

	if strings.HasPrefix(parsedURL.Host, ".") {
		return false
	}

	return true
}

func NormalizeURLUtil(originalURL string) (string, error) {
	u, err := url.Parse(originalURL)
	if err != nil {
		return "", err
	}

	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)

	if u.Path == "/" {
		u.Path = ""
	}

	return u.String(), nil
}
