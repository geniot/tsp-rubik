package main

import (
	"github.com/vkngwrapper/extensions/v2/khr_surface"
	vkng_sdl2 "github.com/vkngwrapper/integrations/sdl2/v2"
)

func (app *HelloTriangleApplication) createSurface() error {
	surfaceLoader := khr_surface.CreateExtensionFromInstance(app.instance)

	surface, err := vkng_sdl2.CreateSurface(app.instance, surfaceLoader, app.window)
	if err != nil {
		return err
	}

	app.surface = surface
	return nil
}
