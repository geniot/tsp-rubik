package main

import (
	"errors"
	vk "github.com/vulkan-go/vulkan"
	"log"
	"unsafe"
)

type Platform struct {
	context *Context

	instance vk.Instance
	gpu      vk.PhysicalDevice
	device   vk.Device

	graphicsQueueIndex uint32
	presentQueueIndex  uint32
	presentQueue       vk.Queue
	graphicsQueue      vk.Queue

	gpuProperties    vk.PhysicalDeviceProperties
	memoryProperties vk.PhysicalDeviceMemoryProperties

	surface       vk.Surface
	debugCallback vk.DebugReportCallback
}

func (p *Platform) Destroy() {
	if p.device != nil {
		vk.DeviceWaitIdle(p.device)
	}
	p.context.destroy()
	p.context = nil
	if p.surface != vk.NullSurface {
		vk.DestroySurface(p.instance, p.surface, nil)
		p.surface = vk.NullSurface
	}
	if p.device != nil {
		vk.DestroyDevice(p.device, nil)
		p.device = nil
	}
	if p.debugCallback != vk.NullDebugReportCallback {
		vk.DestroyDebugReportCallback(p.instance, p.debugCallback, nil)
	}
	if p.instance != nil {
		vk.DestroyInstance(p.instance, nil)
		p.instance = nil
	}
}

func (p *Platform) MemoryProperties() vk.PhysicalDeviceMemoryProperties {
	return p.memoryProperties
}

func (p *Platform) PhysicalDeviceProperies() vk.PhysicalDeviceProperties {
	return p.gpuProperties
}

func (p *Platform) PhysicalDevice() vk.PhysicalDevice {
	return p.gpu
}

func (p *Platform) Surface() vk.Surface {
	return vk.NullSurface
}

func (p *Platform) GraphicsQueueFamilyIndex() uint32 {
	return p.graphicsQueueIndex
}

func (p *Platform) PresentQueueFamilyIndex() uint32 {
	return p.presentQueueIndex
}

func (p *Platform) HasSeparatePresentQueue() bool {
	return p.presentQueueIndex != p.graphicsQueueIndex
}

func (p *Platform) GraphicsQueue() vk.Queue {
	return p.graphicsQueue
}

func (p *Platform) PresentQueue() vk.Queue {
	if p.graphicsQueueIndex != p.presentQueueIndex {
		return p.presentQueue
	}
	return p.graphicsQueue
}

func (p *Platform) Instance() vk.Instance {
	return p.instance
}

func (p *Platform) Device() vk.Device {
	return p.device
}

