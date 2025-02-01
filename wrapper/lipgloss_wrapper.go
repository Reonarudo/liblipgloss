package main

/*
#include <stdbool.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"unsafe"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

//export ColorProfile
func ColorProfile() *C.char {
	profile := lipgloss.ColorProfile()
	var profileStr string
	switch profile {
	case termenv.Ascii:
		profileStr = "ascii"
	case termenv.ANSI:
		profileStr = "ansi"
	case termenv.ANSI256:
		profileStr = "ansi256"
	case termenv.TrueColor:
		profileStr = "truecolor"
	default:
		profileStr = "unknown"
	}

	cs, err := String.CString(profileStr)
	if err != nil {
		Log(LogLevelError, "ColorProfile memory allocation error: %v", err)
		defaultCs, _ := String.CString("ascii") // Safe fallback
		return defaultCs
	}
	Memory.Track(unsafe.Pointer(cs), "ColorProfile result")
	return cs
}

//export HasDarkBackground
func HasDarkBackground() C.bool {
	return C.bool(lipgloss.HasDarkBackground())
}

//export Height
func Height(str *C.char) C.int {
	goStr := String.GoString(str)
	height := lipgloss.Height(goStr)
	Log(LogLevelDebug, "Calculated height %d for string", height)
	return C.int(height)
}

//export JoinHorizontal
func JoinHorizontal(pos C.double, str1 *C.char, str2 *C.char) *C.char {
	goStr1 := String.GoString(str1)
	goStr2 := String.GoString(str2)

	if err := Validate.Position(float64(pos), "horizontal"); err != nil {
		Log(LogLevelError, "JoinHorizontal position error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	joined := lipgloss.JoinHorizontal(lipgloss.Position(pos), goStr1, goStr2)
	cs, err := String.CString(joined)
	if err != nil {
		Log(LogLevelError, "JoinHorizontal memory allocation error: %v", err)
		defaultCs, _ := String.CString("") // Safe fallback
		return defaultCs
	}

	Memory.Track(unsafe.Pointer(cs), "JoinHorizontal result")
	return cs
}

//export JoinVertical
func JoinVertical(pos C.double, str1 *C.char, str2 *C.char) *C.char {
	goStr1 := String.GoString(str1)
	goStr2 := String.GoString(str2)

	joined := lipgloss.JoinVertical(lipgloss.Position(pos), goStr1, goStr2)
	cs, err := String.CString(joined)
	if err != nil {
		Log(LogLevelError, "JoinVertical memory allocation error: %v", err)
		defaultCs, _ := String.CString("") // Safe fallback
		return defaultCs
	}

	Memory.Track(unsafe.Pointer(cs), "JoinVertical result")
	return cs
}

//export Place
func Place(width, height C.int, hPos, vPos C.double, str *C.char) *C.char {
	if err := Validate.Dimension(int(width), "width"); err != nil {
		Log(LogLevelError, "Place width validation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}
	if err := Validate.Dimension(int(height), "height"); err != nil {
		Log(LogLevelError, "Place height validation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	goStr := String.GoString(str)
	placed := lipgloss.Place(int(width), int(height),
		lipgloss.Position(hPos), lipgloss.Position(vPos), goStr)

	cs, err := String.CString(placed)
	if err != nil {
		Log(LogLevelError, "Place memory allocation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	Memory.Track(unsafe.Pointer(cs), "Place result")
	return cs
}

//export PlaceHorizontal
func PlaceHorizontal(width C.int, pos C.double, str *C.char) *C.char {
	if err := Validate.Dimension(int(width), "width"); err != nil {
		Log(LogLevelError, "PlaceHorizontal width validation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	goStr := String.GoString(str)
	placed := lipgloss.PlaceHorizontal(int(width), lipgloss.Position(pos), goStr)

	cs, err := String.CString(placed)
	if err != nil {
		Log(LogLevelError, "PlaceHorizontal memory allocation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	Memory.Track(unsafe.Pointer(cs), "PlaceHorizontal result")
	return cs
}

//export PlaceVertical
func PlaceVertical(height C.int, pos C.double, str *C.char) *C.char {
	if err := Validate.Dimension(int(height), "height"); err != nil {
		Log(LogLevelError, "PlaceVertical height validation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	goStr := String.GoString(str)
	placed := lipgloss.PlaceVertical(int(height), lipgloss.Position(pos), goStr)

	cs, err := String.CString(placed)
	if err != nil {
		Log(LogLevelError, "PlaceVertical memory allocation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	Memory.Track(unsafe.Pointer(cs), "PlaceVertical result")
	return cs
}

//export SetColorProfile
func SetColorProfile(profile *C.char) {
	profileStr := String.GoString(profile)
	var termProfile termenv.Profile

	switch profileStr {
	case "ascii":
		termProfile = termenv.Ascii
	case "ansi":
		termProfile = termenv.ANSI
	case "ansi256":
		termProfile = termenv.ANSI256
	case "truecolor":
		termProfile = termenv.TrueColor
	default:
		Log(LogLevelError, "SetColorProfile received invalid profile: %s", profileStr)
		return
	}

	lipgloss.SetColorProfile(termProfile)
	Log(LogLevelDebug, "Set color profile to: %s", profileStr)
}

//export SetHasDarkBackground
func SetHasDarkBackground(b C.bool) {
	lipgloss.SetHasDarkBackground(bool(b))
	Log(LogLevelDebug, "Set dark background to: %v", bool(b))
}

//export Size
func Size(str *C.char) (C.int, C.int) {
	goStr := String.GoString(str)
	width, height := lipgloss.Size(goStr)
	Log(LogLevelDebug, "Calculated size: width=%d, height=%d", width, height)
	return C.int(width), C.int(height)
}

//export StyleRunes
func StyleRunes(str *C.char, indices *C.int, indicesLen C.int,
	matchedHandle, unmatchedHandle unsafe.Pointer) *C.char {

	if indices == nil || indicesLen <= 0 {
		Log(LogLevelError, "StyleRunes received invalid indices")
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	goStr := String.GoString(str)
	goIndices := unsafe.Slice((*C.int)(indices), indicesLen)
	indicesSlice := make([]int, len(goIndices))

	for i, val := range goIndices {
		indicesSlice[i] = int(val)
	}

	matched := *(*lipgloss.Style)(matchedHandle)
	unmatched := *(*lipgloss.Style)(unmatchedHandle)

	styled := lipgloss.StyleRunes(goStr, indicesSlice, matched, unmatched)
	cs, err := String.CString(styled)
	if err != nil {
		Log(LogLevelError, "StyleRunes memory allocation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}

	Memory.Track(unsafe.Pointer(cs), "StyleRunes result")
	return cs
}

//export Width
func Width(str *C.char) C.int {
	goStr := String.GoString(str)
	width := lipgloss.Width(goStr)
	Log(LogLevelDebug, "Calculated width %d for string", width)
	return C.int(width)
}

func main() {}
