package main

import (
	"flag"
	"github.com/soukengo/gopkg/log"
	"rockimserver/app/access/admin"
)

// go build -ldflags "-X main.version=x.y.z"
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	flag.Parse()
	app, logger, err := admin.New(version)
	if err != nil {
		panic(err)
	}
	log.SetLogger(logger)
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
