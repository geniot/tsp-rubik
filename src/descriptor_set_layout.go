package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) createDescriptorSetLayout() error {
	var err error
	app.descriptorSetLayout, _, err = app.device.CreateDescriptorSetLayout(nil, core1_0.DescriptorSetLayoutCreateInfo{
		Bindings: []core1_0.DescriptorSetLayoutBinding{
			{
				Binding:         0,
				DescriptorType:  core1_0.DescriptorTypeUniformBuffer,
				DescriptorCount: 1,

				StageFlags: core1_0.StageVertex,
			},
			{
				Binding:         1,
				DescriptorType:  core1_0.DescriptorTypeCombinedImageSampler,
				DescriptorCount: 1,

				StageFlags: core1_0.StageFragment,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
