package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("---add-zero-to-chapters-recursive---")
	// call add-zero-to-chapters.go in all sub folders

	cmd, _ := exec.Command("add-zero-to-chapters").Output()
	fmt.Println("cmd=", cmd)

	dir, _ := os.Open(".")
	defer dir.Close()

	fileInfo, _ := dir.Readdir(-1)

	for _, file := range fileInfo {
		cmd2 := exec.Command("add-zero-to-chapters")
		fmt.Println(file.Name())
		cmd2.Dir = file.Name()
		cmd2.Output()
	}
}
