package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	helpFontSize   = int32(20)
	helpWidth      = int32(360)
	helpHeight     = int32(120)
	helpPadding    = int32(10)
	helpLineHeight = int32(20)
)

type TutorialScene struct {
	docTextures [16]rl.Texture2D
	docPointer  int
}

func NewTutorialScene() *TutorialScene {
	//for i := 0; i < len(docTextures); i++ {
	//	textureBytes := orPanicRes(mediaList.ReadFile("media/doc" + strconv.Itoa(i) + ".png"))
	//	docTextures[i] = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", textureBytes, int32(len(textureBytes))))
	//}
	return &TutorialScene{}
}
