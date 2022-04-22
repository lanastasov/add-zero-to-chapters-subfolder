package main

import (
	"fmt"
	"os"
	"os/exec"
)

// https://github.com/lanastasov/add-zero-to-chapters
func main() {
	fmt.Println("---version:0.0.3---")
	fmt.Println("---add-zero-to-chapters-subfolder---")
	fmt.Println("---generate-mp4-video-index-subfolder.exe---")
	fmt.Println("--- go get github.com/lanastasov/add-zero-to-chapters-subfolder ---")
	fmt.Println("--- go get -u github.com/lanastasov/add-zero-to-chapters-subfolder ---")
	
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
