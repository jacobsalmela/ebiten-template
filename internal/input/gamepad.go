// Copyright 2015 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package input

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type VirtualGamepadButton int

const (
	VirtualGamepadButtonLeft VirtualGamepadButton = iota
	VirtualGamepadButtonRight
	VirtualGamepadButtonDown
	VirtualGamepadButtonButtonA
	VirtualGamepadButtonButtonB
)

var VirtualGamepadButtons = []VirtualGamepadButton{
	VirtualGamepadButtonLeft,
	VirtualGamepadButtonRight,
	VirtualGamepadButtonDown,
	VirtualGamepadButtonButtonA,
	VirtualGamepadButtonButtonB,
}

func (v VirtualGamepadButton) StandardGamepadButton() ebiten.StandardGamepadButton {
	switch v {
	case VirtualGamepadButtonLeft:
		return ebiten.StandardGamepadButtonLeftLeft
	case VirtualGamepadButtonRight:
		return ebiten.StandardGamepadButtonLeftRight
	case VirtualGamepadButtonDown:
		return ebiten.StandardGamepadButtonLeftBottom
	case VirtualGamepadButtonButtonA:
		return ebiten.StandardGamepadButtonRightBottom
	case VirtualGamepadButtonButtonB:
		return ebiten.StandardGamepadButtonRightRight
	default:
		panic("not reached")
	}
}

const axisThreshold = 0.75

type axis struct {
	id       int
	positive bool
}

type GamepadConfig struct {
	gamepadID            ebiten.GamepadID
	gamepadIDInitialized bool

	current         VirtualGamepadButton
	buttons         map[VirtualGamepadButton]ebiten.GamepadButton
	axes            map[VirtualGamepadButton]axis
	assignedButtons map[ebiten.GamepadButton]struct{}
	assignedAxes    map[axis]struct{}

	defaultAxesValues map[int]float64
}

func (c *GamepadConfig) SetGamepadID(id ebiten.GamepadID) {
	c.gamepadID = id
	c.gamepadIDInitialized = true
}

func (c *GamepadConfig) ResetGamepadID() {
	c.gamepadID = 0
	c.gamepadIDInitialized = false
}

func (c *GamepadConfig) IsGamepadIDInitialized() bool {
	return c.gamepadIDInitialized
}

func (c *GamepadConfig) NeedsConfiguration() bool {
	return !ebiten.IsStandardGamepadLayoutAvailable(c.gamepadID)
}

func (c *GamepadConfig) initializeIfNeeded() {
	if !c.gamepadIDInitialized {
		panic("not reached")
	}

	if ebiten.IsStandardGamepadLayoutAvailable(c.gamepadID) {
		return
	}

	if c.buttons == nil {
		c.buttons = map[VirtualGamepadButton]ebiten.GamepadButton{}
	}
	if c.axes == nil {
		c.axes = map[VirtualGamepadButton]axis{}
	}
	if c.assignedButtons == nil {
		c.assignedButtons = map[ebiten.GamepadButton]struct{}{}
	}
	if c.assignedAxes == nil {
		c.assignedAxes = map[axis]struct{}{}
	}

	// Set default values.
	// It is assumed that all axes are not pressed here.
	//
	// These default values are used to detect if an axis is actually pressed.
	// For example, on PS4 controllers, L2/R2's axes valuse can be -1.0.
	if c.defaultAxesValues == nil {
		c.defaultAxesValues = map[int]float64{}
		na := ebiten.GamepadAxisCount(c.gamepadID)
		for a := 0; a < na; a++ {
			c.defaultAxesValues[a] = ebiten.GamepadAxisValue(c.gamepadID, a)
		}
	}
}

func (c *GamepadConfig) Reset() {
	c.buttons = nil
	c.axes = nil
	c.assignedButtons = nil
	c.assignedAxes = nil
}

