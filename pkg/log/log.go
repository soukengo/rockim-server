package log

import (
	"context"
	"strings"
)

func Configure(cfg *Config) (err error) {
	kvs := []any{"service.id", cfg.AppId, "service.version", cfg.AppVersion}
	root, err := newLogger(&cfg.LoggerConfig, kvs...)
	if err != nil {
		return
	}
	global.setLogger(newHelper(root))
	for _, item := range cfg.Loggers {
		var lc = item
		l, err := newLogger(&lc, kvs...)
		if err != nil {
			return err
		}
		global.register(lc.Name, newHelper(l))
	}
	// 设置kratos框架默认Logger
	configureKratos(root)
	return
}

type Logger interface {
	Log(level Level, keyvals ...interface{}) error
	WithContext(ctx context.Context) Logger
}

func newLogger(cfg *LoggerConfig, kvs ...any) (Logger, error) {
	lv, ok := levelMap[strings.ToUpper(cfg.Level)]
	if !ok {
		lv = LevelInfo
	}
	if cfg.Appender == nil || cfg.Appender.Console || len(cfg.Appender.FileName) == 0 || !cfg.Split {
		return newSingleLogger(lv, cfg, kvs...)
	}
	return newLeveledLogger(lv, cfg, kvs...)
}
