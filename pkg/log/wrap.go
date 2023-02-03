package log

import (
	"context"
	"path/filepath"
	"rockimserver/pkg/util/file"
	"strings"
)

type leveledLogger struct {
	leveled map[Level]Logger
}

func newLeveledLogger(lv Level, cfg *LoggerConfig, kvs ...any) (Logger, error) {
	leveled := make(map[Level]Logger)
	absPath, err := file.Abs(cfg.Appender.FileName)
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(absPath)
	sourceName := filepath.Base(absPath)
	for _, level := range levelMap {
		if level < lv {
			continue
		}
		levelName := strings.ToLower(level.String())
		fileName := strings.Join([]string{dir, levelName}, string(filepath.Separator))
		fileName = fileName + sourceName
		ac := &AppenderConfig{Pattern: cfg.Appender.Pattern, FileName: fileName, RotationCount: cfg.Appender.RotationCount}
		leveled[level] = newAdapter(level, ac, kvs...)
	}
	return &leveledLogger{leveled: leveled}, nil
}

func (l leveledLogger) Log(level Level, kvs ...interface{}) error {
	if impl, ok := l.leveled[level]; ok {
		return impl.Log(level, kvs...)
	}
	return nil
}

func (l *leveledLogger) WithContext(ctx context.Context) Logger {
	leveled := make(map[Level]Logger)
	for level, item := range l.leveled {
		leveled[level] = item.WithContext(ctx)
	}
	return &leveledLogger{leveled: leveled}
}

type singleLogger struct {
	impl Logger
}

func newSingleLogger(lv Level, cfg *LoggerConfig, kvs ...any) (Logger, error) {
	return &singleLogger{impl: newAdapter(lv, cfg.Appender, kvs...)}, nil
}

func (l singleLogger) Log(level Level, kvs ...interface{}) error {
	return l.impl.Log(level, kvs...)
}
func (l *singleLogger) WithContext(ctx context.Context) Logger {
	return l.impl.WithContext(ctx)
}
