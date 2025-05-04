package main

import vk "github.com/vulkan-go/vulkan"

type platform struct {
	basePlatform

	surface       vk.Surface
	debugCallback vk.DebugReportCallback
}

func (p *platform) Surface() vk.Surface {
	return p.surface
}

func (p *platform) Destroy() {
	if p.device != nil {
		vk.DeviceWaitIdle(p.device)
	}
	p.context.destroy()
	p.context = nil
	if p.surface != vk.NullSurface {
		vk.DestroySurface(p.instance, p.surface, nil)
		p.surface = vk.NullSurface
	}
	if p.device != nil {
		vk.DestroyDevice(p.device, nil)
		p.device = nil
	}
	if p.debugCallback != vk.NullDebugReportCallback {
		vk.DestroyDebugReportCallback(p.instance, p.debugCallback, nil)
	}
	if p.instance != nil {
		vk.DestroyInstance(p.instance, nil)
		p.instance = nil
	}
}
