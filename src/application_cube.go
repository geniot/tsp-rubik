package main

import (
	"github.com/veandco/go-sdl2/sdl"
	vk "github.com/vulkan-go/vulkan"
	"log"
)

type CubeApplication struct {
	*SpinningCubeApplication
	debugEnabled bool
	sdlWindow    *sdl.Window
}

func (a *CubeApplication) VulkanSurface(instance vk.Instance) (surface vk.Surface) {
	surfPtr, err := a.sdlWindow.VulkanCreateSurface(instance)
	if err != nil {
		log.Println("vulkan error:", err)
		return vk.NullSurface
	}
	surf := vk.SurfaceFromPointer(uintptr(surfPtr))
	return surf
}

func (a *CubeApplication) VulkanAppName() string {
	return "VulkanCube"
}

func (a *CubeApplication) VulkanLayers() []string {
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

func (a *CubeApplication) VulkanDebug() bool {
	return false // a.debugEnabled
}

func (a *CubeApplication) VulkanDeviceExtensions() []string {
	return []string{
		"VK_KHR_swapchain",
	}
}

func (a *CubeApplication) VulkanSwapchainDimensions() *SwapchainDimensions {
	return &SwapchainDimensions{
		Width: 1280, Height: 720, Format: vk.FormatB8g8r8a8Unorm,
	}
}

func (a *CubeApplication) VulkanInstanceExtensions() []string {
	extensions := a.sdlWindow.VulkanGetInstanceExtensions()
	if a.debugEnabled {
		extensions = append(extensions, "VK_EXT_debug_report")
	}
	return extensions
}

func NewCubeApplication(debugEnabled bool) *CubeApplication {
	return &CubeApplication{
		SpinningCubeApplication: NewSpinningCube(1.0),
		debugEnabled:            debugEnabled,
	}
}
