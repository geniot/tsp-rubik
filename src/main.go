package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var (
		gamePadId       int32 = 0
		shouldExit            = false
		camera                = rl.Camera3D{}
		rotation              = NewRotation()
		cube                  = NewCube()
		currentRotation       = R_NONE
	)

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(1)
	rl.InitAudioDevice()

	rl.SetClipPlanes(0.5, 100)
	rl.DisableBackfaceCulling()

	prepareTextures()

	camera.Position = rl.NewVector3(10.0, 10.0, 10.0)
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

		dimensions := 3
		for xIterator := 0; xIterator < dimensions; xIterator++ {
			for yIterator := 0; yIterator < dimensions; yIterator++ {
				for zIterator := 0; zIterator < dimensions; zIterator++ {

					if yIterator == 2 {
						rl.PushMatrix()
						rl.Rotatef(rotation.angleX, 1, 0, 0)
						rl.Rotatef(rotation.angleY, 0, 1, 0)
						rl.Rotatef(rotation.angleZ, 0, 0, 1)

						if !rotation.update() {
							cube.rotate(currentRotation)
						}
					}

					x, y, z := float32(xIterator-1)*width, float32(yIterator-1)*height, float32(zIterator-1)*length
					rl.Begin(rl.Quads)
					{
						//front
						rl.SetTexture(colorTextures[cube.cubies[xIterator][yIterator][zIterator].colors[FRONT]].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z+length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z+length/2)
						//back
						rl.SetTexture(colorTextures[cube.cubies[xIterator][yIterator][zIterator].colors[BACK]].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z-length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z-length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z-length/2)
						//top
						rl.SetTexture(colorTextures[cube.cubies[xIterator][yIterator][zIterator].colors[TOP]].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y+height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y+height/2, z+length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z-length/2)
						//bottom
						rl.SetTexture(colorTextures[cube.cubies[xIterator][yIterator][zIterator].colors[BOTTOM]].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y-height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y-height/2, z-length/2)
						//left
						rl.SetTexture(colorTextures[cube.cubies[xIterator][yIterator][zIterator].colors[LEFT]].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z-length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z+length/2)
						//right
						rl.SetTexture(colorTextures[cube.cubies[xIterator][yIterator][zIterator].colors[RIGHT]].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z-length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z+length/2)
					}
					rl.End()
					if yIterator == 2 {
						rl.PopMatrix()
					}

				}
			}
		}

		rl.DrawGrid(10, 1)
		rl.EndMode3D()

		if rl.IsKeyDown(rl.KeyLeft) {
			if rl.IsKeyDown(rl.KeyLeftControl) {
				rotation.rotateX(90)
			} else {
				rotation.rotateY(-90)
				currentRotation = R_TOP_F
			}
		}
		if rl.IsKeyDown(rl.KeyRight) {
			if rl.IsKeyDown(rl.KeyLeftControl) {
				rotation.rotateX(-90)
			} else {
				rotation.rotateY(90)
				currentRotation = R_TOP_B
			}
		}
		if rl.IsKeyDown(rl.KeyUp) {
			rotation.rotateZ(90)
		}
		if rl.IsKeyDown(rl.KeyDown) {
			rotation.rotateZ(-90)
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
