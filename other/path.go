package other

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//获取当前可执行文件所在路径
func GetExecPath() string {
	file, _ := exec.LookPath(os.Args[0])
	//fmt.Println("file ", file)
	path, _ := filepath.Abs(file)
	//fmt.Println("file path ", path)
	index := strings.LastIndex(path, string(os.PathSeparator))
	//fmt.Println("last index ", index)

	return path[:index]
}

