package Inputs

import (
	EIM "github.com/Try-si/EIM"
	"github.com/hajimehoshi/ebiten/v2"
)

var inputManager *InputManager

type InputManager struct {
	im           *EIM.InputManager
	previousKeys map[int]bool
}

func Init(configPath string) {
	inputManager = &InputManager{}
	inputManager.im = EIM.NewInputManager(configPath)
	inputManager.previousKeys = make(map[int]bool)
	for _, key := range inputManager.im.Bindings {
		inputManager.previousKeys[key] = false
	}
}

func BindingIsPressed(key string) bool {
	return inputManager.im.IsKeyPressed(key)
}

func IsBindingJustPressed(binding string) bool {
	key, ok := inputManager.im.Bindings[binding]
	if !ok {
		return false
	}
	currentPressed := BindingIsPressed(binding)
	wasPressed := inputManager.previousKeys[key]
	return currentPressed && !wasPressed
}

func IsBindingJustReleased(binding string) bool {
	key, ok := inputManager.im.Bindings[binding]
	if !ok {
		return false
	}
	currentPressed := BindingIsPressed(binding)
	wasPressed := inputManager.previousKeys[key]
	return !currentPressed && wasPressed
}

func Update() {
	if inputManager == nil {
		return
	}
	if inputManager.previousKeys == nil {
		inputManager.previousKeys = make(map[int]bool)
	}
	for key, _ := range inputManager.previousKeys {
		inputManager.previousKeys[key] = ebiten.IsKeyPressed(ebiten.Key(key))
	}
	for _, key := range inputManager.im.Bindings {
		if _, exists := inputManager.previousKeys[key]; !exists {
			inputManager.previousKeys[key] = ebiten.IsKeyPressed(ebiten.Key(key))
		}
	}
}

func BindingExists(key string) bool {
	return inputManager.im.HasBinding(key)
}

func SetBinding(key string, value int) {
	inputManager.im.SetBinding(key, value)
}

func RemoveBinding(key string) {
	inputManager.im.RemoveBinding(key)
}

func AddBinding(key string, value int) {
	inputManager.im.AddBinding(key, value)
}
