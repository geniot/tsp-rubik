package main

import (
	"encoding/binary"
	"github.com/vkngwrapper/core/v2/core1_0"
)

func (app *HelloTriangleApplication) createVertexBuffer() error {
	var err error
	bufferSize := binary.Size(vertices)

	stagingBuffer, stagingBufferMemory, err := app.createBuffer(bufferSize, core1_0.BufferUsageTransferSrc, core1_0.MemoryPropertyHostVisible|core1_0.MemoryPropertyHostCoherent)
	if stagingBuffer != nil {
		defer stagingBuffer.Destroy(nil)
	}
	if stagingBufferMemory != nil {
		defer stagingBufferMemory.Free(nil)
	}

	if err != nil {
		return err
	}

	err = writeData(stagingBufferMemory, 0, vertices)
	if err != nil {
		return err
	}

	app.vertexBuffer, app.vertexBufferMemory, err = app.createBuffer(bufferSize, core1_0.BufferUsageTransferDst|core1_0.BufferUsageVertexBuffer, core1_0.MemoryPropertyDeviceLocal)
	if err != nil {
		return err
	}

	return app.copyBuffer(stagingBuffer, app.vertexBuffer, bufferSize)
}
