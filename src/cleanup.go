package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vkngwrapper/core/v2/core1_0"
)

func (app *HelloTriangleApplication) cleanupSwapChain() {
	if app.depthImageView != nil {
		app.depthImageView.Destroy(nil)
		app.depthImageView = nil
	}

	if app.depthImage != nil {
		app.depthImage.Destroy(nil)
		app.depthImage = nil
	}

	if app.depthImageMemory != nil {
		app.depthImageMemory.Free(nil)
		app.depthImageMemory = nil
	}

	for _, framebuffer := range app.swapchainFramebuffers {
		framebuffer.Destroy(nil)
	}
	app.swapchainFramebuffers = []core1_0.Framebuffer{}

	if len(app.commandBuffers) > 0 {
		app.device.FreeCommandBuffers(app.commandBuffers)
		app.commandBuffers = []core1_0.CommandBuffer{}
	}

	if app.graphicsPipeline != nil {
		app.graphicsPipeline.Destroy(nil)
		app.graphicsPipeline = nil
	}

	if app.pipelineLayout != nil {
		app.pipelineLayout.Destroy(nil)
		app.pipelineLayout = nil
	}

	if app.renderPass != nil {
		app.renderPass.Destroy(nil)
		app.renderPass = nil
	}

	for _, imageView := range app.swapchainImageViews {
		imageView.Destroy(nil)
	}
	app.swapchainImageViews = []core1_0.ImageView{}

	if app.swapchain != nil {
		app.swapchain.Destroy(nil)
		app.swapchain = nil
	}

	for i := 0; i < len(app.uniformBuffers); i++ {
		app.uniformBuffers[i].Destroy(nil)
	}
	app.uniformBuffers = app.uniformBuffers[:0]

	for i := 0; i < len(app.uniformBuffersMemory); i++ {
		app.uniformBuffersMemory[i].Free(nil)
	}
	app.uniformBuffersMemory = app.uniformBuffersMemory[:0]

	app.descriptorPool.Destroy(nil)
}

func (app *HelloTriangleApplication) cleanup() {
	app.cleanupSwapChain()

	if app.textureSampler != nil {
		app.textureSampler.Destroy(nil)
	}

	if app.textureImageView != nil {
		app.textureImageView.Destroy(nil)
	}

	if app.textureImage != nil {
		app.textureImage.Destroy(nil)
	}

	if app.textureImageMemory != nil {
		app.textureImageMemory.Free(nil)
	}

	if app.descriptorSetLayout != nil {
		app.descriptorSetLayout.Destroy(nil)
	}

	if app.indexBuffer != nil {
		app.indexBuffer.Destroy(nil)
	}

	if app.indexBufferMemory != nil {
		app.indexBufferMemory.Free(nil)
	}

	if app.vertexBuffer != nil {
		app.vertexBuffer.Destroy(nil)
	}

	if app.vertexBufferMemory != nil {
		app.vertexBufferMemory.Free(nil)
	}

	for _, fence := range app.inFlightFence {
		fence.Destroy(nil)
	}

	for _, semaphore := range app.renderFinishedSemaphore {
		semaphore.Destroy(nil)
	}

	for _, semaphore := range app.imageAvailableSemaphore {
		semaphore.Destroy(nil)
	}

	if app.commandPool != nil {
		app.commandPool.Destroy(nil)
	}

	if app.device != nil {
		app.device.Destroy(nil)
	}

	if app.debugMessenger != nil {
		app.debugMessenger.Destroy(nil)
	}

	if app.surface != nil {
		app.surface.Destroy(nil)
	}

	if app.instance != nil {
		app.instance.Destroy(nil)
	}

	if app.window != nil {
		app.window.Destroy()
	}
	sdl.Quit()
}
