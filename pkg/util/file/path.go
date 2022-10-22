package file

import (
	"path/filepath"
	"rockim/pkg/util/env"
)

func Abs(fileName string) (dist string, err error) {
	if filepath.IsAbs(fileName) {
		dist = fileName
		return
	}
	dir, err := env.RootDir()
	if err != nil {
		return
	}
	dist = dir + string(filepath.Separator) + fileName
	dist, err = filepath.Abs(dist)
	return
}
