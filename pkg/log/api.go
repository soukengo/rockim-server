package log

import "context"

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	global.rootProxy.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	global.rootProxy.Debugf(format, a...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	global.rootProxy.Info(a...)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	global.rootProxy.Infof(format, a...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	global.rootProxy.Warn(a...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	global.rootProxy.Warnf(format, a...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	global.rootProxy.Error(a...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	global.rootProxy.Errorf(format, a...)
}

func Default() Helper {
	return global.root
}

func IsEnabled(name string) bool {
	return global.has(name)
}

func Std() Helper {
	return stdProxy
}
func Use(loggerName string) Helper {
	l := global.use(loggerName)
	return l
}

func WithContext(ctx context.Context) Helper {
	l := wrap(global.root.WithContext(ctx))
	return l
}
