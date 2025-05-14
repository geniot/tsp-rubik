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
		selectedRotation       = R_NONE
		isForward              = true
		rotation               = NewRotation()
	)

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()

	rl.SetClipPlanes(0.5, 100)
	rl.DisableBackfaceCulling()

	prepareTextures()

	camera.Position = rl.NewVector3(float32(10.0*cubeSize/2), float32(10.0*cubeSize/2), float32(10.0*cubeSize/2))
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
		//rl.LoadIdentity()

		for xIterator := 0; xIterator < cube.size; xIterator++ {
			for yIterator := 0; yIterator < cube.size; yIterator++ {
				for zIterator := 0; zIterator < cube.size; zIterator++ {

					rl.PushMatrix()

					//cubie := cube.cubies[xIterator+yIterator*(cube.size-1)+zIterator*(cube.size-1)]

					textures := colorTextures
					if cube.shouldSelect(selectedRotation, xIterator, yIterator, zIterator) {
						textures = selectedColorTextures
						if selectedRotation != R_NONE {
							rl.Rotatef(rotation.angleX, 1, 0, 0)
							rl.Rotatef(rotation.angleY, 0, 1, 0)
							rl.Rotatef(rotation.angleZ, 0, 0, 1)
							rotation.update()
						}
					}

					x, y, z := float32(xIterator-1)*width, float32(yIterator-1)*height, float32(zIterator-1)*length
					rl.Begin(rl.Quads)
					{
						//front
						rl.SetTexture(textures[cube.getColor(xIterator, yIterator, zIterator, FRONT)].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z+length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z+length/2)
						//back
						rl.SetTexture(textures[cube.getColor(xIterator, yIterator, zIterator, BACK)].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z-length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z-length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z-length/2)
						//top
						rl.SetTexture(textures[cube.getColor(xIterator, yIterator, zIterator, TOP)].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y+height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y+height/2, z+length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y+height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z-length/2)
						//bottom
						rl.SetTexture(textures[cube.getColor(xIterator, yIterator, zIterator, BOTTOM)].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x+width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x+width/2, y-height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y-height/2, z-length/2)
						//left
						rl.SetTexture(textures[cube.getColor(xIterator, yIterator, zIterator, LEFT)].ID)
						rl.TexCoord2f(0.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z+length/2)
						rl.TexCoord2f(1.0, 0.0)
						rl.Vertex3f(x-width/2, y-height/2, z-length/2)
						rl.TexCoord2f(1.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z-length/2)
						rl.TexCoord2f(0.0, 1.0)
						rl.Vertex3f(x-width/2, y+height/2, z+length/2)
						//right
						rl.SetTexture(textures[cube.getColor(xIterator, yIterator, zIterator, RIGHT)].ID)
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
					rl.PopMatrix()

				}
			}
		}

		rl.DrawGrid(10, 1)
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

		//if rl.IsKeyDown(rl.KeyLeft) {
		//	if rl.IsKeyDown(rl.KeyLeftControl) {
		//		rotation.rotateX(90)
		//	} else {
		//		rotation.rotateY(-90)
		//	}
		//}
		//if rl.IsKeyDown(rl.KeyRight) {
		//	if rl.IsKeyDown(rl.KeyLeftControl) {
		//		rotation.rotateX(-90)
		//	} else {
		//		rotation.rotateY(90)
		//	}
		//}

		if rl.IsKeyDown(rl.KeyUp) {
			isForward = true
			rotation.rotate(selectedRotation, isForward)
		}
		if rl.IsKeyDown(rl.KeyDown) {
			isForward = false
			rotation.rotate(selectedRotation, isForward)
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
