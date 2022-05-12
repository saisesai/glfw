package main

import (
	"fmt"
	"github.com/saisesai/glfw"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Printf("glfw version: %d.%d.%d\n", glfw.VersionMajor, glfw.VersionMinor, glfw.VersionRevision)
	fmt.Println(glfw.DontCare)
}
