package main

import (
	"log"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
	log.SetFlags(log.Lshortfile)
}

func main() {
	NewApplication(true).Start(os.Args)
}
