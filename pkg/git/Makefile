include $(GOROOT)/src/Make.inc

TARG=git

CGOFILES=git.go
GOFILES=defs.go
CLEANFILES+=./tmp

include $(GOROOT)/src/Make.pkg

defs.go: defs.c
	godefs -g git defs.c > defs.go

format: *.go
	gofmt -l -w *.go
