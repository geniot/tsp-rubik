//go:build !wasm
// +build !wasm

package main

func main() {
	application := NewApplication(false)
	for !application.ShouldExit() {
		application.Update()
	}
	application.Exit()
}
