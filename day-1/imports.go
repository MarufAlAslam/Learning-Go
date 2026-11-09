package main

import (
	"fmt"
	"runtime"
)

func showInfo() {
	fmt.Printf("OS: %s GO version: %s", runtime.GOOS, runtime.Version())
}
