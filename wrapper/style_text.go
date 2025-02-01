package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/charmbracelet/lipgloss"
)

// TextStyleError represents an error in text styling operations
type TextStyleError struct {
	Op      string
	Message string
}

func (e *TextStyleError) Error() string {
	return fmt.Sprintf("text style error (op=%s): %s", e.Op, e.Message)
}

//export StyleSetString
func StyleSetString(id C.uint64_t, str *C.char) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "set-string")
	if err != nil {
		Log(LogLevelError, "StyleSetString style error: %v", err)
		return 0
	}

	goStr := String.GoString(str)
	newStyle := style.SetString(goStr)
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new style with set string, ID: %d", uint64(id))
	return id
}

//export StyleGetValue
func StyleGetValue(id C.uint64_t) *C.char {
	style, err := Style.SafeGet(uint64(id), "get-value")
	if err != nil {
		Log(LogLevelError, "StyleGetValue style error: %v", err)
		return C.CString("")
	}

	value := style.Value()
	cs, err := String.CString(value)
	if err != nil {
		Log(LogLevelError, "StyleGetValue memory allocation error: %v", err)
		return C.CString("")
	}

	Memory.Track(unsafe.Pointer(cs), "style value string")
	return cs
}

//export StyleBold
func StyleBold(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "bold")
	if err != nil {
		Log(LogLevelError, "StyleBold style error: %v", err)
		return 0
	}

	newStyle := style.Bold(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new bold style with ID: %d", uint64(id))
	return id
}

//export StyleItalic
func StyleItalic(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "italic")
	if err != nil {
		Log(LogLevelError, "StyleItalic style error: %v", err)
		return 0
	}

	newStyle := style.Italic(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new italic style with ID: %d", uint64(id))
	return id
}

//export StyleUnderline
func StyleUnderline(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "underline")
	if err != nil {
		Log(LogLevelError, "StyleUnderline style error: %v", err)
		return 0
	}

	newStyle := style.Underline(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new underline style with ID: %d", uint64(id))
	return id
}

//export StyleStrikethrough
func StyleStrikethrough(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "strikethrough")
	if err != nil {
		Log(LogLevelError, "StyleStrikethrough style error: %v", err)
		return 0
	}

	newStyle := style.Strikethrough(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new strikethrough style with ID: %d", uint64(id))
	return id
}

//export StyleReverse
func StyleReverse(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "reverse")
	if err != nil {
		Log(LogLevelError, "StyleReverse style error: %v", err)
		return 0
	}

	newStyle := style.Reverse(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new reverse style with ID: %d", uint64(id))
	return id
}

//export StyleBlink
func StyleBlink(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "blink")
	if err != nil {
		Log(LogLevelError, "StyleBlink style error: %v", err)
		return 0
	}

	newStyle := style.Blink(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new blink style with ID: %d", uint64(id))
	return id
}

//export StyleFaint
func StyleFaint(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "faint")
	if err != nil {
		Log(LogLevelError, "StyleFaint style error: %v", err)
		return 0
	}

	newStyle := style.Faint(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new faint style with ID: %d", uint64(id))
	return id
}

// validateTextStyle performs common validation for text style operations
func validateTextStyle(style *lipgloss.Style, op string) error {
	if style == nil {
		return &TextStyleError{
			Op:      op,
			Message: "nil style",
		}
	}
	return nil
}

//export GetTextStyleInfo
func GetTextStyleInfo(id C.uint64_t) *C.char {
	style, err := Style.SafeGet(uint64(id), "get-info")
	if err != nil {
		Log(LogLevelError, "GetTextStyleInfo style error: %v", err)
		return C.CString("Error: Style not found")
	}

	info := fmt.Sprintf("Style ID: %d\nBold: %v\nItalic: %v\nUnderline: %v\nStrikethrough: %v\n",
		id,
		style.GetBold(),
		style.GetItalic(),
		style.GetUnderline(),
		style.GetStrikethrough())

	cs, err := String.CString(info)
	if err != nil {
		Log(LogLevelError, "GetTextStyleInfo memory allocation error: %v", err)
		return C.CString("Error: Memory allocation failed")
	}

	Memory.Track(unsafe.Pointer(cs), "text style info string")
	return cs
}
