package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"github.com/charmbracelet/lipgloss"
)

//export StyleBorder
func StyleBorder(id C.uint64_t, border C.CBorder) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "border")
	if err != nil {
		Log(LogLevelError, "StyleBorder style error: %v", err)
		return 0
	}

	newStyle := style.Border(lipgloss.Border{
		Top:          C.GoString(border.Top),
		Bottom:       C.GoString(border.Bottom),
		Left:         C.GoString(border.Left),
		Right:        C.GoString(border.Right),
		TopLeft:      C.GoString(border.TopLeft),
		TopRight:     C.GoString(border.TopRight),
		BottomLeft:   C.GoString(border.BottomLeft),
		BottomRight:  C.GoString(border.BottomRight),
		MiddleLeft:   C.GoString(border.MiddleLeft),
		MiddleRight:  C.GoString(border.MiddleRight),
		Middle:       C.GoString(border.Middle),
		MiddleTop:    C.GoString(border.MiddleTop),
		MiddleBottom: C.GoString(border.MiddleBottom),
	})

	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new border style with ID: %d", uint64(id))
	return id
}

//export StyleBorderStyle
func StyleBorderStyle(id C.uint64_t, border C.CBorder) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "border-style")
	if err != nil {
		Log(LogLevelError, "StyleBorderStyle style error: %v", err)
		return 0
	}

	newStyle := style.BorderStyle(lipgloss.Border{
		Top:          C.GoString(border.Top),
		Bottom:       C.GoString(border.Bottom),
		Left:         C.GoString(border.Left),
		Right:        C.GoString(border.Right),
		TopLeft:      C.GoString(border.TopLeft),
		TopRight:     C.GoString(border.TopRight),
		BottomLeft:   C.GoString(border.BottomLeft),
		BottomRight:  C.GoString(border.BottomRight),
		MiddleLeft:   C.GoString(border.MiddleLeft),
		MiddleRight:  C.GoString(border.MiddleRight),
		Middle:       C.GoString(border.Middle),
		MiddleTop:    C.GoString(border.MiddleTop),
		MiddleBottom: C.GoString(border.MiddleBottom),
	})

	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new border style with ID: %d", uint64(id))
	return id
}

//export StyleBorderBackground
func StyleBorderBackground(id C.uint64_t, color *C.char) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "border-background")
	if err != nil {
		Log(LogLevelError, "StyleBorderBackground style error: %v", err)
		return 0
	}

	colorStr := String.GoString(color)
	if err := Validate.Color(colorStr, "border-background"); err != nil {
		Log(LogLevelError, "StyleBorderBackground color error: %v", err)
		return 0
	}

	newStyle := style.BorderBackground(lipgloss.Color(colorStr))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new border background style with ID: %d", uint64(id))
	return id
}

//export StyleBorderForeground
func StyleBorderForeground(id C.uint64_t, color *C.char) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "border-foreground")
	if err != nil {
		Log(LogLevelError, "StyleBorderForeground style error: %v", err)
		return 0
	}

	colorStr := String.GoString(color)
	if err := Validate.Color(colorStr, "border-foreground"); err != nil {
		Log(LogLevelError, "StyleBorderForeground color error: %v", err)
		return 0
	}

	newStyle := style.BorderForeground(lipgloss.Color(colorStr))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new border foreground style with ID: %d", uint64(id))
	return id
}

//export StyleGetBorderStyle
func StyleGetBorderStyle(id C.uint64_t) C.CBorder {
	style, err := Style.SafeGet(uint64(id), "get-border-style")
	if err != nil {
		Log(LogLevelError, "StyleGetBorderStyle error: %v", err)
		return C.CBorder{}
	}

	border := style.GetBorderStyle()
	return C.CBorder{
		Top:          C.CString(border.Top),
		Bottom:       C.CString(border.Bottom),
		Left:         C.CString(border.Left),
		Right:        C.CString(border.Right),
		TopLeft:      C.CString(border.TopLeft),
		TopRight:     C.CString(border.TopRight),
		BottomLeft:   C.CString(border.BottomLeft),
		BottomRight:  C.CString(border.BottomRight),
		MiddleLeft:   C.CString(border.MiddleLeft),
		MiddleRight:  C.CString(border.MiddleRight),
		Middle:       C.CString(border.Middle),
		MiddleTop:    C.CString(border.MiddleTop),
		MiddleBottom: C.CString(border.MiddleBottom),
	}
}
