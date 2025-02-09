package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dataCmd := exec.Command("date")

	dataOut, err := dataCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dataOut))

	// 判断错误类型
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("faile to executing")
		case *exec.ExitError:
			fmt.Println("command exit with rc = ", e.ExitCode())
		default:
			panic(err)
		}
	}
	//
	grepCmd := exec.Command("grep", "hello")
	out, _ := grepCmd.StdoutPipe()
	in, _ := grepCmd.StdinPipe()
	_, _ = grepCmd.StderrPipe()
	grepCmd.Start()
	in.Write([]byte("hello grep\nbye grep"))
	in.Close()
	outBytes, _ := io.ReadAll(out)
	grepCmd.Wait() // 等待命令执行完成

	fmt.Println("> grep hello")
	fmt.Println(string(outBytes))

	//
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
