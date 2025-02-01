# Compiler settings
CC = clang
CFLAGS = -Wall -Wextra -Iinclude
LDFLAGS = -L. -llipgloss -Wl,-rpath,.

# Go settings
export CGO_CFLAGS="-I$(PWD)/include"
GOBUILD = go build -o liblipgloss.dylib -buildmode=c-shared

# Targets
all: build test

build:
	$(GOBUILD) wrapper/*.go

test: build
	$(CC) $(CFLAGS) tests/test_lipgloss_wrapper.c -o test_lipgloss_wrapper $(LDFLAGS)
	$(CC) $(CFLAGS) tests/memory_test.c -o memory_test $(LDFLAGS)

clean:
	rm -f liblipgloss.dylib test_lipgloss_wrapper memory_test go.sum liblipgloss.h