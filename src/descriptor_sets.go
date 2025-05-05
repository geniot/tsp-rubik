package main

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

func (app *HelloTriangleApplication) createDescriptorSets() error {
	var allocLayouts []core1_0.DescriptorSetLayout
	for i := 0; i < len(app.swapchainImages); i++ {
		allocLayouts = append(allocLayouts, app.descriptorSetLayout)
	}

	var err error
	app.descriptorSets, _, err = app.device.AllocateDescriptorSets(core1_0.DescriptorSetAllocateInfo{
		DescriptorPool: app.descriptorPool,
		SetLayouts:     allocLayouts,
	})
	if err != nil {
		return err
	}

	for i := 0; i < len(app.swapchainImages); i++ {
		err = app.device.UpdateDescriptorSets([]core1_0.WriteDescriptorSet{
			{
				DstSet:          app.descriptorSets[i],
				DstBinding:      0,
				DstArrayElement: 0,

				DescriptorType: core1_0.DescriptorTypeUniformBuffer,

				BufferInfo: []core1_0.DescriptorBufferInfo{
					{
						Buffer: app.uniformBuffers[i],
						Offset: 0,
						Range:  int(unsafe.Sizeof(UniformBufferObject{})),
					},
				},
			},
			{
				DstSet:          app.descriptorSets[i],
				DstBinding:      1,
				DstArrayElement: 0,

				DescriptorType: core1_0.DescriptorTypeCombinedImageSampler,

				ImageInfo: []core1_0.DescriptorImageInfo{
					{
						ImageView:   app.textureImageView,
						Sampler:     app.textureSampler,
						ImageLayout: core1_0.ImageLayoutShaderReadOnlyOptimal,
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
