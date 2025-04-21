package main

import (
	"github.com/veandco/go-sdl2/sdl"
	as "github.com/vulkan-go/asche"
	vk "github.com/vulkan-go/vulkan"
	"log"
)

type VulkanApplication struct {
	*SpinningCube
	debugEnabled bool
	windowHandle *sdl.Window
}

func (a *VulkanApplication) VulkanSurface(instance vk.Instance) (surface vk.Surface) {
	surfPtr, err := a.windowHandle.VulkanCreateSurface(instance)
	if err != nil {
		log.Println("vulkan error:", err)
		return vk.NullSurface
	}
	surf := vk.SurfaceFromPointer(uintptr(surfPtr))
	return surf
}

func (a *VulkanApplication) VulkanAppName() string {
	return "VulkanCube"
}

func (a *VulkanApplication) VulkanLayers() []string {
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

func (a *VulkanApplication) VulkanDebug() bool {
	return false // a.debugEnabled
}

func (a *VulkanApplication) VulkanDeviceExtensions() []string {
	return []string{
		"VK_KHR_swapchain",
	}
}

func (a *VulkanApplication) VulkanSwapchainDimensions() *as.SwapchainDimensions {
	return &as.SwapchainDimensions{
		Width: 1280, Height: 720, Format: vk.FormatB8g8r8a8Unorm,
	}
}

func (a *VulkanApplication) VulkanInstanceExtensions() []string {
	extensions := a.windowHandle.VulkanGetInstanceExtensions()
	if a.debugEnabled {
		extensions = append(extensions, "VK_EXT_debug_report")
	}
	return extensions
}

func NewVulkanApplication(debugEnabled bool) *VulkanApplication {
	return &VulkanApplication{
		SpinningCube: NewSpinningCube(0.5),

		debugEnabled: debugEnabled,
	}
}
