package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"rockimserver/app/access/gateway"
	"rockimserver/app/logic/user"
	"rockimserver/app/task/job"
	"rockimserver/pkg/log"
)

// go build -ldflags "-X main.version=x.y.z"
var (
	// version is the version of the compiled software.
	version string
)

type Runnable func(version string) (app *kratos.App, logger log.Logger, err error)

var (
	runnable = []Runnable{
		//admin.New,
		gateway.New,
		//comet.New,
		//platform.New,
		user.New,
		//group.New,
		//message.New,
		job.New,
	}
)

func main() {
	flag.Parse()
	for _, item := range runnable {
		go func(r Runnable) {
			app, _, err := r(version)
			if err != nil {
				panic(err)
			}
			err = app.Run()
			if err != nil {
				panic(err)
			}
		}(item)
	}
	err := kratos.New(
		kratos.Name("rockim-all-in-one"),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
	).Run()
	if err != nil {
		panic(err)
	}
}
