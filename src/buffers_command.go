package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) createCommandBuffers() error {

	buffers, _, err := app.device.AllocateCommandBuffers(core1_0.CommandBufferAllocateInfo{
		CommandPool:        app.commandPool,
		Level:              core1_0.CommandBufferLevelPrimary,
		CommandBufferCount: len(app.swapchainImages),
	})
	if err != nil {
		return err
	}
	app.commandBuffers = buffers

	for bufferIdx, buffer := range buffers {
		_, err = buffer.Begin(core1_0.CommandBufferBeginInfo{})
		if err != nil {
			return err
		}

		err = buffer.CmdBeginRenderPass(core1_0.SubpassContentsInline,
			core1_0.RenderPassBeginInfo{
				RenderPass:  app.renderPass,
				Framebuffer: app.swapchainFramebuffers[bufferIdx],
				RenderArea: core1_0.Rect2D{
					Offset: core1_0.Offset2D{X: 0, Y: 0},
					Extent: app.swapchainExtent,
				},
				ClearValues: []core1_0.ClearValue{
					core1_0.ClearValueFloat{0, 0, 0, 1},
					core1_0.ClearValueDepthStencil{Depth: 1.0, Stencil: 0},
				},
			})
		if err != nil {
			return err
		}

		buffer.CmdBindPipeline(core1_0.PipelineBindPointGraphics, app.graphicsPipeline)
		buffer.CmdBindVertexBuffers(0, []core1_0.Buffer{app.vertexBuffer}, []int{0})
		buffer.CmdBindIndexBuffer(app.indexBuffer, 0, core1_0.IndexTypeUInt16)
		buffer.CmdBindDescriptorSets(core1_0.PipelineBindPointGraphics, app.pipelineLayout, 0, []core1_0.DescriptorSet{
			app.descriptorSets[bufferIdx],
		}, nil)
		buffer.CmdDrawIndexed(len(indices), 1, 0, 0, 0)
		buffer.CmdEndRenderPass()

		_, err = buffer.End()
		if err != nil {
			return err
		}
	}

	return nil
}
