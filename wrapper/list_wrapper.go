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

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

// listRegistry manages list instances with thread safety
type listRegistry struct {
	sync.RWMutex
	nextID uint64
	lists  map[uint64]*list.List
}

var listReg = &listRegistry{
	lists: make(map[uint64]*list.List),
}

// Register adds a list to the registry and returns its ID
func (r *listRegistry) Register(l *list.List) uint64 {
	if l == nil {
		return 0
	}
	r.Lock()
	defer r.Unlock()

	id := atomic.AddUint64(&r.nextID, 1)
	r.lists[id] = l
	return id
}

// Get retrieves a list from the registry
func (r *listRegistry) Get(id uint64) *list.List {
	r.RLock()
	defer r.RUnlock()
	return r.lists[id]
}

// Remove deletes a list from the registry
func (r *listRegistry) Remove(id uint64) {
	r.Lock()
	defer r.Unlock()
	delete(r.lists, id)
}

//export NewList
func NewList() C.uint64_t {
	l := list.New()
	return C.uint64_t(listReg.Register(l))
}

//export ListAddItem
func ListAddItem(id C.uint64_t, item *C.char) {
	l := listReg.Get(uint64(id))
	if l == nil {
		return
	}
	l.Item(C.GoString(item))
}

//export ListSetEnumerator
func ListSetEnumerator(id C.uint64_t, enumeratorType C.int) {
	l := listReg.Get(uint64(id))
	if l == nil {
		return
	}
	switch enumeratorType {
	case 0:
		l.Enumerator(list.Bullet)
	case 1:
		l.Enumerator(list.Dash)
	case 2:
		l.Enumerator(list.Alphabet)
	case 3:
		l.Enumerator(list.Arabic)
	case 4:
		l.Enumerator(list.Roman)
	}
}

//export ListSetItemStyle
func ListSetItemStyle(id C.uint64_t, style *C.char) {
	l := listReg.Get(uint64(id))
	if l == nil {
		return
	}
	styled := lipgloss.NewStyle().Foreground(lipgloss.Color(C.GoString(style)))
	l.ItemStyle(styled)
}

//export RenderList
func RenderList(id C.uint64_t) *C.char {
	l := listReg.Get(uint64(id))
	if l == nil {
		return C.CString("")
	}
	result := l.String()
	return C.CString(result)
}

//export FreeList
func FreeList(id C.uint64_t) {
	listReg.Remove(uint64(id))
}
