package main

/*
#include "lipgloss_types.h"
*/
import "C"

// Position represents a position along a horizontal or vertical axis.
// It's used in situations where an axis is involved, such as alignment,
// joining, placement, and so on.
//
// A value of 0 represents the start (the left or top) and 1 represents
// the end (the right or bottom). 0.5 represents the center.
type Position float64

const (
	Top    Position = C.POS_TOP
	Bottom Position = C.POS_BOTTOM
	Center Position = C.POS_CENTER
	Left   Position = C.POS_LEFT
	Right  Position = C.POS_RIGHT
)

//export PositionTop
func PositionTop() C.float {
	pos := Top
	if err := Validate.Position(float64(pos), "top"); err != nil {
		Log(LogLevelError, "PositionTop validation error: %v", err)
		return 0.0
	}
	Log(LogLevelDebug, "Returning top position: %f", pos)
	return C.float(pos)
}

//export PositionBottom
func PositionBottom() C.float {
	pos := Bottom
	if err := Validate.Position(float64(pos), "bottom"); err != nil {
		Log(LogLevelError, "PositionBottom validation error: %v", err)
		return 1.0
	}
	Log(LogLevelDebug, "Returning bottom position: %f", pos)
	return C.float(pos)
}

//export PositionCenter
func PositionCenter() C.float {
	pos := Center
	if err := Validate.Position(float64(pos), "center"); err != nil {
		Log(LogLevelError, "PositionCenter validation error: %v", err)
		return 0.5
	}
	Log(LogLevelDebug, "Returning center position: %f", pos)
	return C.float(pos)
}

//export PositionLeft
func PositionLeft() C.float {
	pos := Left
	if err := Validate.Position(float64(pos), "left"); err != nil {
		Log(LogLevelError, "PositionLeft validation error: %v", err)
		return 0.0
	}
	Log(LogLevelDebug, "Returning left position: %f", pos)
	return C.float(pos)
}

//export PositionRight
func PositionRight() C.float {
	pos := Right
	if err := Validate.Position(float64(pos), "right"); err != nil {
		Log(LogLevelError, "PositionRight validation error: %v", err)
		return 1.0
	}
	Log(LogLevelDebug, "Returning right position: %f", pos)
	return C.float(pos)
}
