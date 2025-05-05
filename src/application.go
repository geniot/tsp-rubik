package main

import (
	"embed"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vkngwrapper/core/v2"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/extensions/v2/ext_debug_utils"
	"github.com/vkngwrapper/extensions/v2/khr_surface"
	"github.com/vkngwrapper/extensions/v2/khr_swapchain"
	vkngmath "github.com/vkngwrapper/math"
	"unsafe"
)

//go:embed shaders images
var fileSystem embed.FS

const MaxFramesInFlight = 2

var validationLayers = []string{
	//"VK_LAYER_KHRONOS_validation"
}
var deviceExtensions = []string{khr_swapchain.ExtensionName}

const enableValidationLayers = true

type QueueFamilyIndices struct {
	GraphicsFamily *int
	PresentFamily  *int
}

func (i *QueueFamilyIndices) IsComplete() bool {
	return i.GraphicsFamily != nil && i.PresentFamily != nil
}

type SwapChainSupportDetails struct {
	Capabilities *khr_surface.SurfaceCapabilities
	Formats      []khr_surface.SurfaceFormat
	PresentModes []khr_surface.PresentMode
}

type Vertex struct {
	Position vkngmath.Vec3[float32]
	Color    vkngmath.Vec3[float32]
	TexCoord vkngmath.Vec2[float32]
}

type UniformBufferObject struct {
	Model vkngmath.Mat4x4[float32]
	View  vkngmath.Mat4x4[float32]
	Proj  vkngmath.Mat4x4[float32]
}

func getVertexBindingDescription() []core1_0.VertexInputBindingDescription {
	v := Vertex{}
	return []core1_0.VertexInputBindingDescription{
		{
			Binding:   0,
			Stride:    int(unsafe.Sizeof(v)),
			InputRate: core1_0.VertexInputRateVertex,
		},
	}
}

func getVertexAttributeDescriptions() []core1_0.VertexInputAttributeDescription {
	v := Vertex{}
	return []core1_0.VertexInputAttributeDescription{
		{
			Binding:  0,
			Location: 0,
			Format:   core1_0.FormatR32G32B32SignedFloat,
			Offset:   int(unsafe.Offsetof(v.Position)),
		},
		{
			Binding:  0,
			Location: 1,
			Format:   core1_0.FormatR32G32B32SignedFloat,
			Offset:   int(unsafe.Offsetof(v.Color)),
		},
		{
			Binding:  0,
			Location: 2,
			Format:   core1_0.FormatR32G32SignedFloat,
			Offset:   int(unsafe.Offsetof(v.TexCoord)),
		},
	}
}

var vertices = []Vertex{
	{Position: vkngmath.Vec3[float32]{X: -0.5, Y: -0.5, Z: 0}, Color: vkngmath.Vec3[float32]{X: 1, Y: 0, Z: 0}, TexCoord: vkngmath.Vec2[float32]{X: 1, Y: 0}},
	{Position: vkngmath.Vec3[float32]{X: 0.5, Y: -0.5, Z: 0}, Color: vkngmath.Vec3[float32]{X: 0, Y: 1, Z: 0}, TexCoord: vkngmath.Vec2[float32]{X: 0, Y: 0}},
	{Position: vkngmath.Vec3[float32]{X: 0.5, Y: 0.5, Z: 0}, Color: vkngmath.Vec3[float32]{X: 0, Y: 0, Z: 1}, TexCoord: vkngmath.Vec2[float32]{X: 0, Y: 1}},
	{Position: vkngmath.Vec3[float32]{X: -0.5, Y: 0.5, Z: 0}, Color: vkngmath.Vec3[float32]{X: 1, Y: 1, Z: 1}, TexCoord: vkngmath.Vec2[float32]{X: 1, Y: 1}},

	{Position: vkngmath.Vec3[float32]{X: -0.5, Y: -0.5, Z: -0.5}, Color: vkngmath.Vec3[float32]{X: 1, Y: 0, Z: 0}, TexCoord: vkngmath.Vec2[float32]{X: 0, Y: 0}},
	{Position: vkngmath.Vec3[float32]{X: 0.5, Y: -0.5, Z: -0.5}, Color: vkngmath.Vec3[float32]{X: 0, Y: 1, Z: 0}, TexCoord: vkngmath.Vec2[float32]{X: 1, Y: 0}},
	{Position: vkngmath.Vec3[float32]{X: 0.5, Y: 0.5, Z: -0.5}, Color: vkngmath.Vec3[float32]{X: 0, Y: 0, Z: 1}, TexCoord: vkngmath.Vec2[float32]{X: 1, Y: 1}},
	{Position: vkngmath.Vec3[float32]{X: -0.5, Y: 0.5, Z: -0.5}, Color: vkngmath.Vec3[float32]{X: 1, Y: 1, Z: 1}, TexCoord: vkngmath.Vec2[float32]{X: 0, Y: 1}},
}

