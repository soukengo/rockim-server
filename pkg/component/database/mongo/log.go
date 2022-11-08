package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"golang.org/x/net/context"
	"rockim/pkg/log"
)

const (
	loggerName = "mongo"
)

type logger struct {
}

func newLogger() Monitor {
	return &logger{}
}

func (l *logger) Started(ctx context.Context, event *event.CommandStartedEvent) {
	if !log.IsEnabled(loggerName) {
		return
	}
	statement, _ := bson.MarshalExtJSON(event.Command, false, false)
	log.Use(loggerName).WithContext(ctx).Infof("exec[%v] started command: %s, statement: %s", event.RequestID, event.CommandName, string(statement))
}

func (l *logger) Succeeded(ctx context.Context, event *event.CommandSucceededEvent) {
	if !log.IsEnabled(loggerName) {
		return
	}
	log.Use(loggerName).WithContext(ctx).Infof("exec[%v] succeeded command: %s,reply: %s", event.RequestID, event.CommandName)
}

func (l *logger) Failed(ctx context.Context, event *event.CommandFailedEvent) {
	if !log.IsEnabled(loggerName) {
		return
	}
	log.Use(loggerName).WithContext(ctx).Errorf("exec[%v] failed command: %s", event.RequestID, event.CommandName)
}
