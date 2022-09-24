package main

import (
	"fmt"
	libseccomp "github.com/seccomp/libseccomp-golang"
	"os"
	"os/exec"
)

func whiteList(syscalls []string) {

	filter, err := libseccomp.NewFilter(libseccomp.ActAllow)
	if err != nil {
		fmt.Printf("Error creating filter: %s", err)
	}
	for _, element := range syscalls {
		syscallID, err := libseccomp.GetSyscallFromName(element)
		if err != nil {
			panic(err)
		}
		filter.AddRule(syscallID, libseccomp.ActAllow)
	}
	filter.Load()
}

func main() {
	var syscalls = []string{}

	whiteList(syscalls)

	cmd := exec.Command("/bin/ls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
