package main

import (
	"fmt"
	"os/exec"
)

func main(){
	cmd0 := exec.Command("echo","-n","my first command from golang")
	// 创建一个能够获取此命令的输出的管道
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command No.0: %s\n", err)
		return
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
		return
	}
	output0 := make([]byte, 30)
	n, err := stdout0.Read(output0)
	if err != nil {
		fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
		return
	}
	fmt.Printf("%s\n", output0[:n])
}
