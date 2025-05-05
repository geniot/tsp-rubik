package main

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

func (app *HelloTriangleApplication) findSupportedFormat(formats []core1_0.Format, tiling core1_0.ImageTiling, features core1_0.FormatFeatureFlags) (core1_0.Format, error) {
	for _, format := range formats {
		props := app.physicalDevice.FormatProperties(format)

		if tiling == core1_0.ImageTilingLinear && (props.LinearTilingFeatures&features) == features {
			return format, nil
		} else if tiling == core1_0.ImageTilingOptimal && (props.OptimalTilingFeatures&features) == features {
			return format, nil
		}
	}

	return 0, errors.Errorf("failed to find supported format for tiling %s, featureset %s", tiling, features)
}

func (app *HelloTriangleApplication) findDepthFormat() (core1_0.Format, error) {
	return app.findSupportedFormat([]core1_0.Format{core1_0.FormatD32SignedFloat, core1_0.FormatD32SignedFloatS8UnsignedInt, core1_0.FormatD24UnsignedNormalizedS8UnsignedInt},
		core1_0.ImageTilingOptimal,
		core1_0.FormatFeatureDepthStencilAttachment)
}

func hasStencilComponent(format core1_0.Format) bool {
	return format == core1_0.FormatD32SignedFloatS8UnsignedInt || format == core1_0.FormatD24UnsignedNormalizedS8UnsignedInt
}

func writeData(memory core1_0.DeviceMemory, offset int, data any) error {
	bufferSize := binary.Size(data)

	memoryPtr, _, err := memory.Map(offset, bufferSize, 0)
	if err != nil {
		return err
	}
	defer memory.Unmap()

	dataBuffer := unsafe.Slice((*byte)(memoryPtr), bufferSize)

	buf := &bytes.Buffer{}
	err = binary.Write(buf, common.ByteOrder, data)
	if err != nil {
		return err
	}

	copy(dataBuffer, buf.Bytes())
	return nil
}

func (app *HelloTriangleApplication) findMemoryType(typeFilter uint32, properties core1_0.MemoryPropertyFlags) (int, error) {
	memProperties := app.physicalDevice.MemoryProperties()
	for i, memoryType := range memProperties.MemoryTypes {
		typeBit := uint32(1 << i)

		if (typeFilter&typeBit) != 0 && (memoryType.PropertyFlags&properties) == properties {
			return i, nil
		}
	}

	return 0, errors.Errorf("failed to find any suitable memory type!")
}
