include $(GOROOT)/src/Make.inc

TARG=go-git

CGOFILES=git.go

GOFILES=defs.go

CGO_CFLAGS:=`pkg-config --cflags libgit2`

CGO_LDFLAGS=`pkg-config --libs libgit2`

CLEANFILES+=defs.go

include $(GOROOT)/src/Make.pkg

defs.go: defs.c
	godefs -g git defs.c > defs.go

format:
	gofmt -l -w *.go
