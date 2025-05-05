package main

import (
	"encoding/binary"
	"github.com/vkngwrapper/core/v2/core1_0"
)

func (app *HelloTriangleApplication) createIndexBuffer() error {
	bufferSize := binary.Size(indices)

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

	err = writeData(stagingBufferMemory, 0, indices)
	if err != nil {
		return err
	}

	app.indexBuffer, app.indexBufferMemory, err = app.createBuffer(bufferSize, core1_0.BufferUsageTransferDst|core1_0.BufferUsageIndexBuffer, core1_0.MemoryPropertyDeviceLocal)
	if err != nil {
		return err
	}

	return app.copyBuffer(stagingBuffer, app.indexBuffer, bufferSize)
}