// Scan scans the current input state and assigns the given virtual gamepad button b
// to the current (pysical) pressed buttons of the gamepad.
func (c *GamepadConfig) Scan(b VirtualGamepadButton) bool {
	if !c.gamepadIDInitialized {
		panic("not reached")
	}

	c.initializeIfNeeded()

	delete(c.buttons, b)
	delete(c.axes, b)

	ebn := ebiten.GamepadButton(ebiten.GamepadButtonCount(c.gamepadID))
	for eb := ebiten.GamepadButton(0); eb < ebn; eb++ {
		if _, ok := c.assignedButtons[eb]; ok {
			continue
		}
		if inpututil.IsGamepadButtonJustPressed(c.gamepadID, eb) {
			c.buttons[b] = eb
			c.assignedButtons[eb] = struct{}{}
			return true
		}
	}

	na := ebiten.GamepadAxisCount(c.gamepadID)
	for a := 0; a < na; a++ {
		v := ebiten.GamepadAxisValue(c.gamepadID, a)
		const delta = 0.25

		// Check |v| < 1.0 because there is a bug that a button returns
		// an axis value wrongly and the value may be over 1 on some platforms.
		if axisThreshold <= v && v <= 1.0 &&
			(v < c.defaultAxesValues[a]-delta || c.defaultAxesValues[a]+delta < v) {
			if _, ok := c.assignedAxes[axis{a, true}]; !ok {
				c.axes[b] = axis{a, true}
				c.assignedAxes[axis{a, true}] = struct{}{}
				return true
			}
		}
		if -1.0 <= v && v <= -axisThreshold &&
			(v < c.defaultAxesValues[a]-delta || c.defaultAxesValues[a]+delta < v) {
			if _, ok := c.assignedAxes[axis{a, false}]; !ok {
				c.axes[b] = axis{a, false}
				c.assignedAxes[axis{a, false}] = struct{}{}
				return true
			}
		}
	}

	return false
}

// IsButtonPressed reports whether the given virtual button b is pressed.
func (c *GamepadConfig) IsButtonPressed(b VirtualGamepadButton) bool {
	if !c.gamepadIDInitialized {
		panic("not reached")
	}

	if ebiten.IsStandardGamepadLayoutAvailable(c.gamepadID) {
		if ebiten.IsStandardGamepadButtonPressed(c.gamepadID, b.StandardGamepadButton()) {
			return true
		}

		const threshold = 0.7
		switch b {
		case VirtualGamepadButtonLeft:
			return ebiten.StandardGamepadAxisValue(c.gamepadID, ebiten.StandardGamepadAxisLeftStickHorizontal) < -threshold
		case VirtualGamepadButtonRight:
			return ebiten.StandardGamepadAxisValue(c.gamepadID, ebiten.StandardGamepadAxisLeftStickHorizontal) > threshold
		case VirtualGamepadButtonDown:
			return ebiten.StandardGamepadAxisValue(c.gamepadID, ebiten.StandardGamepadAxisLeftStickVertical) > threshold
		}
		return false
	}

	c.initializeIfNeeded()

	bb, ok := c.buttons[b]
	if ok {
		return ebiten.IsGamepadButtonPressed(c.gamepadID, bb)
	}

	a, ok := c.axes[b]
	if ok {
		v := ebiten.GamepadAxisValue(c.gamepadID, a.id)
		if a.positive {
			return axisThreshold <= v && v <= 1.0
		}
		return -1.0 <= v && v <= -axisThreshold
	}
	return false
}

// IsButtonJustPressed reports whether the given virtual button b started to be pressed now.
func (c *GamepadConfig) IsButtonJustPressed(b VirtualGamepadButton) bool {
	if !c.gamepadIDInitialized {
		panic("not reached")
	}

	if ebiten.IsStandardGamepadLayoutAvailable(c.gamepadID) {
		return inpututil.IsStandardGamepadButtonJustPressed(c.gamepadID, b.StandardGamepadButton())
	}

	c.initializeIfNeeded()

	bb, ok := c.buttons[b]
	if ok {
		return inpututil.IsGamepadButtonJustPressed(c.gamepadID, bb)
	}
	return false
}

// Name returns the pysical button's name for the given virtual button.
func (c *GamepadConfig) ButtonName(b VirtualGamepadButton) string {
	if !c.gamepadIDInitialized {
		panic("not reached")
	}

	c.initializeIfNeeded()

	bb, ok := c.buttons[b]
	if ok {
		return fmt.Sprintf("Button %d", bb)
	}

	a, ok := c.axes[b]
	if ok {
		if a.positive {
			return fmt.Sprintf("Axis %d+", a.id)
		}
		return fmt.Sprintf("Axis %d-", a.id)
	}

	return ""
}
