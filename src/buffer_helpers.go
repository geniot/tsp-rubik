package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) copyBuffer(srcBuffer core1_0.Buffer, dstBuffer core1_0.Buffer, size int) error {
	buffer, err := app.beginSingleTimeCommands()
	if err != nil {
		return err
	}

	err = buffer.CmdCopyBuffer(srcBuffer, dstBuffer, []core1_0.BufferCopy{
		{
			SrcOffset: 0,
			DstOffset: 0,
			Size:      size,
		},
	})
	if err != nil {
		return err
	}

	return app.endSingleTimeCommands(buffer)
}

func (app *HelloTriangleApplication) endSingleTimeCommands(buffer core1_0.CommandBuffer) error {
	_, err := buffer.End()
	if err != nil {
		return err
	}

	_, err = app.graphicsQueue.Submit(nil, []core1_0.SubmitInfo{
		{
			CommandBuffers: []core1_0.CommandBuffer{buffer},
		},
	})

	if err != nil {
		return err
	}

	_, err = app.graphicsQueue.WaitIdle()
	if err != nil {
		return err
	}

	app.device.FreeCommandBuffers([]core1_0.CommandBuffer{buffer})
	return nil
}

func (app *HelloTriangleApplication) beginSingleTimeCommands() (core1_0.CommandBuffer, error) {
	buffers, _, err := app.device.AllocateCommandBuffers(core1_0.CommandBufferAllocateInfo{
		CommandPool:        app.commandPool,
		Level:              core1_0.CommandBufferLevelPrimary,
		CommandBufferCount: 1,
	})
	if err != nil {
		return nil, err
	}

	buffer := buffers[0]
	_, err = buffer.Begin(core1_0.CommandBufferBeginInfo{
		Flags: core1_0.CommandBufferUsageOneTimeSubmit,
	})
	return buffer, err
}

func (app *HelloTriangleApplication) createBuffer(size int, usage core1_0.BufferUsageFlags, properties core1_0.MemoryPropertyFlags) (core1_0.Buffer, core1_0.DeviceMemory, error) {
	buffer, _, err := app.device.CreateBuffer(nil, core1_0.BufferCreateInfo{
		Size:        size,
		Usage:       usage,
		SharingMode: core1_0.SharingModeExclusive,
	})
	if err != nil {
		return nil, nil, err
	}

	memRequirements := buffer.MemoryRequirements()
	memoryTypeIndex, err := app.findMemoryType(memRequirements.MemoryTypeBits, properties)
	if err != nil {
		return buffer, nil, err
	}

	memory, _, err := app.device.AllocateMemory(nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  memRequirements.Size,
		MemoryTypeIndex: memoryTypeIndex,
	})
	if err != nil {
		return buffer, nil, err
	}

	_, err = buffer.BindBufferMemory(memory, 0)
	return buffer, memory, err
}
