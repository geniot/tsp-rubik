package main

import (
	"github.com/loov/hrtime"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/extensions/v2/khr_swapchain"
	vkngmath "github.com/vkngwrapper/math"
	"math"
)

func (app *HelloTriangleApplication) drawFrame() error {
	fences := []core1_0.Fence{app.inFlightFence[app.currentFrame]}

	_, err := app.device.WaitForFences(true, common.NoTimeout, fences)
	if err != nil {
		return err
	}

	imageIndex, res, err := app.swapchain.AcquireNextImage(common.NoTimeout, app.imageAvailableSemaphore[app.currentFrame], nil)
	if res == khr_swapchain.VKErrorOutOfDate {
		return app.recreateSwapChain()
	} else if err != nil {
		return err
	}

	if app.imagesInFlightFence[imageIndex] != nil {
		_, err := app.imagesInFlightFence[imageIndex].Wait(common.NoTimeout)
		if err != nil {
			return err
		}
	}
	app.imagesInFlightFence[imageIndex] = app.inFlightFence[app.currentFrame]

	_, err = app.device.ResetFences(fences)
	if err != nil {
		return err
	}

	err = app.updateUniformBuffer(imageIndex)
	if err != nil {
		return err
	}

	_, err = app.graphicsQueue.Submit(app.inFlightFence[app.currentFrame], []core1_0.SubmitInfo{
		{
			WaitSemaphores:   []core1_0.Semaphore{app.imageAvailableSemaphore[app.currentFrame]},
			WaitDstStageMask: []core1_0.PipelineStageFlags{core1_0.PipelineStageColorAttachmentOutput},
			CommandBuffers:   []core1_0.CommandBuffer{app.commandBuffers[imageIndex]},
			SignalSemaphores: []core1_0.Semaphore{app.renderFinishedSemaphore[app.currentFrame]},
		},
	})
	if err != nil {
		return err
	}

	res, err = app.swapchainExtension.QueuePresent(app.presentQueue, khr_swapchain.PresentInfo{
		WaitSemaphores: []core1_0.Semaphore{app.renderFinishedSemaphore[app.currentFrame]},
		Swapchains:     []khr_swapchain.Swapchain{app.swapchain},
		ImageIndices:   []int{imageIndex},
	})
	if res == khr_swapchain.VKErrorOutOfDate || res == khr_swapchain.VKSuboptimal {
		return app.recreateSwapChain()
	} else if err != nil {
		return err
	}

	app.currentFrame = (app.currentFrame + 1) % MaxFramesInFlight // 0|1

	return nil
}

func (app *HelloTriangleApplication) updateUniformBuffer(currentImage int) error {
	currentTime := hrtime.Now().Seconds()
	speed := -4.0 //rotation speed: higher-slower, also: use minus to rotate clockwise

	timePeriod := math.Mod(currentTime, speed*2)

	ubo := UniformBufferObject{}
	ubo.Model.SetRotationZ(timePeriod * math.Pi / speed)
	ubo.View.SetLookAt(
		&vkngmath.Vec3[float32]{X: 2, Y: 2, Z: 2},
		&vkngmath.Vec3[float32]{X: 0, Y: 0, Z: 0},
		&vkngmath.Vec3[float32]{X: 0, Y: 0, Z: 1},
	)
	aspectRatio := float32(app.swapchainExtent.Width) / float32(app.swapchainExtent.Height)

	near := float32(0.1)
	far := float32(10.0)
	fovy := math.Pi / 10.0 //how close is the camera: increase to get closer

	ubo.Proj.SetPerspective(fovy, aspectRatio, near, far)

	err := writeData(app.uniformBuffersMemory[currentImage], 0, &ubo)
	return err
}
