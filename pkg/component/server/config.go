package server

import "time"

type Config struct {
	Http *Http
	Grpc *Grpc
}

type Http struct {
	Network string
	Addr    string
	Timeout time.Duration
}

type Grpc struct {
	Network string
	Addr    string
	Timeout time.Duration
}
