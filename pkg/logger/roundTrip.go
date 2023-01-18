package logger

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	Logger io.Writer
	Next   http.RoundTripper
}

func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.Logger, "Time: [%s] | Method: [%s]", time.Now().Format(time.RFC1123), r.Method)
	return l.Next.RoundTrip(r)
}
