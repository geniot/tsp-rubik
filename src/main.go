package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var (
		gamePadId  int32 = 0
		shouldExit       = false
		camera           = rl.Camera3D{}
		angle            = float32(0)
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

		for i := range CubeDescriptors {

			if i >= 0 && i < 9 {
				rl.PushMatrix()
				rl.Rotatef(angle, 0, 1, 0)
				angle += 0 //.1
			}

			cube := CubeDescriptors[i]
			x, y, z := cube.x*width, cube.y*height, cube.z*length

			rl.Begin(rl.Quads)
			{
				//front
				rl.SetTexture(colorTextures[cube.frontColor].ID)
				rl.Normal3f(0.0, 0.0, 1.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
				//back
				rl.SetTexture(colorTextures[cube.backColor].ID)
				rl.Normal3f(0.0, 0.0, -1.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				//up
				rl.SetTexture(colorTextures[cube.upColor].ID)
				rl.Normal3f(0.0, 1.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				//down
				rl.SetTexture(colorTextures[cube.downColor].ID)
				rl.Normal3f(0.0, -1.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				//right
				rl.SetTexture(colorTextures[cube.rightColor].ID)
				rl.Normal3f(1.0, 0.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x+width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x+width/2, y+height/2, z+length/2)
				//left
				rl.SetTexture(colorTextures[cube.leftColor].ID)
				rl.Normal3f(-1.0, 0.0, 0.0)
				rl.TexCoord2f(0.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z+length/2)
				rl.TexCoord2f(1.0, 0.0)
				rl.Vertex3f(x-width/2, y-height/2, z-length/2)
				rl.TexCoord2f(1.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z-length/2)
				rl.TexCoord2f(0.0, 1.0)
				rl.Vertex3f(x-width/2, y+height/2, z+length/2)
			}
			rl.End()
			if i >= 0 && i < 9 {
				rl.PopMatrix()
			}
		}

		rl.DrawGrid(10, 1)
		rl.EndMode3D()

		//exit
		if rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode) {
			shouldExit = true //see WindowShouldClose, it checks if KeyEscape pressed or Close icon pressed
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
