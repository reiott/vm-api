package http

import "net/http"

type MuxOpts struct{}

func NewMux(opts MuxOpts, svc Service) http.Handler {
	return nil
}
