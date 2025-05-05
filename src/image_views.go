package main

import "github.com/vkngwrapper/core/v2/core1_0"

func (app *HelloTriangleApplication) createImageViews() error {
	images, _, err := app.swapchain.SwapchainImages()
	if err != nil {
		return err
	}
	app.swapchainImages = images

	var imageViews []core1_0.ImageView
	for _, image := range images {
		view, err := app.createImageView(image, app.swapchainImageFormat, core1_0.ImageAspectColor)
		if err != nil {
			return err
		}

		imageViews = append(imageViews, view)
	}
	app.swapchainImageViews = imageViews

	return nil
}

func (app *HelloTriangleApplication) createDepthResources() error {
	depthFormat, err := app.findDepthFormat()
	if err != nil {
		return err
	}

	app.depthImage, app.depthImageMemory, err = app.createImage(app.swapchainExtent.Width,
		app.swapchainExtent.Height,
		depthFormat,
		core1_0.ImageTilingOptimal,
		core1_0.ImageUsageDepthStencilAttachment,
		core1_0.MemoryPropertyDeviceLocal)
	if err != nil {
		return err
	}
	app.depthImageView, err = app.createImageView(app.depthImage, depthFormat, core1_0.ImageAspectDepth)
	return err
}

func (app *HelloTriangleApplication) createImageView(image core1_0.Image, format core1_0.Format, aspect core1_0.ImageAspectFlags) (core1_0.ImageView, error) {
	imageView, _, err := app.device.CreateImageView(nil, core1_0.ImageViewCreateInfo{
		Image:    image,
		ViewType: core1_0.ImageViewType2D,
		Format:   format,
		SubresourceRange: core1_0.ImageSubresourceRange{
			AspectMask:     aspect,
			BaseMipLevel:   0,
			LevelCount:     1,
			BaseArrayLayer: 0,
			LayerCount:     1,
		},
	})
	return imageView, err
}
