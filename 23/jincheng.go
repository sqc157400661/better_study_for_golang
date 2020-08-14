package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Println("进程ID:",pid)
	fmt.Println("父进程ID:",ppid)
}