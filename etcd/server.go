package main

import (
	"etcd/handler"
	"net/http"
)

type Server struct {
	server_map  map[string]string
	address_map map[string]string
}

// todo: 单例模式
func NewServer() *Server {
	return &Server{
		server_map:  make(map[string]string),
		address_map: make(map[string]string),
	}
}

func (s *Server) RUN() error {
	http.Handle("/register", &handler.RegisterHandler{})
	http.Handle("/get-serve", &handler.GetServeHandler{})

	err := http.ListenAndServe("127.0.0.1:8000", nil)
	return err
}
