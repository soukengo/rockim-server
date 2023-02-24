package main

import (
	"flag"
	"rockimserver/app/logic/group"
)

// go build -ldflags "-X main.version=x.y.z"
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	flag.Parse()
	app, err := group.New(version)
	if err != nil {
		panic(err)
	}
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
