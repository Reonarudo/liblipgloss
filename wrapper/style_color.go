package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import "github.com/charmbracelet/lipgloss"

//export StyleForeground
func StyleForeground(id C.uint64_t, color *C.char) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "foreground")
	if err != nil {
		Log(LogLevelError, "StyleForeground style error: %v", err)
		return 0
	}

	colorStr := String.GoString(color)
	if err := Validate.Color(colorStr, "foreground"); err != nil {
		Log(LogLevelError, "StyleForeground color error: %v", err)
		return 0
	}

	newStyle := style.Foreground(lipgloss.Color(colorStr))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new foreground style with ID: %d", uint64(id))
	return id
}

//export StyleBackground
func StyleBackground(id C.uint64_t, color *C.char) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "background")
	if err != nil {
		Log(LogLevelError, "StyleBackground style error: %v", err)
		return 0
	}

	colorStr := String.GoString(color)
	if err := Validate.Color(colorStr, "background"); err != nil {
		Log(LogLevelError, "StyleBackground color error: %v", err)
		return 0
	}

	newStyle := style.Background(lipgloss.Color(colorStr))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new background style with ID: %d", uint64(id))
	return id
}

//export StyleColorWhitespace
func StyleColorWhitespace(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "color-whitespace")
	if err != nil {
		Log(LogLevelError, "StyleColorWhitespace error: %v", err)
		return 0
	}

	newStyle := style.ColorWhitespace(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new color-whitespace style with ID: %d", uint64(id))
	return id
}

//export StyleMarginBackground
func StyleMarginBackground(id C.uint64_t, color *C.char) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "margin-background")
	if err != nil {
		Log(LogLevelError, "StyleMarginBackground style error: %v", err)
		return 0
	}

	colorStr := String.GoString(color)
	if err := Validate.Color(colorStr, "marginbackground"); err != nil {
		Log(LogLevelError, "StyleMarginBackground color error: %v", err)
		return 0
	}

	newStyle := style.MarginBackground(lipgloss.Color(colorStr))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new margin-background style with ID: %d", uint64(id))
	return id
}
