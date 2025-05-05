package main

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

func (app *HelloTriangleApplication) createUniformBuffers() error {
	bufferSize := int(unsafe.Sizeof(UniformBufferObject{}))

	for i := 0; i < len(app.swapchainImages); i++ {
		buffer, memory, err := app.createBuffer(bufferSize, core1_0.BufferUsageUniformBuffer, core1_0.MemoryPropertyHostVisible|core1_0.MemoryPropertyHostCoherent)
		if err != nil {
			return err
		}

		app.uniformBuffers = append(app.uniformBuffers, buffer)
		app.uniformBuffersMemory = append(app.uniformBuffersMemory, memory)
	}

	return nil
}
