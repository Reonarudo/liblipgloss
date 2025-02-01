package main

/*
#include <stdbool.h>
#include <stdlib.h>
#include "lipgloss_types.h"
*/
import "C"

import (
	"strconv"
	"unsafe"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

//export MapTerminalColor
func MapTerminalColor(tcHandle unsafe.Pointer) *C.char {
	if tcHandle == nil {
		return C.CString("")
	}

	tc := *(*lipgloss.TerminalColor)(tcHandle)
	renderer := GetRenderer()
	if renderer == nil {
		return C.CString("")
	}

	var colorStr string
	switch t := tc.(type) {
	case lipgloss.Color:
		colorStr = string(t)
	case lipgloss.ANSIColor:
		colorStr = strconv.FormatUint(uint64(t), 10)
	case lipgloss.AdaptiveColor:
		if renderer.HasDarkBackground() {
			colorStr = t.Dark
		} else {
			colorStr = t.Light
		}
	case lipgloss.CompleteColor:
		switch renderer.ColorProfile() {
		case termenv.TrueColor:
			colorStr = t.TrueColor
		case termenv.ANSI256:
			colorStr = t.ANSI256
		default:
			colorStr = t.ANSI
		}
	case lipgloss.CompleteAdaptiveColor:
		color := t.Light
		if renderer.HasDarkBackground() {
			color = t.Dark
		}
		switch renderer.ColorProfile() {
		case termenv.TrueColor:
			colorStr = color.TrueColor
		case termenv.ANSI256:
			colorStr = color.ANSI256
		default:
			colorStr = color.ANSI
		}
	default:
		colorStr = ""
	}

	return C.CString(colorStr)
}

//export GetTerminalColorRGBA
func GetTerminalColorRGBA(tcHandle unsafe.Pointer) (r, g, b, a C.uint32_t) {
	if tcHandle == nil {
		return 0, 0, 0, 0xFFFF
	}

	tc := *(*lipgloss.TerminalColor)(tcHandle)
	rVal, gVal, bVal, aVal := tc.RGBA()
	return C.uint32_t(rVal), C.uint32_t(gVal), C.uint32_t(bVal), C.uint32_t(aVal)
}

// Color specifies a color by hex or ANSI value
type Color string

func (c Color) color(r *lipgloss.Renderer) termenv.Color {
	if r == nil {
		return termenv.NoColor{}
	}
	return r.ColorProfile().Color(string(c))
}

//export ColorRGBA
func ColorRGBA(c *C.char) (r, g, b, a C.uint32_t) {
	if c == nil {
		return 0, 0, 0, 0xFFFF
	}

	parsedColor := Color(C.GoString(c))
	rVal, gVal, bVal, aVal := parsedColor.RGBA()
	return C.uint32_t(rVal), C.uint32_t(gVal), C.uint32_t(bVal), C.uint32_t(aVal)
}

func (c Color) RGBA() (r, g, b, a uint32) {
	renderer := GetRenderer()
	if renderer == nil {
		return 0, 0, 0, 0xFFFF
	}
	return termenv.ConvertToRGB(c.color(renderer)).RGBA()
}

// ANSIColor represents a color specified by an ANSI value
type ANSIColor uint

func (ac ANSIColor) color(r *lipgloss.Renderer) termenv.Color {
	return Color(strconv.FormatUint(uint64(ac), 10)).color(r)
}

//export ANSIColorRGBA
func ANSIColorRGBA(value C.uint) (r, g, b, a C.uint32_t) {
	ac := ANSIColor(value)
	rVal, gVal, bVal, aVal := ac.RGBA()
	return C.uint32_t(rVal), C.uint32_t(gVal), C.uint32_t(bVal), C.uint32_t(aVal)
}

func (ac ANSIColor) RGBA() (r, g, b, a uint32) {
	return Color(strconv.FormatUint(uint64(ac), 10)).RGBA()
}

// AdaptiveColor provides color options for light and dark backgrounds
type AdaptiveColor struct {
	Light string
	Dark  string
}

func (ac AdaptiveColor) color(r *lipgloss.Renderer) termenv.Color {
	if r == nil {
		return termenv.NoColor{}
	}
	if r.HasDarkBackground() {
		return Color(ac.Dark).color(r)
	}
	return Color(ac.Light).color(r)
}

//export AdaptiveColorRGBA
func AdaptiveColorRGBA(light, dark *C.char) (r, g, b, a C.uint32_t) {
	if light == nil || dark == nil {
		return 0, 0, 0, 0xFFFF
	}

	ac := AdaptiveColor{
		Light: C.GoString(light),
		Dark:  C.GoString(dark),
	}
	rVal, gVal, bVal, aVal := ac.RGBA()
	return C.uint32_t(rVal), C.uint32_t(gVal), C.uint32_t(bVal), C.uint32_t(aVal)
}

func (ac AdaptiveColor) RGBA() (r, g, b, a uint32) {
	renderer := GetRenderer()
	if renderer == nil {
		return 0, 0, 0, 0xFFFF
	}
	return termenv.ConvertToRGB(ac.color(renderer)).RGBA()
}

// CompleteColor specifies exact values for truecolor, ANSI256, and ANSI profiles
type CompleteColor struct {
	TrueColor string
	ANSI256   string
	ANSI      string
}

func (cc CompleteColor) color(r *lipgloss.Renderer) termenv.Color {
	if r == nil {
		return termenv.NoColor{}
	}
	switch r.ColorProfile() {
	case termenv.TrueColor:
		return r.ColorProfile().Color(cc.TrueColor)
	case termenv.ANSI256:
		return r.ColorProfile().Color(cc.ANSI256)
	default:
		return r.ColorProfile().Color(cc.ANSI)
	}
}

//export CompleteColorRGBA
func CompleteColorRGBA(trueColor, ansi256, ansi *C.char) (r, g, b, a C.uint32_t) {
	if trueColor == nil || ansi256 == nil || ansi == nil {
		return 0, 0, 0, 0xFFFF
	}

	cc := CompleteColor{
		TrueColor: C.GoString(trueColor),
		ANSI256:   C.GoString(ansi256),
		ANSI:      C.GoString(ansi),
	}
	rVal, gVal, bVal, aVal := cc.RGBA()
	return C.uint32_t(rVal), C.uint32_t(gVal), C.uint32_t(bVal), C.uint32_t(aVal)
}

func (cc CompleteColor) RGBA() (r, g, b, a uint32) {
	renderer := GetRenderer()
	if renderer == nil {
		return 0, 0, 0, 0xFFFF
	}
	return termenv.ConvertToRGB(cc.color(renderer)).RGBA()
}

// CompleteAdaptiveColor specifies colors for light/dark backgrounds with full profiles
type CompleteAdaptiveColor struct {
	Light CompleteColor
	Dark  CompleteColor
}

func (cac CompleteAdaptiveColor) color(r *lipgloss.Renderer) termenv.Color {
	if r == nil {
		return termenv.NoColor{}
	}
	if r.HasDarkBackground() {
		return cac.Dark.color(r)
	}
	return cac.Light.color(r)
}

//export CompleteAdaptiveColorRGBA
func CompleteAdaptiveColorRGBA(lightTrue, lightANSI256, lightANSI, darkTrue, darkANSI256, darkANSI *C.char) (r, g, b, a C.uint32_t) {
	if lightTrue == nil || lightANSI256 == nil || lightANSI == nil ||
		darkTrue == nil || darkANSI256 == nil || darkANSI == nil {
		return 0, 0, 0, 0xFFFF
	}

	cac := CompleteAdaptiveColor{
		Light: CompleteColor{
			TrueColor: C.GoString(lightTrue),
			ANSI256:   C.GoString(lightANSI256),
			ANSI:      C.GoString(lightANSI),
		},
		Dark: CompleteColor{
			TrueColor: C.GoString(darkTrue),
			ANSI256:   C.GoString(darkANSI256),
			ANSI:      C.GoString(darkANSI),
		},
	}
	rVal, gVal, bVal, aVal := cac.RGBA()
	return C.uint32_t(rVal), C.uint32_t(gVal), C.uint32_t(bVal), C.uint32_t(aVal)
}

func (cac CompleteAdaptiveColor) RGBA() (r, g, b, a uint32) {
	renderer := GetRenderer()
	if renderer == nil {
		return 0, 0, 0, 0xFFFF
	}
	return termenv.ConvertToRGB(cac.color(renderer)).RGBA()
}
