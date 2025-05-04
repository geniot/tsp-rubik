package main

import vk "github.com/vulkan-go/vulkan"

type basePlatform struct {
	context *context

	instance vk.Instance
	gpu      vk.PhysicalDevice
	device   vk.Device

	graphicsQueueIndex uint32
	presentQueueIndex  uint32
	presentQueue       vk.Queue
	graphicsQueue      vk.Queue

	gpuProperties    vk.PhysicalDeviceProperties
	memoryProperties vk.PhysicalDeviceMemoryProperties
}

func (p *basePlatform) MemoryProperties() vk.PhysicalDeviceMemoryProperties {
	return p.memoryProperties
}

func (p *basePlatform) PhysicalDeviceProperies() vk.PhysicalDeviceProperties {
	return p.gpuProperties
}

func (p *basePlatform) PhysicalDevice() vk.PhysicalDevice {
	return p.gpu
}

func (p *basePlatform) Surface() vk.Surface {
	return vk.NullSurface
}

func (p *basePlatform) GraphicsQueueFamilyIndex() uint32 {
	return p.graphicsQueueIndex
}

func (p *basePlatform) PresentQueueFamilyIndex() uint32 {
	return p.presentQueueIndex
}

func (p *basePlatform) HasSeparatePresentQueue() bool {
	return p.presentQueueIndex != p.graphicsQueueIndex
}

func (p *basePlatform) GraphicsQueue() vk.Queue {
	return p.graphicsQueue
}

func (p *basePlatform) PresentQueue() vk.Queue {
	if p.graphicsQueueIndex != p.presentQueueIndex {
		return p.presentQueue
	}
	return p.graphicsQueue
}

func (p *basePlatform) Instance() vk.Instance {
	return p.instance
}

func (p *basePlatform) Device() vk.Device {
	return p.device
}
