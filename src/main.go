package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var (
		cubeSize               = 3
		gamePadId        int32 = 0
		shouldExit             = false
		camera                 = rl.Camera3D{}
		cube                   = NewCube(cubeSize)
		selectedRotation       = R_RIGHT
	)

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	rl.SetTargetFPS(60)

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()

	rl.SetClipPlanes(0.5, 100)
	rl.DisableBackfaceCulling()

	prepareTextures()

	zoom := float32(11)
	camera.Position = rl.NewVector3(zoom, zoom, zoom)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 40.0
	camera.Projection = rl.CameraPerspective

	width, height, length := float32(2), float32(2), float32(2)

	for !rl.WindowShouldClose() && !shouldExit {
		//rl.UpdateCamera(&camera, rl.CameraThirdPerson)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.Color4f(1, 1, 1, 1)

		rl.BeginMode3D(camera)

		for xIterator := 0; xIterator < cube.size; xIterator++ {
			for yIterator := 0; yIterator < cube.size; yIterator++ {
				for zIterator := 0; zIterator < cube.size; zIterator++ {

					if zIterator == 0 || zIterator == 1 ||
						(zIterator == 2 && xIterator == 0 && yIterator == 0) ||
						(zIterator == 2 && xIterator == 0 && yIterator == 1) ||
						(zIterator == 2 && xIterator == 0 && yIterator == 2) ||
						(zIterator == 2 && xIterator == 1 && yIterator == 0) ||
						(zIterator == 2 && xIterator == 1 && yIterator == 1) ||
						(zIterator == 2 && xIterator == 1 && yIterator == 2) ||
						(zIterator == 2 && xIterator == 2 && yIterator == 0) ||
						(zIterator == 2 && xIterator == 2 && yIterator == 1) ||
						(zIterator == 2 && xIterator == 2 && yIterator == 2) {

						cubie := cube.cubies[xIterator][yIterator][zIterator]
						cubie.update()

						rl.PushMatrix()
						textures := colorTextures
						if cubie.shouldSelect(selectedRotation) {
							textures = selectedColorTextures
						}
						rl.Translatef(cubie.x*width, cubie.y*height, cubie.z*length)
						rl.Rotatef(cubie.angleX, cubie.vecX.X, cubie.vecX.Y, cubie.vecX.Z)
						rl.Rotatef(cubie.angleY, cubie.vecY.X, cubie.vecY.Y, cubie.vecY.Z)
						rl.Rotatef(cubie.angleZ, cubie.vecZ.X, cubie.vecZ.Y, cubie.vecZ.Z)

						rl.Begin(rl.Quads)
						{
							//front
							rl.SetTexture(textures[cubie.colors[FRONT]].ID)
							rl.TexCoord2f(0.0, 0.0)
							rl.Vertex3f(-width/2, -height/2, length/2)
							rl.TexCoord2f(1.0, 0.0)
							rl.Vertex3f(width/2, -height/2, length/2)
							rl.TexCoord2f(1.0, 1.0)
							rl.Vertex3f(width/2, height/2, length/2)
							rl.TexCoord2f(0.0, 1.0)
							rl.Vertex3f(-width/2, height/2, length/2)
							//back
							rl.SetTexture(textures[cubie.colors[BACK]].ID)
							rl.TexCoord2f(0.0, 0.0)
							rl.Vertex3f(-width/2, -height/2, -length/2)
							rl.TexCoord2f(1.0, 0.0)
							rl.Vertex3f(width/2, -height/2, -length/2)
							rl.TexCoord2f(1.0, 1.0)
							rl.Vertex3f(width/2, height/2, -length/2)
							rl.TexCoord2f(0.0, 1.0)
							rl.Vertex3f(-width/2, height/2, -length/2)
							//top
							rl.SetTexture(textures[cubie.colors[TOP]].ID)
							rl.TexCoord2f(0.0, 0.0)
							rl.Vertex3f(-width/2, height/2, length/2)
							rl.TexCoord2f(1.0, 0.0)
							rl.Vertex3f(width/2, height/2, length/2)
							rl.TexCoord2f(1.0, 1.0)
							rl.Vertex3f(width/2, height/2, -length/2)
							rl.TexCoord2f(0.0, 1.0)
							rl.Vertex3f(-width/2, height/2, -length/2)
							//bottom
							rl.SetTexture(textures[cubie.colors[BOTTOM]].ID)
							rl.TexCoord2f(0.0, 0.0)
							rl.Vertex3f(-width/2, -height/2, length/2)
							rl.TexCoord2f(1.0, 0.0)
							rl.Vertex3f(width/2, -height/2, length/2)
							rl.TexCoord2f(1.0, 1.0)
							rl.Vertex3f(width/2, -height/2, -length/2)
							rl.TexCoord2f(0.0, 1.0)
							rl.Vertex3f(-width/2, -height/2, -length/2)
							//left
							rl.SetTexture(textures[cubie.colors[LEFT]].ID)
							rl.TexCoord2f(0.0, 0.0)
							rl.Vertex3f(-width/2, -height/2, length/2)
							rl.TexCoord2f(1.0, 0.0)
							rl.Vertex3f(-width/2, -height/2, -length/2)
							rl.TexCoord2f(1.0, 1.0)
							rl.Vertex3f(-width/2, height/2, -length/2)
							rl.TexCoord2f(0.0, 1.0)
							rl.Vertex3f(-width/2, height/2, length/2)
							//right
							rl.SetTexture(textures[cubie.colors[RIGHT]].ID)
							rl.TexCoord2f(0.0, 0.0)
							rl.Vertex3f(width/2, -height/2, length/2)
							rl.TexCoord2f(1.0, 0.0)
							rl.Vertex3f(width/2, -height/2, -length/2)
							rl.TexCoord2f(1.0, 1.0)
							rl.Vertex3f(width/2, height/2, -length/2)
							rl.TexCoord2f(0.0, 1.0)
							rl.Vertex3f(width/2, height/2, length/2)
						}
						rl.End()
						rl.PopMatrix()
					}
				}
			}
		}

		//rl.DrawGrid(10, 1)
		rl.EndMode3D()

		if rl.IsKeyDown(rl.KeyZero) {
			selectedRotation = R_NONE
		}
		if rl.IsKeyDown(rl.KeyOne) {
			selectedRotation = R_FRONT
		}
		if rl.IsKeyDown(rl.KeyTwo) {
			selectedRotation = R_FB_MIDDLE
		}
		if rl.IsKeyDown(rl.KeyThree) {
			selectedRotation = R_BACK
		}
		if rl.IsKeyDown(rl.KeyFour) {
			selectedRotation = R_LEFT
		}
		if rl.IsKeyDown(rl.KeyFive) {
			selectedRotation = R_LR_MIDDLE
		}
		if rl.IsKeyDown(rl.KeySix) {
			selectedRotation = R_RIGHT
		}
		if rl.IsKeyDown(rl.KeySeven) {
			selectedRotation = R_TOP
		}
		if rl.IsKeyDown(rl.KeyEight) {
			selectedRotation = R_TB_MIDDLE
		}
		if rl.IsKeyDown(rl.KeyNine) {
			selectedRotation = R_BOTTOM
		}
		if rl.IsKeyDown(rl.KeyUp) {
			cube.startRotation(selectedRotation, true)
		}
		if rl.IsKeyDown(rl.KeyDown) {
			cube.startRotation(selectedRotation, false)
		}

		//exit
		if rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode) {
			shouldExit = true //see WindowShouldClose, it checks if KeyEscape pressed or Close icon pressed
		}

		rl.DrawFPS(5, 5)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
