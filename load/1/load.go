package load

import (
	"io"
	"net/http"
)

type LoadTester struct {
	URL               string
	Output, ErrOutput io.Writer
	HTTPClient        *http.Client
	Fanout            int
	UserAgent         string
}

func NewLoadTester(URL string, stdout, stderr io.Writer, httpClient *http.Client, fanout int, userAgent string) *LoadTester {
	return &LoadTester{
		URL:        URL,
		Output:     stdout,
		ErrOutput:  stderr,
		HTTPClient: httpClient,
		Fanout:     fanout,
		UserAgent:  userAgent,
	}
}
