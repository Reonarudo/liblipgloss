# Compiler settings
CC = clang
CFLAGS = -Wall -Wextra -Iinclude
LDFLAGS = -L. -llipgloss -Wl,-rpath,.

# Go settings
export CGO_CFLAGS="-I$(PWD)/include"
GOBUILD = go build -o liblipgloss.dylib -buildmode=c-shared

# Installation paths
PREFIX ?= /opt/homebrew
EXEC_PREFIX = $(PREFIX)
LIBDIR = $(EXEC_PREFIX)/lib
INCLUDEDIR = $(PREFIX)/include
PKGCONFIGDIR = $(LIBDIR)/pkgconfig

# Targets
all: build test

build:
	$(GOBUILD) wrapper/*.go

test: build
	$(CC) $(CFLAGS) tests/test_lipgloss_wrapper.c -o test_lipgloss_wrapper $(LDFLAGS)
	$(CC) $(CFLAGS) tests/memory_test.c -o memory_test $(LDFLAGS)

clean:
	rm -f liblipgloss.dylib test_lipgloss_wrapper memory_test go.sum liblipgloss.h

install:
	mkdir -p $(DESTDIR)$(LIBDIR)
	mkdir -p $(DESTDIR)$(INCLUDEDIR)
	mkdir -p $(DESTDIR)$(PKGCONFIGDIR)
	
	install -m 0755 liblipgloss.dylib $(DESTDIR)$(LIBDIR)/
	install -m 0644 include/*.h $(DESTDIR)$(INCLUDEDIR)/

	# Install pkg-config file
	echo "prefix=$(PREFIX)" > liblipgloss.pc
	echo "exec_prefix=\$${prefix}" >> liblipgloss.pc
	echo "libdir=\$${exec_prefix}/lib" >> liblipgloss.pc
	echo "includedir=\$${prefix}/include" >> liblipgloss.pc
	echo "" >> liblipgloss.pc
	echo "Name: liblipgloss" >> liblipgloss.pc
	echo "Description: C Wrapper for Lipgloss, a terminal UI styling library" >> liblipgloss.pc
	echo "Version: 0.1.0" >> liblipgloss.pc
	echo "Libs: -L\$${libdir} -llipgloss" >> liblipgloss.pc
	echo "Cflags: -I\$${includedir}" >> liblipgloss.pc
	
	install -m 0644 liblipgloss.pc $(DESTDIR)$(PKGCONFIGDIR)/
	rm -f liblipgloss.pc

