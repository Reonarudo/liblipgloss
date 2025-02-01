package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/charmbracelet/lipgloss"
)

// styleRegistry manages style instances with thread safety
type styleRegistry struct {
	sync.RWMutex
	nextID uint64
	styles map[uint64]*lipgloss.Style
}

var styleReg = &styleRegistry{
	styles: make(map[uint64]*lipgloss.Style),
}

// RegistryError represents an error in registry operations
type RegistryError struct {
	Op      string
	ID      uint64
	Message string
}

func (e *RegistryError) Error() string {
	return fmt.Sprintf("registry error (op=%s, id=%d): %s", e.Op, e.ID, e.Message)
}

// Register adds a style to the registry and returns its ID
func (r *styleRegistry) Register(style *lipgloss.Style) uint64 {
	if style == nil {
		Log(LogLevelError, "Attempted to register nil style")
		return 0
	}

	r.Lock()
	defer r.Unlock()

	id := atomic.AddUint64(&r.nextID, 1)
	r.styles[id] = style
	Log(LogLevelDebug, "Registered new style with ID: %d", id)
	return id
}

// Get retrieves a style from the registry
func (r *styleRegistry) Get(id uint64) *lipgloss.Style {
	r.RLock()
	defer r.RUnlock()

	style, exists := r.styles[id]
	if !exists {
		Log(LogLevelError, "Style not found with ID: %d", id)
		return nil
	}
	return style
}

// Remove deletes a style from the registry
func (r *styleRegistry) Remove(id uint64) {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.styles[id]; exists {
		delete(r.styles, id)
		Log(LogLevelDebug, "Removed style with ID: %d", id)
	} else {
		Log(LogLevelWarn, "Attempted to remove non-existent style with ID: %d", id)
	}
}

// GetStats returns statistics about the registry
func (r *styleRegistry) GetStats() string {
	r.RLock()
	defer r.RUnlock()

	return fmt.Sprintf("Total styles: %d, Next ID: %d", len(r.styles), r.nextID)
}

//export NewStyle
func NewStyle() C.uint64_t {
	style := lipgloss.NewStyle()
	id := styleReg.Register(&style)
	Log(LogLevelDebug, "Created new style with ID: %d", id)
	return C.uint64_t(id)
}

//export CopyStyle
func CopyStyle(id C.uint64_t) C.uint64_t {
	style := styleReg.Get(uint64(id))
	if style == nil {
		Log(LogLevelError, "CopyStyle failed: source style not found with ID: %d", uint64(id))
		return 0
	}

	newStyle := style.Copy()
	newID := styleReg.Register(&newStyle)
	Log(LogLevelDebug, "Copied style %d to new style with ID: %d", uint64(id), newID)
	return C.uint64_t(newID)
}

//export FreeStyle
func FreeStyle(id C.uint64_t) {
	styleReg.Remove(uint64(id))
}

//export FreeString
func FreeString(str *C.char) {
	if str != nil {
		Memory.Untrack(unsafe.Pointer(str))
		C.free(unsafe.Pointer(str))
	}
}

//export GetStyleStats
func GetStyleStats() *C.char {
	stats := styleReg.GetStats()
	cs, err := String.CString(stats)
	if err != nil {
		Log(LogLevelError, "GetStyleStats memory allocation error: %v", err)
		return C.CString("Error getting stats")
	}
	Memory.Track(unsafe.Pointer(cs), "style stats string")
	return cs
}

// Helper function to safely get a style with error handling
func getStyleSafe(id uint64, op string) (*lipgloss.Style, error) {
	style := styleReg.Get(id)
	if style == nil {
		return nil, &RegistryError{
			Op:      op,
			ID:      id,
			Message: "style not found",
		}
	}
	return style, nil
}
