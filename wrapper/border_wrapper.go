package main

/*
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include "lipgloss_types.h"
*/
import "C"

import (
	"unsafe"

	"github.com/charmbracelet/lipgloss"
)

// toBorder converts a lipgloss.Border to C.CBorder
func toBorder(b lipgloss.Border) C.CBorder {
	return C.CBorder{
		Top:          C.CString(b.Top),
		Bottom:       C.CString(b.Bottom),
		Left:         C.CString(b.Left),
		Right:        C.CString(b.Right),
		TopLeft:      C.CString(b.TopLeft),
		TopRight:     C.CString(b.TopRight),
		BottomLeft:   C.CString(b.BottomLeft),
		BottomRight:  C.CString(b.BottomRight),
		MiddleLeft:   C.CString(b.MiddleLeft),
		MiddleRight:  C.CString(b.MiddleRight),
		Middle:       C.CString(b.Middle),
		MiddleTop:    C.CString(b.MiddleTop),
		MiddleBottom: C.CString(b.MiddleBottom),
	}
}

// freeBorder frees memory allocated for C.CBorder strings
func freeBorder(b C.CBorder) {
	C.free(unsafe.Pointer(b.Top))
	C.free(unsafe.Pointer(b.Bottom))
	C.free(unsafe.Pointer(b.Left))
	C.free(unsafe.Pointer(b.Right))
	C.free(unsafe.Pointer(b.TopLeft))
	C.free(unsafe.Pointer(b.TopRight))
	C.free(unsafe.Pointer(b.BottomLeft))
	C.free(unsafe.Pointer(b.BottomRight))
	C.free(unsafe.Pointer(b.MiddleLeft))
	C.free(unsafe.Pointer(b.MiddleRight))
	C.free(unsafe.Pointer(b.Middle))
	C.free(unsafe.Pointer(b.MiddleTop))
	C.free(unsafe.Pointer(b.MiddleBottom))
}

//export BlockBorder
func BlockBorder() C.CBorder {
	Log(LogLevelDebug, "Created block border")
	return toBorder(lipgloss.BlockBorder())
}

//export DoubleBorder
func DoubleBorder() C.CBorder {
	Log(LogLevelDebug, "Created double border")
	return toBorder(lipgloss.DoubleBorder())
}

//export HiddenBorder
func HiddenBorder() C.CBorder {
	Log(LogLevelDebug, "Created hidden border")
	return toBorder(lipgloss.HiddenBorder())
}

//export InnerHalfBlockBorder
func InnerHalfBlockBorder() C.CBorder {
	Log(LogLevelDebug, "Created inner half block border")
	return toBorder(lipgloss.InnerHalfBlockBorder())
}

//export NormalBorder
func NormalBorder() C.CBorder {
	Log(LogLevelDebug, "Created normal border")
	return toBorder(lipgloss.NormalBorder())
}

//export OuterHalfBlockBorder
func OuterHalfBlockBorder() C.CBorder {
	Log(LogLevelDebug, "Created outer half block border")
	return toBorder(lipgloss.OuterHalfBlockBorder())
}

//export RoundedBorder
func RoundedBorder() C.CBorder {
	Log(LogLevelDebug, "Created rounded border")
	return toBorder(lipgloss.RoundedBorder())
}

//export ThickBorder
func ThickBorder() C.CBorder {
	Log(LogLevelDebug, "Created thick border")
	return toBorder(lipgloss.ThickBorder())
}

//export FreeBorder
func FreeBorder(b C.CBorder) {
	freeBorder(b)
}

//export GetBottomSize
func GetBottomSize(b C.CBorder) C.int {
	border := lipgloss.Border{
		Bottom: C.GoString(b.Bottom),
	}
	return C.int(border.GetBottomSize())
}

//export GetLeftSize
func GetLeftSize(b C.CBorder) C.int {
	border := lipgloss.Border{
		Left: C.GoString(b.Left),
	}
	return C.int(border.GetLeftSize())
}

//export GetRightSize
func GetRightSize(b C.CBorder) C.int {
	border := lipgloss.Border{
		Right: C.GoString(b.Right),
	}
	return C.int(border.GetRightSize())
}

//export GetTopSize
func GetTopSize(b C.CBorder) C.int {
	border := lipgloss.Border{
		Top: C.GoString(b.Top),
	}
	return C.int(border.GetTopSize())
}

//export CreateCustomBorder
func CreateCustomBorder(top, bottom, left, right, topLeft, topRight,
	bottomLeft, bottomRight, middleLeft, middleRight,
	middle, middleTop, middleBottom *C.char) C.CBorder {

	return C.CBorder{
		Top:          C.CString(C.GoString(top)),
		Bottom:       C.CString(C.GoString(bottom)),
		Left:         C.CString(C.GoString(left)),
		Right:        C.CString(C.GoString(right)),
		TopLeft:      C.CString(C.GoString(topLeft)),
		TopRight:     C.CString(C.GoString(topRight)),
		BottomLeft:   C.CString(C.GoString(bottomLeft)),
		BottomRight:  C.CString(C.GoString(bottomRight)),
		MiddleLeft:   C.CString(C.GoString(middleLeft)),
		MiddleRight:  C.CString(C.GoString(middleRight)),
		Middle:       C.CString(C.GoString(middle)),
		MiddleTop:    C.CString(C.GoString(middleTop)),
		MiddleBottom: C.CString(C.GoString(middleBottom)),
	}
}