func NewPlatform(app *CubeApplication) (pFace *Platform, err error) {
	// defer checkErr(&err)
	p := Platform{
		context: &Context{
			// TODO: make configurable
			// defines count of slots allocated in swapchain
			frameLag: 3,
		},
	}
	p.context.platform = &p

	// Select instance extensions
	requiredInstanceExtensions := safeStrings(app.VulkanInstanceExtensions())
	actualInstanceExtensions, err := InstanceExtensions()
	orPanic(err)
	instanceExtensions, missing := checkExisting(actualInstanceExtensions, requiredInstanceExtensions)
	if missing > 0 {
		log.Println("vulkan warning: missing", missing, "required instance extensions during init")
	}
	log.Printf("vulkan: enabling %d instance extensions", len(instanceExtensions))

	// Select instance layers
	var validationLayers []string

	requiredValidationLayers := safeStrings(app.VulkanLayers())
	actualValidationLayers, err := ValidationLayers()
	orPanic(err)
	validationLayers, missing = checkExisting(actualValidationLayers, requiredValidationLayers)
	if missing > 0 {
		log.Println("vulkan warning: missing", missing, "required validation layers during init")
	}

	// Create instance
	var instance vk.Instance
	ret := vk.CreateInstance(&vk.InstanceCreateInfo{
		SType: vk.StructureTypeInstanceCreateInfo,
		PApplicationInfo: &vk.ApplicationInfo{
			SType:              vk.StructureTypeApplicationInfo,
			ApiVersion:         uint32(app.VulkanAPIVersion()),
			ApplicationVersion: uint32(app.VulkanAppVersion()),
			PApplicationName:   safeString(app.VulkanAppName()),
			PEngineName:        "vulkango.com\x00",
		},
		EnabledExtensionCount:   uint32(len(instanceExtensions)),
		PpEnabledExtensionNames: instanceExtensions,
		EnabledLayerCount:       uint32(len(validationLayers)),
		PpEnabledLayerNames:     validationLayers,
	}, nil, &instance)
	orPanic(NewError(ret))
	p.instance = instance
	vk.InitInstance(instance)

	if app.VulkanDebug() {
		// Register a debug callback
		ret := vk.CreateDebugReportCallback(instance, &vk.DebugReportCallbackCreateInfo{
			SType:       vk.StructureTypeDebugReportCallbackCreateInfo,
			Flags:       vk.DebugReportFlags(vk.DebugReportErrorBit | vk.DebugReportWarningBit),
			PfnCallback: dbgCallbackFunc,
		}, nil, &p.debugCallback)
		orPanic(NewError(ret))
		log.Println("vulkan: DebugReportCallback enabled by application")
	}

	// Find a suitable GPU
	var gpuCount uint32
	ret = vk.EnumeratePhysicalDevices(p.instance, &gpuCount, nil)
	orPanic(NewError(ret))
	if gpuCount == 0 {
		return nil, errors.New("vulkan error: no GPU devices found")
	}
	gpus := make([]vk.PhysicalDevice, gpuCount)
	ret = vk.EnumeratePhysicalDevices(p.instance, &gpuCount, gpus)
	orPanic(NewError(ret))
	// get the first one, multiple GPUs not supported yet
	p.gpu = gpus[0]
	vk.GetPhysicalDeviceProperties(p.gpu, &p.gpuProperties)
	p.gpuProperties.Deref()
	vk.GetPhysicalDeviceMemoryProperties(p.gpu, &p.memoryProperties)
	p.memoryProperties.Deref()

	// Select device extensions
	requiredDeviceExtensions := safeStrings(app.VulkanDeviceExtensions())
	actualDeviceExtensions, err := DeviceExtensions(p.gpu)
	orPanic(err)
	deviceExtensions, missing := checkExisting(actualDeviceExtensions, requiredDeviceExtensions)
	if missing > 0 {
		log.Println("vulkan warning: missing", missing, "required device extensions during init")
	}
	log.Printf("vulkan: enabling %d device extensions", len(deviceExtensions))

	// Make sure the surface is here if required
	mode := app.VulkanMode()
	if mode.Has(VulkanPresent) { // so, a surface is required and provided
		p.surface = app.VulkanSurface(p.instance)
		if p.surface == vk.NullSurface {
			return nil, errors.New("vulkan error: surface required but not provided")
		}
	}

	// Get queue family properties
	var queueCount uint32
	vk.GetPhysicalDeviceQueueFamilyProperties(p.gpu, &queueCount, nil)
	queueProperties := make([]vk.QueueFamilyProperties, queueCount)
	vk.GetPhysicalDeviceQueueFamilyProperties(p.gpu, &queueCount, queueProperties)
	if queueCount == 0 { // probably should try another GPU
		return nil, errors.New("vulkan error: no queue families found on GPU 0")
	}

	// Find a suitable queue family for the target Vulkan mode
	var graphicsFound bool
	var presentFound bool
	var separateQueue bool
	for i := uint32(0); i < queueCount; i++ {
		var (
			required        vk.QueueFlags
			supportsPresent vk.Bool32
			needsPresent    bool
		)
		if graphicsFound {
			// looking for separate present queue
			separateQueue = true
			vk.GetPhysicalDeviceSurfaceSupport(p.gpu, i, p.surface, &supportsPresent)
			if supportsPresent.B() {
				p.presentQueueIndex = i
				presentFound = true
				break
			}
		}
		if mode.Has(VulkanCompute) {
			required |= vk.QueueFlags(vk.QueueComputeBit)
		}
		if mode.Has(VulkanGraphics) {
			required |= vk.QueueFlags(vk.QueueGraphicsBit)
		}
		if mode.Has(VulkanPresent) {
			needsPresent = true
			vk.GetPhysicalDeviceSurfaceSupport(p.gpu, i, p.surface, &supportsPresent)
		}
		queueProperties[i].Deref()
		if queueProperties[i].QueueFlags&required != 0 {
			if !needsPresent || (needsPresent && supportsPresent.B()) {
				p.graphicsQueueIndex = i
				graphicsFound = true
				break
			} else if needsPresent {
				p.graphicsQueueIndex = i
				graphicsFound = true
				// need present, but this one doesn't support
				// continue lookup
			}
		}
	}
	if separateQueue && !presentFound {
		err := errors.New("vulkan error: could not found separate queue with present capabilities")
		return nil, err
	}
	if !graphicsFound {
		err := errors.New("vulkan error: could not find a suitable queue family for the target Vulkan mode")
		return nil, err
	}

	// Create a Vulkan device
	queueInfos := []vk.DeviceQueueCreateInfo{{
		SType:            vk.StructureTypeDeviceQueueCreateInfo,
		QueueFamilyIndex: p.graphicsQueueIndex,
		QueueCount:       1,
		PQueuePriorities: []float32{1.0},
	}}
	if separateQueue {
		queueInfos = append(queueInfos, vk.DeviceQueueCreateInfo{
			SType:            vk.StructureTypeDeviceQueueCreateInfo,
			QueueFamilyIndex: p.presentQueueIndex,
			QueueCount:       1,
			PQueuePriorities: []float32{1.0},
		})
	}

	var device vk.Device
	ret = vk.CreateDevice(p.gpu, &vk.DeviceCreateInfo{
		SType:                   vk.StructureTypeDeviceCreateInfo,
		QueueCreateInfoCount:    uint32(len(queueInfos)),
		PQueueCreateInfos:       queueInfos,
		EnabledExtensionCount:   uint32(len(deviceExtensions)),
		PpEnabledExtensionNames: deviceExtensions,
		EnabledLayerCount:       uint32(len(validationLayers)),
		PpEnabledLayerNames:     validationLayers,
	}, nil, &device)
	orPanic(NewError(ret))
	p.device = device
	p.context.device = device
	app.VulkanInit(p.context)

	var queue vk.Queue
	vk.GetDeviceQueue(p.device, p.graphicsQueueIndex, 0, &queue)
	p.graphicsQueue = queue

	if mode.Has(VulkanPresent) { // init a swapchain for surface
		if separateQueue {
			var presentQueue vk.Queue
			vk.GetDeviceQueue(p.device, p.presentQueueIndex, 0, &presentQueue)
			p.presentQueue = presentQueue
		}
		p.context.preparePresent()

		dimensions := &SwapchainDimensions{
			// some default preferences here
			Width: 640, Height: 480,
			Format: vk.FormatB8g8r8a8Unorm,
		}
		dimensions = app.VulkanSwapchainDimensions()
		p.context.prepareSwapchain(p.gpu, p.surface, dimensions)
	}
	p.context.SetOnPrepare(app.VulkanContextPrepare)
	p.context.SetOnCleanup(app.VulkanContextCleanup)
	p.context.SetOnInvalidate(app.VulkanContextInvalidate)
	if mode.Has(VulkanPresent) {
		p.context.prepare(false)
	}
	return &p, nil
}

