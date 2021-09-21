package path

import (
	"os"
	"path"
	"runtime"
)

func getCurrentPkgFilePath() string {
	_, file, _, _ := runtime.Caller(2)
	return file
}

// CurrentFilePath 返回当前文件路径
func CurrentFilePath() string {
	return getCurrentPkgFilePath()
}

// CurrentFileDirPath 返回当前文件所在目录路径
func CurrentFileDirPath() string {
	file := getCurrentPkgFilePath()
	return path.Dir(file)
}

// ExecutablePath 返回可执行目录路径
func ExecutablePath() (path string, err error){
	return os.Executable()
}