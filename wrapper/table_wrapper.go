package main

/*
#include <stdlib.h>
#include <stdint.h>
#include "lipgloss_types.h"
*/
import "C"
import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// tableRegistry manages table instances with thread safety
type tableRegistry struct {
	sync.RWMutex
	nextID uint64
	tables map[uint64]*table.Table
}

var tableReg = &tableRegistry{
	tables: make(map[uint64]*table.Table),
}

// Register adds a table to the registry and returns its ID
func (r *tableRegistry) Register(t *table.Table) uint64 {
	if t == nil {
		return 0
	}
	r.Lock()
	defer r.Unlock()

	id := atomic.AddUint64(&r.nextID, 1)
	r.tables[id] = t
	return id
}

// Get retrieves a table from the registry
func (r *tableRegistry) Get(id uint64) *table.Table {
	r.RLock()
	defer r.RUnlock()
	return r.tables[id]
}

// Remove deletes a table from the registry
func (r *tableRegistry) Remove(id uint64) {
	r.Lock()
	defer r.Unlock()
	delete(r.tables, id)
}

//export NewTable
func NewTable() C.uint64_t {
	t := table.New()
	return C.uint64_t(tableReg.Register(t))
}

//export TableAddHeaders
func TableAddHeaders(id C.uint64_t, headers **C.char, count C.int) {
	t := tableReg.Get(uint64(id))
	if t == nil {
		return
	}
	goHeaders := make([]string, int(count))
	for i := 0; i < int(count); i++ {
		goHeaders[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(headers)) + uintptr(i)*unsafe.Sizeof(*headers))))
	}
	t.Headers(goHeaders...)
}

//export TableAddRow
func TableAddRow(id C.uint64_t, row **C.char, count C.int) {
	t := tableReg.Get(uint64(id))
	if t == nil {
		return
	}
	goRow := make([]string, int(count))
	for i := 0; i < int(count); i++ {
		goRow[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(row)) + uintptr(i)*unsafe.Sizeof(*row))))
	}
	t.Row(goRow...)
}

//export TableSetWidth
func TableSetWidth(id C.uint64_t, width C.int) {
	t := tableReg.Get(uint64(id))
	if t != nil {
		t.Width(int(width))
	}
}

//export TableSetHeight
func TableSetHeight(id C.uint64_t, height C.int) {
	t := tableReg.Get(uint64(id))
	if t != nil {
		t.Height(int(height))
	}
}

//export TableSetBorder
func TableSetBorder(id C.uint64_t, borderType C.int) {
	t := tableReg.Get(uint64(id))
	if t == nil {
		return
	}
	switch borderType {
	case 0:
		t.Border(lipgloss.NormalBorder())
	case 1:
		t.Border(lipgloss.RoundedBorder())
	case 2:
		t.Border(lipgloss.ThickBorder())
	}
}

//export RenderTable
func RenderTable(id C.uint64_t) *C.char {
	t := tableReg.Get(uint64(id))
	if t == nil {
		return C.CString("")
	}
	result := t.Render()
	return C.CString(result)
}

//export FreeTable
func FreeTable(id C.uint64_t) {
	tableReg.Remove(uint64(id))
}
