package redis

import (
	"time"
)

type Config struct {
	Addr         []string
	Password     string
	Prefix       string
	Active       int
	Idle         int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}
