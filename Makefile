include $(GOROOT)/src/Make.inc

TARG=go-git

CGOFILES=git.go

GOFILES=defs.go

CGO_CFLAGS:=`pkg-config --cflags libgit2`

CGO_LDFLAGS=`pkg-config --libs libgit2`

CLEANFILES+=defs.go

include $(GOROOT)/src/Make.pkg

defs.go: git_defs.c
	godefs -g git git_defs.c > defs.go

format:
	gofmt -l -w *.go

all: format test defs.go
