package main

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/extensions/v2/khr_portability_subset"
)

func (app *HelloTriangleApplication) createLogicalDevice() error {
	indices, err := app.findQueueFamilies(app.physicalDevice)
	if err != nil {
		return err
	}

	uniqueQueueFamilies := []int{*indices.GraphicsFamily}
	if uniqueQueueFamilies[0] != *indices.PresentFamily {
		uniqueQueueFamilies = append(uniqueQueueFamilies, *indices.PresentFamily)
	}

	var queueFamilyOptions []core1_0.DeviceQueueCreateInfo
	queuePriority := float32(1.0)
	for _, queueFamily := range uniqueQueueFamilies {
		queueFamilyOptions = append(queueFamilyOptions, core1_0.DeviceQueueCreateInfo{
			QueueFamilyIndex: queueFamily,
			QueuePriorities:  []float32{queuePriority},
		})
	}

	var extensionNames []string
	extensionNames = append(extensionNames, deviceExtensions...)

	// Makes this example compatible with vulkan portability, necessary to run on mobile & mac
	extensions, _, err := app.physicalDevice.EnumerateDeviceExtensionProperties()
	if err != nil {
		return err
	}

	_, supported := extensions[khr_portability_subset.ExtensionName]
	if supported {
		extensionNames = append(extensionNames, khr_portability_subset.ExtensionName)
	}

	app.device, _, err = app.physicalDevice.CreateDevice(nil, core1_0.DeviceCreateInfo{
		QueueCreateInfos: queueFamilyOptions,
		EnabledFeatures: &core1_0.PhysicalDeviceFeatures{
			SamplerAnisotropy: true,
		},
		EnabledExtensionNames: extensionNames,
	})
	if err != nil {
		return err
	}

	app.graphicsQueue = app.device.GetQueue(*indices.GraphicsFamily, 0)
	app.presentQueue = app.device.GetQueue(*indices.PresentFamily, 0)
	return nil
}
