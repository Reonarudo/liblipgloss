package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import "fmt"

// validateDimension validates width/height values
func validateDimension(value int, name string) error {
	if value < 0 {
		return &StyleError{
			Op:      "dimension-validation",
			Message: fmt.Sprintf("invalid %s: %d (must be non-negative)", name, value),
		}
	}
	return nil
}

//export StyleWidth
func StyleWidth(id C.uint64_t, width C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "width")
	if err != nil {
		Log(LogLevelError, "StyleWidth style error: %v", err)
		return 0
	}

	if err := validateDimension(int(width), "width"); err != nil {
		Log(LogLevelError, "StyleWidth validation error: %v", err)
		return 0
	}

	newStyle := style.Width(int(width))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new width style with ID: %d", uint64(id))
	return id
}

//export StyleHeight
func StyleHeight(id C.uint64_t, height C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "height")
	if err != nil {
		Log(LogLevelError, "StyleHeight style error: %v", err)
		return 0
	}

	if err := validateDimension(int(height), "height"); err != nil {
		Log(LogLevelError, "StyleHeight validation error: %v", err)
		return 0
	}

	newStyle := style.Height(int(height))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new height style with ID: %d", uint64(id))
	return id
}

//export StyleMaxWidth
func StyleMaxWidth(id C.uint64_t, width C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "max-width")
	if err != nil {
		Log(LogLevelError, "StyleMaxWidth style error: %v", err)
		return 0
	}

	if err := validateDimension(int(width), "max-width"); err != nil {
		Log(LogLevelError, "StyleMaxWidth validation error: %v", err)
		return 0
	}

	newStyle := style.MaxWidth(int(width))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new max-width style with ID: %d", uint64(id))
	return id
}

//export StyleMaxHeight
func StyleMaxHeight(id C.uint64_t, height C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "max-height")
	if err != nil {
		Log(LogLevelError, "StyleMaxHeight style error: %v", err)
		return 0
	}

	if err := validateDimension(int(height), "max-height"); err != nil {
		Log(LogLevelError, "StyleMaxHeight validation error: %v", err)
		return 0
	}

	newStyle := style.MaxHeight(int(height))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new max-height style with ID: %d", uint64(id))
	return id
}

//export StyleInline
func StyleInline(id C.uint64_t, v C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "inline")
	if err != nil {
		Log(LogLevelError, "StyleInline style error: %v", err)
		return 0
	}

	newStyle := style.Inline(String.ToBool(v))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new inline style with ID: %d", uint64(id))
	return id
}

//export StyleTabWidth
func StyleTabWidth(id C.uint64_t, width C.int) C.uint64_t {
	style, err := Style.SafeGet(uint64(id), "tab-width")
	if err != nil {
		Log(LogLevelError, "StyleTabWidth style error: %v", err)
		return 0
	}

	// Special case: -1 is allowed for NoTabConversion
	if width != -1 {
		if err := validateDimension(int(width), "tab-width"); err != nil {
			Log(LogLevelError, "StyleTabWidth validation error: %v", err)
			return 0
		}
	}

	newStyle := style.TabWidth(int(width))
	id = C.uint64_t(styleReg.Register(&newStyle))
	Log(LogLevelDebug, "Created new tab-width style with ID: %d", uint64(id))
	return id
}
