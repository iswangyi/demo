package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd1 := exec.Command("umask", "0022")
	err := cmd1.Run()
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	if err != nil {
		fmt.Println(err)
	}

	cmd2 := exec.Command("umask")
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err = cmd2.Run()
	syscall.Umask(0027)
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("123", os.ModePerm)
	cmd := exec.Command("umask", "0000")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("u2", os.ModePerm)
	os.Mkdir("u3-777", 0777)
	os.Create("123")

}
