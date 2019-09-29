package utils

import (
	"github.com/go-errors/errors"
	"os"
	"runtime"
	"strings"
)

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获取当前文件绝对路径
func CurrentFile() string {

	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return substr(file, 0, strings.LastIndex(file, "/"))
}
