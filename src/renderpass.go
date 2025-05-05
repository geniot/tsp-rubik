package main

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/extensions/v2/khr_swapchain"
)

func (app *HelloTriangleApplication) createRenderPass() error {
	depthFormat, err := app.findDepthFormat()
	if err != nil {
		return err
	}

	renderPass, _, err := app.device.CreateRenderPass(nil, core1_0.RenderPassCreateInfo{
		Attachments: []core1_0.AttachmentDescription{
			{
				Format:         app.swapchainImageFormat,
				Samples:        core1_0.Samples1,
				LoadOp:         core1_0.AttachmentLoadOpClear,
				StoreOp:        core1_0.AttachmentStoreOpStore,
				StencilLoadOp:  core1_0.AttachmentLoadOpDontCare,
				StencilStoreOp: core1_0.AttachmentStoreOpDontCare,
				InitialLayout:  core1_0.ImageLayoutUndefined,
				FinalLayout:    khr_swapchain.ImageLayoutPresentSrc,
			},
			{
				Format:         depthFormat,
				Samples:        core1_0.Samples1,
				LoadOp:         core1_0.AttachmentLoadOpClear,
				StoreOp:        core1_0.AttachmentStoreOpDontCare,
				StencilLoadOp:  core1_0.AttachmentLoadOpDontCare,
				StencilStoreOp: core1_0.AttachmentStoreOpDontCare,
				InitialLayout:  core1_0.ImageLayoutUndefined,
				FinalLayout:    core1_0.ImageLayoutDepthStencilAttachmentOptimal,
			},
		},
		Subpasses: []core1_0.SubpassDescription{
			{
				PipelineBindPoint: core1_0.PipelineBindPointGraphics,
				ColorAttachments: []core1_0.AttachmentReference{
					{
						Attachment: 0,
						Layout:     core1_0.ImageLayoutColorAttachmentOptimal,
					},
				},
				DepthStencilAttachment: &core1_0.AttachmentReference{
					Attachment: 1,
					Layout:     core1_0.ImageLayoutDepthStencilAttachmentOptimal,
				},
			},
		},
		SubpassDependencies: []core1_0.SubpassDependency{
			{
				SrcSubpass: core1_0.SubpassExternal,
				DstSubpass: 0,

				SrcStageMask:  core1_0.PipelineStageColorAttachmentOutput | core1_0.PipelineStageEarlyFragmentTests,
				SrcAccessMask: 0,

				DstStageMask:  core1_0.PipelineStageColorAttachmentOutput | core1_0.PipelineStageEarlyFragmentTests,
				DstAccessMask: core1_0.AccessColorAttachmentWrite | core1_0.AccessDepthStencilAttachmentWrite,
			},
		},
	})
	if err != nil {
		return err
	}

	app.renderPass = renderPass

	return nil
}
