package main

/*
#include <stdio.h>
#include <stdbool.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"fmt"
	"os"
	"sync"
	"unsafe"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// RendererError represents an error in renderer operations
type RendererError struct {
	Op      string
	Message string
}

func (e *RendererError) Error() string {
	return fmt.Sprintf("renderer error (op=%s): %s", e.Op, e.Message)
}

// rendererRegistry provides thread-safe access to renderers
type rendererRegistry struct {
	sync.RWMutex
	defaultRenderer *lipgloss.Renderer
	activeOutput    *os.File
}

var (
	rendererReg = &rendererRegistry{}
)

// getRenderer safely gets the default renderer
func GetRenderer() *lipgloss.Renderer {
	rendererReg.RLock()
	defer rendererReg.RUnlock()
	return rendererReg.defaultRenderer
}

// setRenderer safely sets the default renderer and tracks the output
func setRenderer(r *lipgloss.Renderer, output *os.File) {
	rendererReg.Lock()
	defer rendererReg.Unlock()

	if rendererReg.defaultRenderer != nil {
		Log(LogLevelDebug, "Replacing existing renderer")
	}

	rendererReg.defaultRenderer = r
	rendererReg.activeOutput = output
	Log(LogLevelDebug, "Set new default renderer with output: %v", output)
}

// validateRenderer checks if a renderer is available and valid
func validateRenderer(op string) error {
	renderer := GetRenderer()
	if renderer == nil {
		return &RendererError{
			Op:      op,
			Message: "no renderer available",
		}
	}

	if rendererReg.activeOutput == nil {
		return &RendererError{
			Op:      op,
			Message: "renderer has no valid output",
		}
	}

	return nil
}

//export DefaultRenderer
func DefaultRenderer() {
	setRenderer(lipgloss.DefaultRenderer(), os.Stdout)
	Log(LogLevelDebug, "Initialized default renderer with stdout")
}

//export NewRenderer
func NewRenderer(w *C.FILE) {
	if w == nil {
		Log(LogLevelError, "NewRenderer received nil file pointer")
		return
	}

	file := os.NewFile(uintptr(C.fileno(w)), "cfile")
	if file == nil {
		Log(LogLevelError, "Failed to create file from descriptor")
		return
	}

	renderer := lipgloss.NewRenderer(file)
	setRenderer(renderer, file)
	Log(LogLevelDebug, "Created new renderer with custom output")
}

//export RendererColorProfile
func RendererColorProfile() *C.char {
	if err := validateRenderer("color-profile"); err != nil {
		Log(LogLevelError, "RendererColorProfile error: %v", err)
		return C.CString("ascii") // Safe default
	}

	renderer := GetRenderer()
	profile := renderer.ColorProfile()
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
		profileStr = "ascii"
	}

	cs, err := String.CString(profileStr)
	if err != nil {
		Log(LogLevelError, "RendererColorProfile memory allocation error: %v", err)
		return C.CString("ascii")
	}

	Memory.Track(unsafe.Pointer(cs), "color profile string")
	return cs
}

//export RendererHasDarkBackground
func RendererHasDarkBackground() C.bool {
	if err := validateRenderer("dark-background"); err != nil {
		Log(LogLevelError, "RendererHasDarkBackground error: %v", err)
		return C.bool(false)
	}

	renderer := GetRenderer()
	return C.bool(renderer.HasDarkBackground())
}

//export RendererNewStyle
func RendererNewStyle() unsafe.Pointer {
	renderer := GetRenderer()
	if renderer == nil {
		Log(LogLevelWarn, "No renderer available, using default")
		renderer = lipgloss.DefaultRenderer()
		setRenderer(renderer, os.Stdout)
	}

	style := renderer.NewStyle()
	Log(LogLevelDebug, "Created new style from renderer")
	return unsafe.Pointer(&style)
}

//export RendererPlace
func RendererPlace(width, height C.int, hPos, vPos C.double, str *C.char) *C.char {
	if err := validateRenderer("place"); err != nil {
		Log(LogLevelError, "RendererPlace error: %v", err)
		return C.CString(String.GoString(str))
	}
	if err := Validate.Dimension(int(width), "place width"); err != nil {
		Log(LogLevelError, "RendererPlace width validation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}
	if err := Validate.Dimension(int(height), "place height"); err != nil {
		Log(LogLevelError, "RendererPlace height validation error: %v", err)
		defaultCs, _ := String.CString("")
		return defaultCs
	}
	if err := Validate.Position(float64(hPos), "horizontal"); err != nil {
		Log(LogLevelError, "RendererPlace horizontal position error: %v", err)
		return C.CString(String.GoString(str))
	}

	if err := Validate.Position(float64(vPos), "vertical"); err != nil {
		Log(LogLevelError, "RendererPlace vertical position error: %v", err)
		return C.CString(String.GoString(str))
	}

	renderer := GetRenderer()
	goStr := String.GoString(str)
	placed := renderer.Place(int(width), int(height),
		lipgloss.Position(hPos), lipgloss.Position(vPos), goStr)

	cs, err := String.CString(placed)
	if err != nil {
		Log(LogLevelError, "RendererPlace memory allocation error: %v", err)
		return C.CString(goStr)
	}

	Memory.Track(unsafe.Pointer(cs), "placed string")
	return cs
}

//export RendererPlaceHorizontal
func RendererPlaceHorizontal(width C.int, pos C.double, str *C.char) *C.char {
	if err := validateRenderer("place-horizontal"); err != nil {
		Log(LogLevelError, "RendererPlaceHorizontal error: %v", err)
		return C.CString(String.GoString(str))
	}

	if err := Validate.Dimension(int(width), "place-horizontal"); err != nil {
		Log(LogLevelError, "RendererPlaceHorizontal width error: %v", err)
		return C.CString(String.GoString(str))
	}

	if err := Validate.Position(float64(pos), "horizontal"); err != nil {
		Log(LogLevelError, "RendererPlaceHorizontal position error: %v", err)
		return C.CString(String.GoString(str))
	}

	renderer := GetRenderer()
	goStr := String.GoString(str)
	placed := renderer.PlaceHorizontal(int(width), lipgloss.Position(pos), goStr)

	cs, err := String.CString(placed)
	if err != nil {
		Log(LogLevelError, "RendererPlaceHorizontal memory allocation error: %v", err)
		return C.CString(goStr)
	}

	Memory.Track(unsafe.Pointer(cs), "horizontally placed string")
	return cs
}

//export RendererPlaceVertical
func RendererPlaceVertical(height C.int, pos C.double, str *C.char) *C.char {
	if err := validateRenderer("place-vertical"); err != nil {
		Log(LogLevelError, "RendererPlaceVertical error: %v", err)
		return C.CString(String.GoString(str))
	}

	if err := Validate.Dimension(int(height), "place-vertical"); err != nil {
		Log(LogLevelError, "RendererPlaceVertical height error: %v", err)
		return C.CString(String.GoString(str))
	}

	if err := Validate.Position(float64(pos), "vertical"); err != nil {
		Log(LogLevelError, "RendererPlaceVertical position error: %v", err)
		return C.CString(String.GoString(str))
	}

	renderer := GetRenderer()
	goStr := String.GoString(str)
	placed := renderer.PlaceVertical(int(height), lipgloss.Position(pos), goStr)

	cs, err := String.CString(placed)
	if err != nil {
		Log(LogLevelError, "RendererPlaceVertical memory allocation error: %v", err)
		return C.CString(goStr)
	}

	Memory.Track(unsafe.Pointer(cs), "vertically placed string")
	return cs
}

//export RendererSetColorProfile
func RendererSetColorProfile(p *C.char) {
	if err := validateRenderer("set-color-profile"); err != nil {
		Log(LogLevelError, "RendererSetColorProfile error: %v", err)
		return
	}

	profileStr := String.GoString(p)
	var profile termenv.Profile

	switch profileStr {
	case "ascii":
		profile = termenv.Ascii
	case "ansi":
		profile = termenv.ANSI
	case "ansi256":
		profile = termenv.ANSI256
	case "truecolor":
		profile = termenv.TrueColor
	default:
		Log(LogLevelError, "Invalid color profile specified: %s", profileStr)
		return
	}

	renderer := GetRenderer()
	renderer.SetColorProfile(profile)
	Log(LogLevelDebug, "Set color profile to: %s", profileStr)
}

//export RendererSetHasDarkBackground
func RendererSetHasDarkBackground(b C.bool) {
	if err := validateRenderer("set-dark-background"); err != nil {
		Log(LogLevelError, "RendererSetHasDarkBackground error: %v", err)
		return
	}

	renderer := GetRenderer()
	renderer.SetHasDarkBackground(bool(b))
	Log(LogLevelDebug, "Set dark background to: %v", bool(b))
}

//export RendererSetOutput
func RendererSetOutput(o *C.FILE) {
	if err := validateRenderer("set-output"); err != nil {
		Log(LogLevelError, "RendererSetOutput error: %v", err)
		return
	}

	if o == nil {
		Log(LogLevelError, "RendererSetOutput received nil file pointer")
		return
	}

	file := os.NewFile(uintptr(C.fileno(o)), "cfile")
	if file == nil {
		Log(LogLevelError, "Failed to create file from descriptor")
		return
	}

	output := termenv.NewOutput(file)
	renderer := GetRenderer()
	renderer.SetOutput(output)
	rendererReg.activeOutput = file
	Log(LogLevelDebug, "Set new renderer output")
}
