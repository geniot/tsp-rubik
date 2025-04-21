package main

import (
	"embed"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"io/fs"
)

var (
	//go:embed media/*
	mediaList embed.FS
)

type SurfTexture struct {
	W int32
	H int32
	T *sdl.Texture
	S *sdl.Surface
}

type ImageDescriptor struct {
	OffsetX     int32
	OffsetY     int32
	Width       int32
	Height      int32
	ResourceKey ResourceKey
}

func LoadMediaFile(fileName string) *sdl.RWops {
	file, _ := mediaList.Open("media/" + fileName)
	return GetResource(file)
}

func GetResource(file fs.File) *sdl.RWops {
	stat, _ := file.Stat()
	size := stat.Size()
	buf := make([]byte, size)
	if _, err := file.Read(buf); err != nil {
		println(err.Error())
	}
	rwOps, _ := sdl.RWFromMem(buf)
	return rwOps
}

func LoadTexture(fileName string, sdlRenderer *sdl.Renderer) *sdl.Texture {
	return LoadSurfTexture(fileName, sdlRenderer).T
}

func LoadSurfTexture(fileName string, sdlRenderer *sdl.Renderer) *SurfTexture {
	surface, err := img.LoadRW(LoadMediaFile(fileName), true)
	if err != nil {
		println(err.Error())
	}
	defer surface.Free()
	txt, err := sdlRenderer.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return &SurfTexture{T: txt, W: surface.W, H: surface.H}
}
