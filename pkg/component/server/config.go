package server

import (
	"rockimserver/pkg/component/server/job"
	"rockimserver/pkg/component/server/socket/network/tcp"
	"rockimserver/pkg/component/server/socket/network/ws"
	"time"
)

type Config struct {
	Http   *Http
	Grpc   *Grpc
	Socket *Socket
	Job    *job.Config
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

type Socket struct {
	TCP *tcp.Config
	WS  *ws.Config
}
