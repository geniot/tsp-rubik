package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) createSampler() error {
	properties, err := app.physicalDevice.Properties()
	if err != nil {
		return err
	}

	app.textureSampler, _, err = app.device.CreateSampler(nil, core1_0.SamplerCreateInfo{
		MagFilter:    core1_0.FilterLinear,
		MinFilter:    core1_0.FilterLinear,
		AddressModeU: core1_0.SamplerAddressModeRepeat,
		AddressModeV: core1_0.SamplerAddressModeRepeat,
		AddressModeW: core1_0.SamplerAddressModeRepeat,

		AnisotropyEnable: true,
		MaxAnisotropy:    properties.Limits.MaxSamplerAnisotropy,

		BorderColor: core1_0.BorderColorIntOpaqueBlack,

		MipmapMode: core1_0.SamplerMipmapModeLinear,
	})

	return err
}
