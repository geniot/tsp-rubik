package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameScene struct {
	a    *Application
	cube *Cube
}

func NewGameScene(a *Application) *GameScene {
	return &GameScene{cube: NewCube(3, split(CUBE_CORRECT), a)}
}

func (gs *GameScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (gs *GameScene) Update(camera *rl.Camera) {
	//rl.UpdateCamera(&camera, rl.CameraOrbital)
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.Color4f(1, 1, 1, 1)

	rl.BeginMode3D(*camera)
	gs.cube.update()
	gs.cube.draw()
	//rl.DrawGrid(10, 1)
	rl.EndMode3D()

	//rl.DrawText("The Breathing Cube", helpPadding*2, helpPadding*2, helpFontSize*2, rl.Blue)
	//rl.DrawText("It's breathing, when it's correct.", helpPadding*2+helpPadding/2, helpPadding*8, helpFontSize, rl.DarkGreen)

	//rl.DrawFPS(5, 5)
	rl.EndDrawing()
}
