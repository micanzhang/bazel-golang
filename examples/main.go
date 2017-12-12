package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	cmd := exec.Command("bazel", "version")
	bazelVer, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Go version: %s, bazel version: %s\n", runtime.Version(), string(bazelVer))
}
