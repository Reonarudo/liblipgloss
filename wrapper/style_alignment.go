package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// validatePadding validates padding/margin values are non-negative
func validatePadding(value int, name string) error {
	if value < 0 {
		return &StyleError{
			Op:      "padding-validation",
			Message: fmt.Sprintf("invalid %s value: %d (must be non-negative)", name, value),
		}
	}
	return nil
}

//export StyleAlignHorizontal
func StyleAlignHorizontal(id C.uint64_t, position C.double) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "align-horizontal")
	if err != nil {
		Log(LogLevelError, "StyleAlignHorizontal style error: %v", err)
		return 0
	}

	if err := Validate.Position(float64(position), "align-horizontal"); err != nil {
		Log(LogLevelError, "StyleAlignHorizontal position error: %v", err)
		return 0
	}

	newStyle := style.Align(lipgloss.Position(position))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new horizontal alignment style with ID: %d", uint64(id))
	return id
}

//export StyleAlignVertical
func StyleAlignVertical(id C.uint64_t, position C.double) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "align-vertical")
	if err != nil {
		Log(LogLevelError, "StyleAlignVertical style error: %v", err)
		return 0
	}

	if err := Validate.Position(float64(position), "align-vertical"); err != nil {
		Log(LogLevelError, "StyleAlignVertical position error: %v", err)
		return 0
	}

	newStyle := style.AlignVertical(lipgloss.Position(position))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new vertical alignment style with ID: %d", uint64(id))
	return id
}

//export StylePadding
func StylePadding(id C.uint64_t, top, right, bottom, left C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "padding")
	if err != nil {
		Log(LogLevelError, "StylePadding style error: %v", err)
		return 0
	}

	// Validate all padding values
	for _, v := range []struct {
		value int
		name  string
	}{
		{int(top), "top"},
		{int(right), "right"},
		{int(bottom), "bottom"},
		{int(left), "left"},
	} {
		if err := validatePadding(v.value, v.name); err != nil {
			Log(LogLevelError, "StylePadding validation error: %v", err)
			return 0
		}
	}

	newStyle := style.Padding(int(top), int(right), int(bottom), int(left))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new padding style with ID: %d", uint64(id))
	return id
}

//export StylePaddingTop
func StylePaddingTop(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "padding-top")
	if err != nil {
		Log(LogLevelError, "StylePaddingTop style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "top"); err != nil {
		Log(LogLevelError, "StylePaddingTop validation error: %v", err)
		return 0
	}

	newStyle := style.PaddingTop(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new padding-top style with ID: %d", uint64(id))
	return id
}

//export StylePaddingRight
func StylePaddingRight(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "padding-right")
	if err != nil {
		Log(LogLevelError, "StylePaddingRight style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "right"); err != nil {
		Log(LogLevelError, "StylePaddingRight validation error: %v", err)
		return 0
	}

	newStyle := style.PaddingRight(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new padding-right style with ID: %d", uint64(id))
	return id
}

//export StylePaddingBottom
func StylePaddingBottom(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "padding-bottom")
	if err != nil {
		Log(LogLevelError, "StylePaddingBottom style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "bottom"); err != nil {
		Log(LogLevelError, "StylePaddingBottom validation error: %v", err)
		return 0
	}

	newStyle := style.PaddingBottom(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new padding-bottom style with ID: %d", uint64(id))
	return id
}

//export StylePaddingLeft
func StylePaddingLeft(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "padding-left")
	if err != nil {
		Log(LogLevelError, "StylePaddingLeft style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "left"); err != nil {
		Log(LogLevelError, "StylePaddingLeft validation error: %v", err)
		return 0
	}

	newStyle := style.PaddingLeft(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new padding-left style with ID: %d", uint64(id))
	return id
}

//export StyleMargin
func StyleMargin(id C.uint64_t, top, right, bottom, left C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "margin")
	if err != nil {
		Log(LogLevelError, "StyleMargin style error: %v", err)
		return 0
	}

	// Validate all margin values
	for _, v := range []struct {
		value int
		name  string
	}{
		{int(top), "top"},
		{int(right), "right"},
		{int(bottom), "bottom"},
		{int(left), "left"},
	} {
		if err := validatePadding(v.value, v.name); err != nil {
			Log(LogLevelError, "StyleMargin validation error: %v", err)
			return 0
		}
	}

	newStyle := style.Margin(int(top), int(right), int(bottom), int(left))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new margin style with ID: %d", uint64(id))
	return id
}

//export StyleMarginTop
func StyleMarginTop(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "margin-top")
	if err != nil {
		Log(LogLevelError, "StyleMarginTop style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "top"); err != nil {
		Log(LogLevelError, "StyleMarginTop validation error: %v", err)
		return 0
	}

	newStyle := style.MarginTop(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new margin-top style with ID: %d", uint64(id))
	return id
}

//export StyleMarginRight
func StyleMarginRight(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "margin-right")
	if err != nil {
		Log(LogLevelError, "StyleMarginRight style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "right"); err != nil {
		Log(LogLevelError, "StyleMarginRight validation error: %v", err)
		return 0
	}

	newStyle := style.MarginRight(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new margin-right style with ID: %d", uint64(id))
	return id
}

//export StyleMarginBottom
func StyleMarginBottom(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "margin-bottom")
	if err != nil {
		Log(LogLevelError, "StyleMarginBottom style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "bottom"); err != nil {
		Log(LogLevelError, "StyleMarginBottom validation error: %v", err)
		return 0
	}

	newStyle := style.MarginBottom(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new margin-bottom style with ID: %d", uint64(id))
	return id
}

//export StyleMarginLeft
func StyleMarginLeft(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "margin-left")
	if err != nil {
		Log(LogLevelError, "StyleMarginLeft style error: %v", err)
		return 0
	}

	if err := validatePadding(int(v), "left"); err != nil {
		Log(LogLevelError, "StyleMarginLeft validation error: %v", err)
		return 0
	}

	newStyle := style.MarginLeft(int(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new margin-left style with ID: %d", uint64(id))
	return id
}
