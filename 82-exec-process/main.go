package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 先找到执行路径
	lookPath, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	fmt.Println(lookPath)

	// 设置参数，第一个必须是程序名，这里是ls
	args := []string{"ls", "-a", "-l", "-h"}

	// 设置env
	env := os.Environ()

	// 执行，如果成功则go 进程功成身退，否则就处理错误
	err = syscall.Exec(lookPath, args, env)
	if err != nil {
		panic(err)
	}

}
