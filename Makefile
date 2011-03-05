include $(GOROOT)/src/Make.inc

TARG=go-git

CGOFILES=git.go
GOFILES=defs.go

CGO_CFLAGS:=`pkg-config --cflags libgit2`
CGO_LDFLAGS=`pkg-config --libs libgit2`

CFLAGS:=`pkg-config --cflags libgit2`
LDFLAGS=`pkg-config --libs libgit2`

CLEANFILES+=defs.go ./tmp ctest

include $(GOROOT)/src/Make.pkg

ctest: clean ctest.c
	gcc -g -O0 -Wall $(CFLAGS) $(LDFLAGS) $@.c -o $@
	./$@

defs.go: defs.c
	godefs -g git defs.c > defs.go

format:
	gofmt -l -w *.go

all: ctest test

defs.go: defs.c
