package main

import (
	"github.com/vkngwrapper/extensions/v2/ext_debug_utils"
	"log"
)

func (app *HelloTriangleApplication) setupDebugMessenger() error {
	if !enableValidationLayers {
		return nil
	}

	var err error
	debugLoader := ext_debug_utils.CreateExtensionFromInstance(app.instance)
	app.debugMessenger, _, err = debugLoader.CreateDebugUtilsMessenger(app.instance, nil, app.debugMessengerOptions())
	if err != nil {
		return err
	}

	return nil
}

func (app *HelloTriangleApplication) logDebug(msgType ext_debug_utils.DebugUtilsMessageTypeFlags, severity ext_debug_utils.DebugUtilsMessageSeverityFlags, data *ext_debug_utils.DebugUtilsMessengerCallbackData) bool {
	log.Printf("[%s %s] - %s", severity, msgType, data.Message)
	return false
}

func (app *HelloTriangleApplication) debugMessengerOptions() ext_debug_utils.DebugUtilsMessengerCreateInfo {
	return ext_debug_utils.DebugUtilsMessengerCreateInfo{
		MessageSeverity: ext_debug_utils.SeverityError | ext_debug_utils.SeverityWarning,
		MessageType:     ext_debug_utils.TypeGeneral | ext_debug_utils.TypeValidation | ext_debug_utils.TypePerformance,
		UserCallback:    app.logDebug,
	}
}