func dbgCallbackFunc(flags vk.DebugReportFlags, objectType vk.DebugReportObjectType,
	object uint64, location uint, messageCode int32, pLayerPrefix string,
	pMessage string, pUserData unsafe.Pointer) vk.Bool32 {

	switch {
	case flags&vk.DebugReportFlags(vk.DebugReportInformationBit) != 0:
		log.Printf("INFORMATION: [%s] Code %d : %s", pLayerPrefix, messageCode, pMessage)
	case flags&vk.DebugReportFlags(vk.DebugReportWarningBit) != 0:
		log.Printf("WARNING: [%s] Code %d : %s", pLayerPrefix, messageCode, pMessage)
	case flags&vk.DebugReportFlags(vk.DebugReportPerformanceWarningBit) != 0:
		log.Printf("PERFORMANCE WARNING: [%s] Code %d : %s", pLayerPrefix, messageCode, pMessage)
	case flags&vk.DebugReportFlags(vk.DebugReportErrorBit) != 0:
		log.Printf("ERROR: [%s] Code %d : %s", pLayerPrefix, messageCode, pMessage)
	case flags&vk.DebugReportFlags(vk.DebugReportDebugBit) != 0:
		log.Printf("DEBUG: [%s] Code %d : %s", pLayerPrefix, messageCode, pMessage)
	default:
		log.Printf("INFORMATION: [%s] Code %d : %s", pLayerPrefix, messageCode, pMessage)
	}
	return vk.Bool32(vk.False)
}
