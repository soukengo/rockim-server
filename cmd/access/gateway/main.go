package main

import (
	"flag"
	"rockimserver/app/access/gateway"
	"rockimserver/pkg/log"
)

// go build -ldflags "-X main.version=x.y.z"
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	flag.Parse()
	app, logger, err := gateway.New(version)
	if err != nil {
		panic(err)
	}
	log.SetLogger(logger)
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
