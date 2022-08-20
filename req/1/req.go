package req

import (
	"errors"
	"net/http"
)

var ErrRateLimit = errors.New("rate limit")

func Request(URL string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusTooManyRequests {
		return ErrRateLimit
	}
	return nil
}
