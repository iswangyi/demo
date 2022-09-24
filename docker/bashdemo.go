package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("help")
	}
}

func run() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(syscall.Setenv("NAME", "ZHENGJUNXIONG"))
	println(os.Getenv("NAME"))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
