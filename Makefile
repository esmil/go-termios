include $(GOROOT)/src/Make.$(GOARCH)

TARG=termios
GOFILES=\
	termios.go\

include $(GOROOT)/src/Make.pkg
