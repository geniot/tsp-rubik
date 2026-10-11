//go:build !wasm
// +build !wasm

package main

func main() {
	application := NewApplication(false)
	for !application.ShouldExit() {
		application.ProcessInput()
		application.Update()
		application.Render()
	}
	application.Exit()
}
