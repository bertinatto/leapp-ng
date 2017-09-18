package web

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type Handler struct {
	mux     *http.ServeMux
	context context.Context
	errorCh chan error
}

func (h *Handler) Run() {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: h.mux,
	}

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		h.errorCh <- err
	} else {
		h.errorCh <- srv.Serve(listener)
	}
}

func (h *Handler) ErrorCh() <-chan error {
	return h.errorCh
}

func New() *Handler {
	h := &Handler{
		mux:     http.NewServeMux(),
		errorCh: make(chan error),
	}

	h.mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "{}")
	})

	return h
}
