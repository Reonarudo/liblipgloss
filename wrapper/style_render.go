package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import "unsafe"

//export StyleRender
func StyleRender(id C.uint64_t, str *C.char) *C.char {
	style, err := Style.SafeGet(uint64(id), "render")
	if err != nil {
		Log(LogLevelError, "StyleRender error: %v", err)
		cs, _ := String.CString("")
		return cs
	}

	inputStr := String.GoString(str)
	result := style.Render(inputStr)
	cs, err := String.CString(result)
	if err != nil {
		Log(LogLevelError, "StyleRender memory allocation error: %v", err)
		cs, _ = String.CString("")
		return cs
	}

	Memory.Track(unsafe.Pointer(cs), "StyleRender result")
	return cs
}

//export StyleInherited
func StyleInherited(id C.uint64_t) C.int {
	style, err := Style.SafeGet(uint64(id), "inherited")
	if err != nil {
		Log(LogLevelError, "StyleInherited error: %v", err)
		return 0
	}

	// Check all inheritable properties
	border := style.GetBorderStyle()
	hasBorder := border.Top != "" || border.Bottom != "" ||
		border.Left != "" || border.Right != ""

	// Add checks for other inheritable properties
	hasColor := style.GetForeground() != nil || style.GetBackground() != nil
	hasMargin := style.GetMarginTop() != 0 || style.GetMarginRight() != 0 ||
		style.GetMarginBottom() != 0 || style.GetMarginLeft() != 0
	hasPadding := style.GetPaddingTop() != 0 || style.GetPaddingRight() != 0 ||
		style.GetPaddingBottom() != 0 || style.GetPaddingLeft() != 0

	return String.ToCInt(hasBorder || hasColor || hasMargin || hasPadding)
}

//export StyleString
func StyleString(id C.uint64_t) *C.char {
	style, err := Style.SafeGet(uint64(id), "string")
	if err != nil {
		Log(LogLevelError, "StyleString error: %v", err)
		cs, _ := String.CString("")
		return cs
	}

	cs, err := String.CString(style.String())
	if err != nil {
		Log(LogLevelError, "StyleString memory allocation error: %v", err)
		cs, _ = String.CString("")
		return cs
	}

	Memory.Track(unsafe.Pointer(cs), "StyleString result")
	return cs
}

//export StyleInherit
func StyleInherit(baseID, inheritID C.uint64_t) C.uint64_t {
	baseStyle, err := Style.SafeGet(uint64(baseID), "inherit-base")
	if err != nil {
		Log(LogLevelError, "StyleInherit base error: %v", err)
		return 0
	}

	inheritStyle, err := Style.SafeGet(uint64(inheritID), "inherit-from")
	if err != nil {
		Log(LogLevelError, "StyleInherit inherit error: %v", err)
		return 0
	}

	newStyle := baseStyle.Inherit(*inheritStyle)
	id := styleReg.Register(&newStyle)
	Log(LogLevelDebug, "Created new inherited style with ID: %d", id)

	return C.uint64_t(id)
}

// Cleanup helper
//
//export StyleCleanup
func StyleCleanup() *C.char {
	return GetMemoryLeaks()
}
