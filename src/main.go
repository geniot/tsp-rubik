package main

func main() {
	application := NewApplication()
	for !application.ShouldExit() {
		application.Draw()
	}
	application.Exit()
}
