package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) createFramebuffers() error {
	for _, imageView := range app.swapchainImageViews {
		framebuffer, _, err := app.device.CreateFramebuffer(nil, core1_0.FramebufferCreateInfo{
			RenderPass: app.renderPass,
			Layers:     1,
			Attachments: []core1_0.ImageView{
				imageView,
				app.depthImageView,
			},
			Width:  app.swapchainExtent.Width,
			Height: app.swapchainExtent.Height,
		})
		if err != nil {
			return err
		}

		app.swapchainFramebuffers = append(app.swapchainFramebuffers, framebuffer)
	}

	return nil
}
