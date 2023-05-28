package input

// Dir represents a direction.
type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

// String returns a string representing the direction.
func (d Dir) String() string {
	switch d {
	case DirUp:
		return "Up"
	case DirRight:
		return "Right"
	case DirDown:
		return "Down"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

// Vector returns a [-1, 1] value for each axis.
func (d Dir) Vector() (x, y int) {
	switch d {
	case DirUp:
		return 0, -1
	case DirRight:
		return 1, 0
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	}
	panic("not reach")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func vecToDir(dx, dy int) (Dir, bool) {
	if abs(dx) < 4 && abs(dy) < 4 {
		return 0, false
	}
	if abs(dx) < abs(dy) {
		if dy < 0 {
			return DirUp, true
		}
		return DirDown, true
	}
	if dx < 0 {
		return DirLeft, true
	}
	return DirRight, true
}

// Dir returns a currently pressed direction.
// Dir returns false if no direction key is pressed.
func (i *Input) Dir() (int, bool) {
	// if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
	// 	return DirUp, true
	// }
	// if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
	// 	return DirLeft, true
	// }
	// if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
	// 	return DirRight, true
	// }
	// if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
	// 	return DirDown, true
	// }
	if i.mouseState == mouseStateSettled {
		return i.mouseDir, true
	}
	if i.touchState == touchStateSettled {
		return i.touchDir, true
	}
	return 0, false
}
