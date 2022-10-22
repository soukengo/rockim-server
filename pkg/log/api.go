package log

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	global.root.Debug(a...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	global.root.Debugf(format, a...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	global.root.Info(a...)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	global.root.Infof(format, a...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	global.root.Warn(a...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	global.root.Warnf(format, a...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	global.root.Error(a...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	global.root.Errorf(format, a...)
}
func Use(name string) Helper {
	return global.use(name)
}
func Default() Helper {
	return global.root
}
