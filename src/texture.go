package main

import vk "github.com/vulkan-go/vulkan"

type Texture struct {
	sampler vk.Sampler

	image       vk.Image
	imageLayout vk.ImageLayout

	memAlloc *vk.MemoryAllocateInfo
	mem      vk.DeviceMemory
	view     vk.ImageView

	texWidth  int32
	texHeight int32
}

func (t *Texture) Destroy(dev vk.Device) {
	vk.DestroyImageView(dev, t.view, nil)
	vk.FreeMemory(dev, t.mem, nil)
	vk.DestroyImage(dev, t.image, nil)
	vk.DestroySampler(dev, t.sampler, nil)
}

func (t *Texture) DestroyImage(dev vk.Device) {
	vk.FreeMemory(dev, t.mem, nil)
	vk.DestroyImage(dev, t.image, nil)
}
