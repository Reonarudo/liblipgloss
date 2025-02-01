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
	"github.com/charmbracelet/lipgloss/tree"
)

// treeRegistry manages tree instances with thread safety
type treeRegistry struct {
	sync.RWMutex
	nextID uint64
	trees  map[uint64]*tree.Tree
}

var treeReg = &treeRegistry{
	trees: make(map[uint64]*tree.Tree),
}

// Register adds a tree to the registry and returns its ID
func (r *treeRegistry) Register(t *tree.Tree) uint64 {
	if t == nil {
		return 0
	}
	r.Lock()
	defer r.Unlock()

	id := atomic.AddUint64(&r.nextID, 1)
	r.trees[id] = t
	return id
}

// Get retrieves a tree from the registry
func (r *treeRegistry) Get(id uint64) *tree.Tree {
	r.RLock()
	defer r.RUnlock()
	return r.trees[id]
}

// Remove deletes a tree from the registry
func (r *treeRegistry) Remove(id uint64) {
	r.Lock()
	defer r.Unlock()
	delete(r.trees, id)
}

//export NewTree
func NewTree() C.uint64_t {
	t := tree.New()
	return C.uint64_t(treeReg.Register(t))
}

//export TreeAddChildValue
func TreeAddChildValue(parentID C.uint64_t, value *C.char) {
	parent := treeReg.Get(uint64(parentID))
	if parent == nil {
		return
	}
	parent.Child(C.GoString(value))
}

//export TreeAddChildTree
func TreeAddChildTree(parentID C.uint64_t, childID C.uint64_t) {
	parent := treeReg.Get(uint64(parentID))
	child := treeReg.Get(uint64(childID))
	if parent == nil || child == nil {
		return
	}
	parent.Child(child)
}

//export TreeSetEnumerator
func TreeSetEnumerator(id C.uint64_t, enumType C.int) {
	t := treeReg.Get(uint64(id))
	if t == nil {
		return
	}
	switch enumType {
	case 0:
		t.Enumerator(tree.DefaultEnumerator)
	case 1:
		t.Enumerator(tree.RoundedEnumerator)
	}
}

//export TreeSetIndenter
func TreeSetIndenter(id C.uint64_t, indentType C.int) {
	t := treeReg.Get(uint64(id))
	if t == nil {
		return
	}
	switch indentType {
	case 0:
		t.Indenter(tree.DefaultIndenter)
	case 1:
		t.Indenter(func(children tree.Children, index int) string {
			return "    "
		})
	}
}

//export TreeSetItemStyle
func TreeSetItemStyle(id C.uint64_t, style *C.char) {
	t := treeReg.Get(uint64(id))
	if t == nil {
		return
	}
	styled := lipgloss.NewStyle().Foreground(lipgloss.Color(C.GoString(style)))
	t.ItemStyle(styled)
}

//export RenderTree
func RenderTree(id C.uint64_t) *C.char {
	t := treeReg.Get(uint64(id))
	if t == nil {
		return C.CString("(empty tree)")
	}
	result := t.String()
	if result == "" {
		result = "(empty tree)"
	}
	return C.CString(result)
}

//export FreeTree
func FreeTree(id C.uint64_t) {
	treeReg.Remove(uint64(id))
}
