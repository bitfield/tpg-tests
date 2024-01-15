package load

import (
	"io"
	"net/http"
	"os"
)

type LoadTester struct {
	URL               string
	Output, ErrOutput io.Writer
	HTTPClient        *http.Client
	Fanout            int
	UserAgent         string
}

func NewLoadTester(URL string, opts ...TesterOption) *LoadTester {
	lt := &LoadTester{
		URL:        URL,
		Output:     os.Stdout,
		ErrOutput:  os.Stderr,
		HTTPClient: http.DefaultClient,
		Fanout:     10,
		UserAgent:  "load v2",
	}
	for _, opt := range opts {
		opt(lt)
	}
	return lt
}

type TesterOption func(*LoadTester)

func WithOutput(w io.Writer) TesterOption {
	return func(lt *LoadTester) {
		lt.Output = w
	}
}

func WithErrOutput(w io.Writer) TesterOption {
	return func(lt *LoadTester) {
		lt.ErrOutput = w
	}
}

func WithHTTPClient(c *http.Client) TesterOption {
	return func(lt *LoadTester) {
		lt.HTTPClient = c
	}
}

func WithConcurrentRequests(fanout int) TesterOption {
	return func(lt *LoadTester) {
		lt.Fanout = fanout
	}
}

func WithUserAgent(agent string) TesterOption {
	return func(lt *LoadTester) {
		lt.UserAgent = agent
	}
}
