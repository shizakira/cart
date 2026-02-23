package product

import (
	"fmt"
	"net/http"
	"time"
)

type RetryTransport struct {
	Base       http.RoundTripper
	MaxRetries int
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.MaxRetries; attempt++ {
		resp, err = t.Base.RoundTrip(req)
		if err != nil {
			return nil, fmt.Errorf("t.Base.RoundTrip: %w", err)
		}

		if resp.StatusCode != 420 && resp.StatusCode != 429 {
			return resp, nil
		}

		resp.Body.Close()

		if attempt == t.MaxRetries {
			return resp, nil
		}

		time.Sleep(time.Duration(500*(1<<attempt)) * time.Millisecond)
	}

	return resp, err
}
