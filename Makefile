include $(GOROOT)/src/Make.inc

TARG=termios
GOFILES=\
	$(GOOS).go\
	termios.go\

include $(GOROOT)/src/Make.pkg
