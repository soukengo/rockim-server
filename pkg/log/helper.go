package log

import (
	"context"
	"fmt"
)

var DefaultMessageKey = "msg"

type Helper interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	WithContext(context.Context) Helper
}

type helper struct {
	msgKey string
	logger Logger
}

func newHelper(logger Logger) Helper {
	return &helper{logger: logger, msgKey: DefaultMessageKey}
}

func (h helper) Debug(args ...interface{}) {
	_ = h.logger.Log(LevelDebug, h.msgKey, fmt.Sprint(args...))
}

func (h helper) Debugf(format string, args ...interface{}) {
	_ = h.logger.Log(LevelDebug, h.msgKey, fmt.Sprintf(format, args...))
}

func (h helper) Info(args ...interface{}) {
	_ = h.logger.Log(LevelInfo, h.msgKey, fmt.Sprint(args...))
}

func (h helper) Infof(format string, args ...interface{}) {
	_ = h.logger.Log(LevelInfo, h.msgKey, fmt.Sprintf(format, args...))
}

func (h helper) Warn(args ...interface{}) {
	_ = h.logger.Log(LevelWarn, h.msgKey, fmt.Sprint(args...))
}

func (h helper) Warnf(format string, args ...interface{}) {
	_ = h.logger.Log(LevelWarn, h.msgKey, fmt.Sprintf(format, args...))
}

func (h helper) Error(args ...interface{}) {
	_ = h.logger.Log(LevelError, h.msgKey, fmt.Sprint(args...))
}

func (h helper) Errorf(format string, args ...interface{}) {
	_ = h.logger.Log(LevelError, h.msgKey, fmt.Sprintf(format, args...))
}

func (h helper) WithContext(ctx context.Context) Helper {
	return &helper{
		msgKey: h.msgKey,
		logger: h.logger.WithContext(ctx),
	}
}

func wrap(hpl Helper) *logWrapper {
	return &logWrapper{hpl: hpl}
}

type logWrapper struct {
	hpl Helper
}

func (p *logWrapper) Debug(args ...interface{}) {
	p.hpl.Debug(args...)
}

func (p *logWrapper) Debugf(format string, args ...interface{}) {
	p.hpl.Debugf(format, args...)
}

func (p *logWrapper) Info(args ...interface{}) {
	p.hpl.Info(args...)
}

func (p *logWrapper) Infof(format string, args ...interface{}) {
	p.hpl.Infof(format, args...)
}

func (p *logWrapper) Warn(args ...interface{}) {
	p.hpl.Warn(args...)
}

func (p *logWrapper) Warnf(format string, args ...interface{}) {
	p.hpl.Warnf(format, args...)
}

func (p logWrapper) Error(args ...interface{}) {
	p.hpl.Error(args...)
}

func (p *logWrapper) Errorf(format string, args ...interface{}) {
	p.hpl.Errorf(format, args...)
}

func (p *logWrapper) WithContext(ctx context.Context) Helper {
	return wrap(wrap(p.hpl.WithContext(ctx)))
}

type emptyHelper struct {
}

func (l *emptyHelper) Debug(args ...interface{}) {

}
func (l *emptyHelper) Debugf(format string, args ...interface{}) {

}

func (l *emptyHelper) Info(args ...interface{}) {

}

func (l *emptyHelper) Infof(format string, args ...interface{}) {

}

func (l *emptyHelper) Warn(args ...interface{}) {

}

func (l *emptyHelper) Warnf(format string, args ...interface{}) {

}

func (l *emptyHelper) Error(args ...interface{}) {

}

func (l *emptyHelper) Errorf(format string, args ...interface{}) {

}

func (l *emptyHelper) WithContext(ctx context.Context) Helper {
	return l
}
