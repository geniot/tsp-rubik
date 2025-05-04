package main

import vk "github.com/vulkan-go/vulkan"

type VulkanMode uint32

const (
	VulkanNone VulkanMode = (1 << iota) >> 1
	VulkanCompute
	VulkanGraphics
	VulkanPresent
)

func (v VulkanMode) Has(mode VulkanMode) bool {
	return v&mode != 0
}

var (
	DefaultVulkanAppVersion = vk.MakeVersion(1, 0, 0)
	DefaultVulkanAPIVersion = vk.MakeVersion(1, 0, 0)
	DefaultVulkanMode       = VulkanCompute | VulkanGraphics | VulkanPresent
)

// SwapchainDimensions describes the size and format of the swapchain.
type SwapchainDimensions struct {
	// Width of the swapchain.
	Width uint32
	// Height of the swapchain.
	Height uint32
	// Format is the pixel format of the swapchain.
	Format vk.Format
}

type BaseApplication struct {
	context *Context
}

func (app *BaseApplication) Context() *Context {
	return app.context
}

func (app *BaseApplication) VulkanInit(ctx *Context) error {
	app.context = ctx
	return nil
}

func (app *BaseApplication) VulkanAPIVersion() vk.Version {
	return vk.Version(vk.MakeVersion(1, 0, 0))
}

func (app *BaseApplication) VulkanAppVersion() vk.Version {
	return vk.Version(vk.MakeVersion(1, 0, 0))
}

func (app *BaseApplication) VulkanAppName() string {
	return "base"
}

func (app *BaseApplication) VulkanMode() VulkanMode {
	return DefaultVulkanMode
}

func (app *BaseApplication) VulkanSurface(instance vk.Instance) vk.Surface {
	return vk.NullSurface
}

func (app *BaseApplication) VulkanInstanceExtensions() []string {
	return nil
}

func (app *BaseApplication) VulkanDeviceExtensions() []string {
	return nil
}

func (app *BaseApplication) VulkanDebug() bool {
	return false
}
