package main

/*
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/charmbracelet/lipgloss"
)

// Exported package-level variables and constants
var (
	String   = &StringUtil{}
	Validate = &ValidationUtil{}
	Memory   = &MemoryUtil{
		allocations: make(map[unsafe.Pointer]string),
	}
	Style = &StyleUtil{}
)

const (
	LogLevelError = C.LOG_ERROR
	LogLevelWarn  = C.LOG_WARN
	LogLevelInfo  = C.LOG_INFO
	LogLevelDebug = C.LOG_DEBUG
)

// Exported log function - capitalized for visibility across package
func Log(level int, format string, args ...interface{}) {
	if level <= CurrentLogLevel {
		log.Printf(format, args...)
	}
}

var CurrentLogLevel = LogLevelError

type StringUtil struct{}
type ValidationUtil struct{}
type MemoryUtil struct {
	allocations map[unsafe.Pointer]string
}
type StyleUtil struct{}

//export SetLogLevel
func SetLogLevel(level C.int) {
	CurrentLogLevel = int(level)
}

// Error types
type StyleError struct {
	Op      string
	Message string
}

func (e *StyleError) Error() string {
	return fmt.Sprintf("style error (op=%s): %s", e.Op, e.Message)
}

type ValidationError struct {
	Op      string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error (op=%s): %s", e.Op, e.Message)
}

type MemoryError struct {
	Op      string
	Message string
}

func (e *MemoryError) Error() string {
	return fmt.Sprintf("memory error (op=%s): %s", e.Op, e.Message)
}

// SafeCString converts a Go string to a C string with error handling
func (su *StringUtil) CString(s string) (*C.char, error) {
	cs := C.CString(s)
	if cs == nil {
		return nil, fmt.Errorf("failed to allocate memory for string: %s", s)
	}
	return cs, nil
}

// GoString safely converts a C string to a Go string
func (su *StringUtil) GoString(cs *C.char) string {
	if cs == nil {
		return ""
	}
	return C.GoString(cs)
}

// ToCInt converts a Go bool to a C int
func (su *StringUtil) ToCInt(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

// ToBool converts a C int to a Go bool
func (su *StringUtil) ToBool(i C.int) bool {
	return i != 0
}

// Position validation
func (vu *ValidationUtil) Position(pos float64, name string) error {
	if pos < 0 || pos > 1 {
		return fmt.Errorf("invalid %s position: %f (must be between 0 and 1)", name, pos)
	}
	return nil
}

// Dimension validation
func (vu *ValidationUtil) Dimension(value int, name string) error {
	if value < 0 {
		return fmt.Errorf("invalid %s: %d (must be non-negative)", name, value)
	}
	return nil
}

// Padding validation
func (vu *ValidationUtil) Padding(value int, name string) error {
	if value < 0 {
		return fmt.Errorf("invalid %s padding: %d (must be non-negative)", name, value)
	}
	return nil
}

// Renderer validation
func (vu *ValidationUtil) Renderer(value int, name string) error {
	if value < 0 {
		return fmt.Errorf("invalid %s padding: %d (must be non-negative)", name, value)
	}
	return nil
}

func (vu *ValidationUtil) Color(color string, op string) error {
	if color == "" {
		return &ValidationError{
			Op:      op,
			Message: "empty color string",
		}
	}

	if color[0] == '#' {
		// Validate hex color format
		if len(color) != 4 && len(color) != 7 {
			return &ValidationError{
				Op:      op,
				Message: fmt.Sprintf("invalid hex color format: %s", color),
			}
		}
	}
	return nil
}

// Track allocation
func (mu *MemoryUtil) Track(ptr unsafe.Pointer, desc string) {
	if CurrentLogLevel >= LogLevelDebug {
		mu.allocations[ptr] = desc
		Log(LogLevelDebug, "allocated: %v (%s)", ptr, desc)
	}
}

// Untrack allocation
func (mu *MemoryUtil) Untrack(ptr unsafe.Pointer) {
	if CurrentLogLevel >= LogLevelDebug {
		if desc, ok := mu.allocations[ptr]; ok {
			Log(LogLevelDebug, "freed: %v (%s)", ptr, desc)
			delete(mu.allocations, ptr)
		} else {
			Log(LogLevelWarn, "attempting to free untracked pointer: %v", ptr)
		}
	}
}

// CheckMemoryLeaks returns information about potential memory leaks
func (mu *MemoryUtil) CheckLeaks() (*C.char, error) {
	if len(mu.allocations) == 0 {
		return String.CString("No memory leaks detected")
	}

	var leaks string
	for ptr, desc := range mu.allocations {
		leaks += fmt.Sprintf("Leak: %v (%s)\n", ptr, desc)
	}
	return String.CString(leaks)
}

//export GetMemoryLeaks
func GetMemoryLeaks() *C.char {
	cs, err := Memory.CheckLeaks()
	if err != nil {
		Log(LogLevelError, "Failed to get memory leaks: %v", err)
		defaultCs, _ := String.CString("Error checking memory leaks")
		return defaultCs
	}
	return cs
}

// SafeGet safely retrieves a style from the registry
func (su *StyleUtil) SafeGet(id uint64, op string) (*lipgloss.Style, error) {
	style := styleReg.Get(id)
	if style == nil {
		return nil, fmt.Errorf("style error (op=%s): style not found with ID %d", op, id)
	}
	return style, nil
}

// Register adds a style to the registry
func (su *StyleUtil) Register(style *lipgloss.Style) uint64 {
	return styleReg.Register(style)
}

// Free removes a style from the registry
func (su *StyleUtil) Free(id uint64) {
	styleReg.Remove(id)
}
