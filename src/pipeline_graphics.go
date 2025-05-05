package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) createGraphicsPipeline() error {
	// Load vertex shader
	vertShaderBytes, err := fileSystem.ReadFile("shaders/vert.spv")
	if err != nil {
		return err
	}

	vertShader, _, err := app.device.CreateShaderModule(nil, core1_0.ShaderModuleCreateInfo{
		Code: bytesToBytecode(vertShaderBytes),
	})
	if err != nil {
		return err
	}
	defer vertShader.Destroy(nil)

	// Load fragment shader
	fragShaderBytes, err := fileSystem.ReadFile("shaders/frag.spv")
	if err != nil {
		return err
	}

	fragShader, _, err := app.device.CreateShaderModule(nil, core1_0.ShaderModuleCreateInfo{
		Code: bytesToBytecode(fragShaderBytes),
	})
	if err != nil {
		return err
	}
	defer fragShader.Destroy(nil)

	vertexInput := &core1_0.PipelineVertexInputStateCreateInfo{
		VertexBindingDescriptions:   getVertexBindingDescription(),
		VertexAttributeDescriptions: getVertexAttributeDescriptions(),
	}

	inputAssembly := &core1_0.PipelineInputAssemblyStateCreateInfo{
		Topology:               core1_0.PrimitiveTopologyTriangleList,
		PrimitiveRestartEnable: false,
	}

	vertStage := core1_0.PipelineShaderStageCreateInfo{
		Stage:  core1_0.StageVertex,
		Module: vertShader,
		Name:   "main",
	}

	fragStage := core1_0.PipelineShaderStageCreateInfo{
		Stage:  core1_0.StageFragment,
		Module: fragShader,
		Name:   "main",
	}

	viewport := &core1_0.PipelineViewportStateCreateInfo{
		Viewports: []core1_0.Viewport{
			{
				X:        0,
				Y:        0,
				Width:    float32(app.swapchainExtent.Width),
				Height:   float32(app.swapchainExtent.Height),
				MinDepth: 0,
				MaxDepth: 1,
			},
		},
		Scissors: []core1_0.Rect2D{
			{
				Offset: core1_0.Offset2D{X: 0, Y: 0},
				Extent: app.swapchainExtent,
			},
		},
	}

	rasterization := &core1_0.PipelineRasterizationStateCreateInfo{
		DepthClampEnable:        false,
		RasterizerDiscardEnable: false,

		PolygonMode: core1_0.PolygonModeFill,
		CullMode:    core1_0.CullModeBack,
		FrontFace:   core1_0.FrontFaceCounterClockwise,

		DepthBiasEnable: false,

		LineWidth: 1.0,
	}

	multisample := &core1_0.PipelineMultisampleStateCreateInfo{
		SampleShadingEnable:  false,
		RasterizationSamples: core1_0.Samples1,
		MinSampleShading:     1.0,
	}

	depthStencil := &core1_0.PipelineDepthStencilStateCreateInfo{
		DepthTestEnable:  true,
		DepthWriteEnable: true,
		DepthCompareOp:   core1_0.CompareOpLess,
	}

	colorBlend := &core1_0.PipelineColorBlendStateCreateInfo{
		LogicOpEnabled: false,
		LogicOp:        core1_0.LogicOpCopy,

		BlendConstants: [4]float32{0, 0, 0, 0},
		Attachments: []core1_0.PipelineColorBlendAttachmentState{
			{
				BlendEnabled:   false,
				ColorWriteMask: core1_0.ColorComponentRed | core1_0.ColorComponentGreen | core1_0.ColorComponentBlue | core1_0.ColorComponentAlpha,
			},
		},
	}

	app.pipelineLayout, _, err = app.device.CreatePipelineLayout(nil, core1_0.PipelineLayoutCreateInfo{
		SetLayouts: []core1_0.DescriptorSetLayout{
			app.descriptorSetLayout,
		},
	})

	pipelines, _, err := app.device.CreateGraphicsPipelines(nil, nil, []core1_0.GraphicsPipelineCreateInfo{
		{
			Stages: []core1_0.PipelineShaderStageCreateInfo{
				vertStage,
				fragStage,
			},
			VertexInputState:   vertexInput,
			InputAssemblyState: inputAssembly,
			ViewportState:      viewport,
			RasterizationState: rasterization,
			MultisampleState:   multisample,
			DepthStencilState:  depthStencil,
			ColorBlendState:    colorBlend,
			Layout:             app.pipelineLayout,
			RenderPass:         app.renderPass,
			Subpass:            0,
			BasePipelineIndex:  -1,
		},
	})
	if err != nil {
		return err
	}
	app.graphicsPipeline = pipelines[0]

	return nil
}

func bytesToBytecode(b []byte) []uint32 {
	byteCode := make([]uint32, len(b)/4)
	for i := 0; i < len(byteCode); i++ {
		byteIndex := i * 4
		byteCode[i] = 0
		byteCode[i] |= uint32(b[byteIndex])
		byteCode[i] |= uint32(b[byteIndex+1]) << 8
		byteCode[i] |= uint32(b[byteIndex+2]) << 16
		byteCode[i] |= uint32(b[byteIndex+3]) << 24
	}

	return byteCode
}
