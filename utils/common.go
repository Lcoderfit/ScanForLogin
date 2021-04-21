package utils

import (
	"net/url"
	"path"
)

// UrlJoin 对URL进行拼接
func UrlJoin(paths ...string) (string, error) {
	baseUrl, err := url.Parse(paths[0])
	if err != nil {
		return "", err
	}
	baseUrl.Path = path.Join(paths[1:]...)
	return baseUrl.String(), nil
}