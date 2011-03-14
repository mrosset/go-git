include $(GOROOT)/src/Make.inc

export LD_LIBRARY_PATH := libgit2/build/shared

TARG=git

CGOFILES=git.go
GOFILES=defs.go

CGO_CFLAGS:=`pkg-config --cflags libgit2`
CGO_LDFLAGS=`pkg-config --libs libgit2`

CFLAGS:=`pkg-config --cflags libgit2`
LDFLAGS=`pkg-config --libs libgit2`

CLEANFILES+=defs.go ./tmp ctest

.PHONY: libgit2 libgit2clean

include $(GOROOT)/src/Make.pkg

libgit2:
	make -C $@

ctest: clean ctest.c
	gcc -g -O0 -Wall $(CFLAGS) $(LDFLAGS) $@.c -o $@
	./$@

defs.go: defs.c
	godefs -g git defs.c > defs.go

format: *.go
	gofmt -l -w *.go

all: test

defs.go: defs.c
