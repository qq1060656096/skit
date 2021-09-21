package file

import (
	"path"
	"runtime"
)

// /Users/zhaoweijie/go/src/github.com/qq1060656096/skit/file/file.go
func CurrentFile() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}

// /Users/zhaoweijie/go/src/github.com/qq1060656096/skit/file
func CurrentFileDir() string {
	file := CurrentFile()
	return path.Dir(file)
}