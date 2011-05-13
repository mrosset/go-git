include $(GOROOT)/src/Make.inc

TARG=gogit

GOFILES=gogit.go
#CLEANFILES+=./tmp

include $(GOROOT)/src/Make.cmd

test: clean all
	./${TARG} -t
	cd ./tmp; git log
