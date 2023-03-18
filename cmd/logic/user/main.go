package main

import (
	"flag"
	"rockimserver/app/logic/user"
	"rockimserver/pkg/log"
)

// go build -ldflags "-X main.version=x.y.z"
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	flag.Parse()
	app, logger, err := user.New(version)
	if err != nil {
		panic(err)
	}
	log.SetLogger(logger)
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
