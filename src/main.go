package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"unsafe"
)

func main() {

	//rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint) //should be set before window initialization!
	rl.SetTargetFPS(60)

	rl.InitWindow(1280, 720, "TrimUI Rubik")
	rl.SetWindowMonitor(0)
	rl.InitAudioDevice()

	prepareTextures()

	var (
		cubeSize               = 3
		gamePadId        int32 = 0
		shouldExit             = false
		camera                 = rl.Camera3D{}
		cube                   = NewCube(cubeSize)
		selectedRotation       = R_RIGHT
	)

	rl.SetClipPlanes(0.5, 100)
	rl.DisableBackfaceCulling()

	//mesh := GenMeshCustom()
	//mesh := rl.GenMeshCube(2, 2, 2)
	////texCoords := []float32{0,0.5,0.5,1.0}
	//var texcoords []float32
	//mesh.Texcoords = unsafe.SliceData(texcoords)
	//model := rl.LoadModelFromMesh(GenMeshCustom())
	//model := rl.LoadModelFromMesh(mesh)
	//rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, combinedTexture)
	//rl.SetModelMeshMaterial(&model, 0, int32(combinedTexture.ID))

	zoom := float32(11)
	camera.Position = rl.NewVector3(zoom, zoom, zoom)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 40.0
	camera.Projection = rl.CameraPerspective

	for !rl.WindowShouldClose() && !shouldExit {
		rl.UpdateCamera(&camera, rl.CameraThirdPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.Color4f(1, 1, 1, 1)

		rl.BeginMode3D(camera)

		for xIterator := 0; xIterator < cube.size; xIterator++ {
			for yIterator := 0; yIterator < cube.size; yIterator++ {
				for zIterator := 0; zIterator < cube.size; zIterator++ {

					//if zIterator == 0 || zIterator == 1 ||
					//	(zIterator == 2 && xIterator == 0 && yIterator == 0) ||
					//	(zIterator == 2 && xIterator == 0 && yIterator == 1) ||
					//	(zIterator == 2 && xIterator == 0 && yIterator == 2) ||
					//	(zIterator == 2 && xIterator == 1 && yIterator == 0) ||
					//	(zIterator == 2 && xIterator == 1 && yIterator == 1) ||
					//	(zIterator == 2 && xIterator == 1 && yIterator == 2) ||
					//	(zIterator == 2 && xIterator == 2 && yIterator == 0) ||
					//	(zIterator == 2 && xIterator == 2 && yIterator == 1) ||
					//	(zIterator == 2 && xIterator == 2 && yIterator == 2) {

					cubie := cube.cubies[xIterator][yIterator][zIterator]
					cubie.update()
					cubie.model.Transform = rl.MatrixRotateXYZ(rl.Vector3{X: rl.Deg2rad * cubie.angleX, Y: rl.Deg2rad * cubie.angleY, Z: rl.Deg2rad * cubie.angleZ})
					rl.DrawModel(cubie.model, rl.Vector3{X: cubie.x, Y: cubie.y, Z: cubie.z}, 1.0, rl.White)
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

// GenMeshCustom generates a simple triangle mesh from code
func GenMeshCustom() rl.Mesh {
	mesh := rl.Mesh{
		TriangleCount: 1,
		VertexCount:   6,
	}

	var vertices, normals, texcoords []float32

	// 4 vertices
	vertices = addCoord(vertices, 0, 0, 0)
	vertices = addCoord(vertices, 0, 0, 2)
	vertices = addCoord(vertices, 2, 0, 2)

	vertices = addCoord(vertices, 0, 0, 0)
	vertices = addCoord(vertices, 2, 0, 0)
	vertices = addCoord(vertices, 2, 0, 2)
	mesh.Vertices = unsafe.SliceData(vertices)

	// 4 normals
	//normals = addCoord(normals, 0, 1, 0)
	//normals = addCoord(normals, 0, 1, 0)
	//normals = addCoord(normals, 0, 1, 0)
	//normals = addCoord(normals, 0, 1, 0)
	//normals = addCoord(normals, 0, 1, 0)
	//normals = addCoord(normals, 0, 1, 0)

	normals = addCoord(normals, 1, 0, 0)
	normals = addCoord(normals, 1, 0, 0)
	normals = addCoord(normals, 1, 0, 0)
	normals = addCoord(normals, 1, 0, 0)
	normals = addCoord(normals, 1, 0, 0)
	normals = addCoord(normals, 1, 0, 0)
	mesh.Normals = unsafe.SliceData(normals)

	// 4 texcoords
	offsetX := float32(0.0)
	offsetY := float32(0.0)

	texcoords = addCoord(texcoords, 0+offsetX, 0+offsetY)
	texcoords = addCoord(texcoords, 0+offsetX, 0.5+offsetY)
	texcoords = addCoord(texcoords, 0.5+offsetX, 0.5+offsetY)

	texcoords = addCoord(texcoords, 0+offsetX, 0+offsetY)
	texcoords = addCoord(texcoords, 0+offsetX, 0.5+offsetY)
	texcoords = addCoord(texcoords, 0.5+offsetX, 0.5+offsetY)
	mesh.Texcoords = unsafe.SliceData(texcoords)

	// Upload mesh data from CPU (RAM) to GPU (VRAM) memory
	rl.UploadMesh(&mesh, false)

	return mesh
}

func addCoord(slice []float32, values ...float32) []float32 {
	for _, value := range values {
		slice = append(slice, value)
	}
	return slice
}
