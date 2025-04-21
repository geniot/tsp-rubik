package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/tevino/abool/v2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	as "github.com/vulkan-go/asche"
	vk "github.com/vulkan-go/vulkan"
	"github.com/xlab/closer"
	"log"
	"os"
	"runtime"
	"runtime/debug"
)

type Application struct {
	*SpinningCube
	settings           *Settings
	debugEnabled       bool
	sdlWindow          *sdl.Window
	sdlGameController  *sdl.GameController
	platform           as.Platform
	font               *ttf.Font
	joysticks          [16]*sdl.Joystick
	pressedKeysCodes   mapset.Set[sdl.Keycode]
	pressedButtonCodes mapset.Set[ButtonCode]
	isRunning          *abool.AtomicBool
}

func (app *Application) Start(args []string) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Unhandled error: %v\n", r)
			log.Println("Stack trace:")
			debug.PrintStack()
			os.Exit(-1)
		}
	}()

	orPanic(sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS | sdl.INIT_JOYSTICK | sdl.INIT_GAMECONTROLLER))
	defer sdl.Quit()

	orPanic(ttf.Init())
	app.font = orPanicRes(ttf.OpenFontRW(LoadMediaFile("pixelberry.ttf"), 1, 20))

	sdl.JoystickEventState(sdl.ENABLE)
	for i := 0; i < sdl.NumJoysticks(); i++ {
		if sdl.IsGameController(i) {
			app.sdlGameController = sdl.GameControllerOpen(i)
		}
	}
	println(runtime.GOARCH)
	if runtime.GOARCH == "arm64" { //most likely it's a TSP device
		orPanic(app.sdlGameController != nil)
	}

	orPanic(sdl.VulkanLoadLibrary(""))
	defer sdl.VulkanUnloadLibrary()

	vk.SetGetInstanceProcAddr(sdl.VulkanGetVkGetInstanceProcAddr())
	orPanic(vk.Init())
	defer closer.Close()

	//reqDim := app.VulkanSwapchainDimensions()

	app.sdlWindow = orPanicRes(sdl.CreateWindow("VulkanCube (SDL2)",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		TSP_SCREEN_WIDTH, TSP_SCREEN_HEIGHT,
		sdl.WINDOW_VULKAN))

	// creates a new platform, also initializes Vulkan context in the app
	app.platform = orPanicRes(as.NewPlatform(app))

	dim := app.Context().SwapchainDimensions()
	log.Printf("Initialized %s with %+v swapchain", app.VulkanAppName(), dim)

	app.isRunning.Set()
	for app.isRunning.IsSet() {
		app.UpdateEvents()
		app.UpdatePhysics()
		app.UpdateView()
	}
	app.releaseResources()
}

func (app *Application) releaseResources() {
	app.settings.Save(app.sdlWindow)
	sdl.VulkanUnloadLibrary()
	app.sdlGameController.Close()
	app.font.Close()
	ttf.Quit()
	sdl.Quit()
}

func (app *Application) UpdateView() {
	imageIdx, outdated, err := app.Context().AcquireNextImage()
	orPanic(err)
	if outdated {
		imageIdx, _, err = app.Context().AcquireNextImage()
		orPanic(err)
	}
	_, err = app.Context().PresentImage(imageIdx)
	orPanic(err)
}

func (app *Application) Stop() {
	app.isRunning.UnSet()
}

func (app *Application) UpdatePhysics() {
	x := If(app.pressedButtonCodes.Contains(BUTTON_CODE_UP), -1.0, If(app.pressedButtonCodes.Contains(BUTTON_CODE_DOWN), 1.0, 0.0))
	y := If(app.pressedButtonCodes.Contains(BUTTON_CODE_LEFT), -1.0, If(app.pressedButtonCodes.Contains(BUTTON_CODE_RIGHT), 1.0, 0.0))
	app.NextFrame(float32(x), float32(y))
	if app.pressedKeysCodes.Contains(sdl.K_q) || (app.pressedButtonCodes.Contains(BUTTON_CODE_MENU) && app.pressedButtonCodes.Contains(BUTTON_CODE_START)) {
		app.Stop()
	}
}

func (app *Application) UpdateEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.JoyAxisEvent:
			// Convert the value to a -1.0 - 1.0 range
			//value := float64(t.Value) / 32768.0
			break

		case *sdl.ControllerButtonEvent:
			if t.State == sdl.PRESSED {
				println(t.Button)
				app.pressedButtonCodes.Add(t.Button)
			} else {
				app.pressedButtonCodes.Remove(t.Button)
			}
			break

		case *sdl.JoyDeviceAddedEvent:
			// Open joystick for use
			app.joysticks[int(t.Which)] = sdl.JoystickOpen(int(t.Which))
			if app.joysticks[int(t.Which)] != nil {
				fmt.Println("Joystick", t.Which, "connected")
			}
			break
		case *sdl.JoyDeviceRemovedEvent:
			if joystick := app.joysticks[int(t.Which)]; joystick != nil {
				joystick.Close()
			}
			fmt.Println("Joystick", t.Which, "disconnected")
			break

		case *sdl.KeyboardEvent:
			if t.Repeat > 0 {
				break
			}
			if t.State == sdl.PRESSED {
				app.pressedKeysCodes.Add(t.Keysym.Sym)
			} else { // if t.State == sdl.RELEASED {
				app.pressedKeysCodes.Remove(t.Keysym.Sym)
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				app.settings.SaveWindowState(app.sdlWindow)
			}
			break

		case *sdl.QuitEvent:
			app.Stop()
			break
		}
	}
}

func (a *Application) VulkanSurface(instance vk.Instance) (surface vk.Surface) {
	surfPtr, err := a.sdlWindow.VulkanCreateSurface(instance)
	if err != nil {
		log.Println("vulkan error:", err)
		return vk.NullSurface
	}
	surf := vk.SurfaceFromPointer(uintptr(surfPtr))
	return surf
}

func (a *Application) VulkanAppName() string {
	return "VulkanCube"
}

func (a *Application) VulkanLayers() []string {
	return []string{
		// "VK_LAYER_GOOGLE_threading",
		// "VK_LAYER_LUNARG_parameter_validation",
		// "VK_LAYER_LUNARG_object_tracker",
		// "VK_LAYER_LUNARG_core_validation",
		// "VK_LAYER_LUNARG_api_dump",
		// "VK_LAYER_LUNARG_swapchain",
		// "VK_LAYER_GOOGLE_unique_objects",
	}
}

func (a *Application) VulkanDebug() bool {
	return false // a.debugEnabled
}

func (a *Application) VulkanDeviceExtensions() []string {
	return []string{
		"VK_KHR_swapchain",
	}
}

//func (a *Application) VulkanSwapchainDimensions() *as.SwapchainDimensions {
//	return &as.SwapchainDimensions{
//		Width: 1280, Height: 720, Format: vk.FormatB8g8r8a8Unorm,
//	}
//}

func (a *Application) VulkanInstanceExtensions() []string {
	extensions := a.sdlWindow.VulkanGetInstanceExtensions()
	if a.debugEnabled {
		extensions = append(extensions, "VK_EXT_debug_report")
	}
	return extensions
}

func NewApplication(debugEnabled bool) *Application {
	return &Application{
		SpinningCube:       NewSpinningCube(0.5),
		debugEnabled:       debugEnabled,
		pressedKeysCodes:   mapset.NewSet[sdl.Keycode](),
		pressedButtonCodes: mapset.NewSet[ButtonCode](),
		isRunning:          abool.New(),
		settings:           NewSettings(),
	}
}
