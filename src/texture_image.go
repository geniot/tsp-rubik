package main

import (
	"bufio"
	"bytes"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/core1_0"
	vkngmath "github.com/vkngwrapper/math"
	"image/png"
)

func (app *HelloTriangleApplication) createTextureImageView() error {
	var err error
	app.textureImageView, err = app.createImageView(app.textureImage, core1_0.FormatR8G8B8A8SRGB, core1_0.ImageAspectColor)
	return err
}

func (app *HelloTriangleApplication) createCubeColors() ([]byte, error) {
	// https://www.schemecolor.com/rubik-cube-colors.php
	var (
		//black     = vkngmath.Vec4[float64]{X: 0, Y: 0, Z: 0, W: 255}
		green  = vkngmath.Vec4[float64]{X: 0, Y: 155, Z: 72, W: 255}
		red    = vkngmath.Vec4[float64]{X: 185, Y: 0, Z: 0, W: 255}
		blue   = vkngmath.Vec4[float64]{X: 0, Y: 69, Z: 173, W: 255}
		orange = vkngmath.Vec4[float64]{X: 255, Y: 89, Z: 0, W: 255}
		//white     = vkngmath.Vec4[float64]{X: 255, Y: 255, Z: 255, W: 255}
		//yellow    = vkngmath.Vec4[float64]{X: 255, Y: 213, Z: 0, W: 255}
		//allColors = []vkngmath.Vec4[float64]{black, green, red, blue, orange, white, yellow}
		//allColors = []vkngmath.Vec4[float64]{green, red, blue, orange}
	)
	var (
		width  = 100
		height = 100
	)
	bytesBuffer := new(bytes.Buffer)
	dc := gg.NewContext(width*2, height*2)

	app.draw(green, dc, 0, 0, width, height)
	app.draw(red, dc, 0, 1, width, height)
	app.draw(blue, dc, 1, 0, width, height)
	app.draw(orange, dc, 1, 1, width, height)

	w := bufio.NewWriter(bytesBuffer)
	if err := dc.EncodePNG(w); err != nil {
		return nil, err
	}
	if err := dc.SavePNG("out.png"); err != nil {
		return nil, err
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

func (app *HelloTriangleApplication) draw(color vkngmath.Vec4[float64], dc *gg.Context, x, y, width, height int) {
	dc.DrawRectangle(float64(x*width), float64(y*height), float64(width), float64(height))
	dc.SetRGBA255(int(color.X), int(color.Y), int(color.Z), int(color.W))
	dc.Fill()
}

func (app *HelloTriangleApplication) createTextureImage() error {
	//Put image data into staging buffer
	imageBytes, err := app.createCubeColors()
	if err != nil {
		return err
	}

	decodedImage, err := png.Decode(bytes.NewBuffer(imageBytes))
	if err != nil {
		return err
	}
	imageBounds := decodedImage.Bounds()
	imageDims := imageBounds.Size()
	imageSize := imageDims.X * imageDims.Y * 4

	stagingBuffer, stagingMemory, err := app.createBuffer(imageSize, core1_0.BufferUsageTransferSrc, core1_0.MemoryPropertyHostVisible|core1_0.MemoryPropertyHostCoherent)
	if err != nil {
		return err
	}

	var pixelData []byte

	for y := imageBounds.Min.Y; y < imageBounds.Max.Y; y++ {
		for x := imageBounds.Min.X; x < imageBounds.Max.Y; x++ {
			r, g, b, a := decodedImage.At(x, y).RGBA()
			pixelData = append(pixelData, byte(r), byte(g), byte(b), byte(a))
		}
	}

	err = writeData(stagingMemory, 0, pixelData)
	if err != nil {
		return err
	}

	//Create final image
	app.textureImage, app.textureImageMemory, err = app.createImage(imageDims.X, imageDims.Y, core1_0.FormatR8G8B8A8SRGB, core1_0.ImageTilingOptimal, core1_0.ImageUsageTransferDst|core1_0.ImageUsageSampled, core1_0.MemoryPropertyDeviceLocal)
	if err != nil {
		return err
	}

	// Copy staging to final
	err = app.transitionImageLayout(app.textureImage, core1_0.FormatR8G8B8A8SRGB, core1_0.ImageLayoutUndefined, core1_0.ImageLayoutTransferDstOptimal)
	if err != nil {
		return err
	}
	err = app.copyBufferToImage(stagingBuffer, app.textureImage, imageDims.X, imageDims.Y)
	if err != nil {
		return err
	}
	err = app.transitionImageLayout(app.textureImage, core1_0.FormatR8G8B8A8SRGB, core1_0.ImageLayoutTransferDstOptimal, core1_0.ImageLayoutShaderReadOnlyOptimal)
	if err != nil {
		return err
	}

	stagingBuffer.Destroy(nil)
	stagingMemory.Free(nil)

	return nil
}

func (app *HelloTriangleApplication) transitionImageLayout(image core1_0.Image, format core1_0.Format, oldLayout core1_0.ImageLayout, newLayout core1_0.ImageLayout) error {
	buffer, err := app.beginSingleTimeCommands()
	if err != nil {
		return err
	}

	var sourceStage, destStage core1_0.PipelineStageFlags
	var sourceAccess, destAccess core1_0.AccessFlags

	if oldLayout == core1_0.ImageLayoutUndefined && newLayout == core1_0.ImageLayoutTransferDstOptimal {
		sourceAccess = 0
		destAccess = core1_0.AccessTransferWrite
		sourceStage = core1_0.PipelineStageTopOfPipe
		destStage = core1_0.PipelineStageTransfer
	} else if oldLayout == core1_0.ImageLayoutTransferDstOptimal && newLayout == core1_0.ImageLayoutShaderReadOnlyOptimal {
		sourceAccess = core1_0.AccessTransferWrite
		destAccess = core1_0.AccessShaderRead
		sourceStage = core1_0.PipelineStageTransfer
		destStage = core1_0.PipelineStageFragmentShader
	} else {
		return errors.Errorf("unexpected layout transition: %s -> %s", oldLayout, newLayout)
	}

	err = buffer.CmdPipelineBarrier(sourceStage, destStage, 0, nil, nil, []core1_0.ImageMemoryBarrier{
		{
			OldLayout:           oldLayout,
			NewLayout:           newLayout,
			SrcQueueFamilyIndex: -1,
			DstQueueFamilyIndex: -1,
			Image:               image,
			SubresourceRange: core1_0.ImageSubresourceRange{
				AspectMask:     core1_0.ImageAspectColor,
				BaseMipLevel:   0,
				LevelCount:     1,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			SrcAccessMask: sourceAccess,
			DstAccessMask: destAccess,
		},
	})
	if err != nil {
		return err
	}

	return app.endSingleTimeCommands(buffer)
}

func (app *HelloTriangleApplication) createImage(
	width, height int,
	format core1_0.Format,
	tiling core1_0.ImageTiling,
	usage core1_0.ImageUsageFlags,
	memoryProperties core1_0.MemoryPropertyFlags,
) (core1_0.Image, core1_0.DeviceMemory, error) {

	image, _, err := app.device.CreateImage(nil, core1_0.ImageCreateInfo{
		ImageType: core1_0.ImageType2D,
		Extent: core1_0.Extent3D{
			Width:  width,
			Height: height,
			Depth:  1,
		},
		MipLevels:     1,
		ArrayLayers:   1,
		Format:        format,
		Tiling:        tiling,
		InitialLayout: core1_0.ImageLayoutUndefined,
		Usage:         usage,
		SharingMode:   core1_0.SharingModeExclusive,
		Samples:       core1_0.Samples1,
	})
	if err != nil {
		return nil, nil, err
	}

	memReqs := image.MemoryRequirements()
	memoryIndex, err := app.findMemoryType(memReqs.MemoryTypeBits, memoryProperties)
	if err != nil {
		return nil, nil, err
	}

	imageMemory, _, err := app.device.AllocateMemory(nil, core1_0.MemoryAllocateInfo{
		AllocationSize:  memReqs.Size,
		MemoryTypeIndex: memoryIndex,
	})

	_, err = image.BindImageMemory(imageMemory, 0)
	if err != nil {
		return nil, nil, err
	}

	return image, imageMemory, nil
}

func (app *HelloTriangleApplication) copyBufferToImage(buffer core1_0.Buffer, image core1_0.Image, width, height int) error {
	cmdBuffer, err := app.beginSingleTimeCommands()
	if err != nil {
		return err
	}

	err = cmdBuffer.CmdCopyBufferToImage(buffer, image, core1_0.ImageLayoutTransferDstOptimal, []core1_0.BufferImageCopy{
		{
			BufferOffset:      0,
			BufferRowLength:   0,
			BufferImageHeight: 0,

			ImageSubresource: core1_0.ImageSubresourceLayers{
				AspectMask:     core1_0.ImageAspectColor,
				MipLevel:       0,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			ImageOffset: core1_0.Offset3D{X: 0, Y: 0, Z: 0},
			ImageExtent: core1_0.Extent3D{Width: width, Height: height, Depth: 1},
		},
	})
	if err != nil {
		return err
	}

	return app.endSingleTimeCommands(cmdBuffer)
}
