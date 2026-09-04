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
