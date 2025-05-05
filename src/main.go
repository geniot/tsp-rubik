package main

import (
	"embed"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	vk "github.com/vulkan-go/vulkan"
)

var (
	//go:embed media/*
	mediaList embed.FS
)

func GetResource(fileName string) []byte {
	file, _ := mediaList.Open("media/" + fileName)
	stat, _ := file.Stat()
	size := stat.Size()
	buf := make([]byte, size)
	file.Read(buf)
	return buf
}

func init() {
	runtime.LockOSThread()
	log.SetFlags(log.Lshortfile)
}

func main() {
	orPanic(sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS))
	defer sdl.Quit()

	orPanic(sdl.VulkanLoadLibrary(""))
	defer sdl.VulkanUnloadLibrary()

	vk.SetGetInstanceProcAddr(sdl.VulkanGetVkGetInstanceProcAddr())
	orPanic(vk.Init())

	app := NewCubeApplication(true, 1.0)
	reqDim := app.VulkanSwapchainDimensions()
	window, err := sdl.CreateWindow("VulkanCube (SDL2)",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(reqDim.Width), int32(reqDim.Height),
		sdl.WINDOW_VULKAN)
	orPanic(err)
	app.sdlWindow = window

	// creates a new platform, also initializes Vulkan context in the app
	platform, err := NewPlatform(app)
	orPanic(err)

	dim := app.Context().SwapchainDimensions()
	log.Printf("Initialized %s with %+v swapchain", app.VulkanAppName(), dim)

	// some sync logic
	doneC := make(chan struct{}, 2)
	exitC := make(chan struct{}, 2)

	fpsDelay := time.Second / 60
	fpsTicker := time.NewTicker(fpsDelay)
	start := time.Now()
	frames := 0
_MainLoop:
	for {
		select {
		case <-exitC:
			fmt.Printf("FPS: %.2f\n", float64(frames)/time.Now().Sub(start).Seconds())
			app.Destroy()
			platform.Destroy()
			window.Destroy()
			fpsTicker.Stop()
			doneC <- struct{}{}
			return
		case <-fpsTicker.C:
			frames++
			var event sdl.Event
			for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.KeyboardEvent:
					if t.Keysym.Sym == sdl.K_ESCAPE {
						exitC <- struct{}{}
						continue _MainLoop
					}
				case *sdl.QuitEvent:
					exitC <- struct{}{}
					continue _MainLoop
				}
			}
			app.NextFrame()

			imageIdx, outdated, err := app.Context().AcquireNextImage()
			orPanic(err)
			if outdated {
				imageIdx, _, err = app.Context().AcquireNextImage()
				orPanic(err)
			}
			_, err = app.Context().PresentImage(imageIdx)
			orPanic(err)
		}
	}
}

//func orPanic(err interface{}) {
//	switch v := err.(type) {
//	case error:
//		if v != nil {
//			panic(err)
//		}
//	case vk.Result:
//		if err := vk.Error(v); err != nil {
//			panic(err)
//		}
//	case bool:
//		if !v {
//			panic("condition failed: != true")
//		}
//	}
//}
