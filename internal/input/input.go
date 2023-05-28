package input

import "github.com/hajimehoshi/ebiten/v2"

// Input manages the input state including gamepads and keyboards.
// Input represents the current key states.
type Input struct {
	mouseState    mouseState
	mouseInitPosX int
	mouseInitPosY int
	mouseDir      int

	touches       []ebiten.TouchID
	touchState    touchState
	touchID       ebiten.TouchID
	touchInitPosX int
	touchInitPosY int
	touchLastPosX int
	touchLastPosY int
	touchDir      int

	gamepadIDs                 []ebiten.GamepadID
	VirtualGamepadButtonStates map[VirtualGamepadButton]int
	GamepadConfig              GamepadConfig
}
