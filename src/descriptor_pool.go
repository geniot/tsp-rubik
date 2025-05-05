package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) createDescriptorPool() error {
	var err error
	app.descriptorPool, _, err = app.device.CreateDescriptorPool(nil, core1_0.DescriptorPoolCreateInfo{
		MaxSets: len(app.swapchainImages),
		PoolSizes: []core1_0.DescriptorPoolSize{
			{
				Type:            core1_0.DescriptorTypeUniformBuffer,
				DescriptorCount: len(app.swapchainImages),
			},
			{
				Type:            core1_0.DescriptorTypeCombinedImageSampler,
				DescriptorCount: len(app.swapchainImages),
			},
		},
	})
	return err
}
