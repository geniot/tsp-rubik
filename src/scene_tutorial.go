package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TutorialScene struct {
	docTextures [16]rl.Texture2D
	docPointer  int
}

func (ts *TutorialScene) ShouldExit() bool {
	return rl.IsKeyPressed(rl.KeyEscape) || (rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode))
}

func (ts *TutorialScene) Draw(camera *rl.Camera) {

}

func NewTutorialScene() *TutorialScene {
	//for i := 0; i < len(docTextures); i++ {
	//	textureBytes := orPanicRes(mediaList.ReadFile("media/doc" + strconv.Itoa(i) + ".png"))
	//	docTextures[i] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", textureBytes, int32(len(textureBytes))))
	//}
	return &TutorialScene{}
}