var indices = []uint16{
	0, 1, 2, 2, 3, 0,
	4, 5, 6, 6, 7, 4,
}

type HelloTriangleApplication struct {
	window *sdl.Window
	loader core.Loader

	instance       core1_0.Instance
	debugMessenger ext_debug_utils.DebugUtilsMessenger
	surface        khr_surface.Surface

	physicalDevice core1_0.PhysicalDevice
	device         core1_0.Device

	graphicsQueue core1_0.Queue
	presentQueue  core1_0.Queue

	swapchainExtension    khr_swapchain.Extension
	swapchain             khr_swapchain.Swapchain
	swapchainImages       []core1_0.Image
	swapchainImageFormat  core1_0.Format
	swapchainExtent       core1_0.Extent2D
	swapchainImageViews   []core1_0.ImageView
	swapchainFramebuffers []core1_0.Framebuffer

	renderPass          core1_0.RenderPass
	descriptorPool      core1_0.DescriptorPool
	descriptorSets      []core1_0.DescriptorSet
	descriptorSetLayout core1_0.DescriptorSetLayout
	pipelineLayout      core1_0.PipelineLayout
	graphicsPipeline    core1_0.Pipeline

	commandPool    core1_0.CommandPool
	commandBuffers []core1_0.CommandBuffer

	imageAvailableSemaphore []core1_0.Semaphore
	renderFinishedSemaphore []core1_0.Semaphore
	inFlightFence           []core1_0.Fence
	imagesInFlight          []core1_0.Fence
	currentFrame            int
	frameStart              float64

	vertexBuffer       core1_0.Buffer
	vertexBufferMemory core1_0.DeviceMemory
	indexBuffer        core1_0.Buffer
	indexBufferMemory  core1_0.DeviceMemory

	uniformBuffers       []core1_0.Buffer
	uniformBuffersMemory []core1_0.DeviceMemory

	textureImage       core1_0.Image
	textureImageMemory core1_0.DeviceMemory
	textureImageView   core1_0.ImageView
	textureSampler     core1_0.Sampler

	depthImage       core1_0.Image
	depthImageMemory core1_0.DeviceMemory
	depthImageView   core1_0.ImageView
}

func (app *HelloTriangleApplication) Run() error {
	err := app.initWindow()
	if err != nil {
		return err
	}
	err = app.initVulkan()
	if err != nil {
		return err
	}
	defer app.cleanup()
	return app.mainLoop()
}

func (app *HelloTriangleApplication) mainLoop() error {
	rendering := true

appLoop:
	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				break appLoop
			case *sdl.WindowEvent:
				switch e.Event {
				case sdl.WINDOWEVENT_MINIMIZED:
					rendering = false
				case sdl.WINDOWEVENT_RESTORED:
					rendering = true
				case sdl.WINDOWEVENT_RESIZED:
					w, h := app.window.GetSize()
					if w > 0 && h > 0 {
						rendering = true
						app.recreateSwapChain()
					} else {
						rendering = false
					}
				}
			}
		}
		if rendering {
			err := app.drawFrame()
			if err != nil {
				return err
			}
		}
	}

	_, err := app.device.WaitIdle()
	return err
}
