/*
 * This file is part of go-termios.
 *
 * go-termios is free software: you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of
 * the License, or(at your option) any later version.
 *
 * go-termios is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with go-termios.  If not, see <http://www.gnu.org/licenses/>.
 */
package termios

import (
	"os"
	"syscall"
	"unsafe"
)

// termios struct as defined in bits/termios.h
const NCCS = 32

type Termios struct {
	IFlag  uint
	OFlag  uint
	CFlag  uint
	LFlag  uint
	Line   byte
	CC     [NCCS]byte
	ISpeed uint
	OSpeed uint
}

// c_cc control characters
const (
	VINTR = iota
	VQUIT
	VERASE
	VKILL
	VEOF
	VTIME
	VMIN
	VSWTC
	VSTART
	VSTOP
	VSUSP
	VEOL
	VREPRINT
	VDISCARD
	VWERASE
	VLNEXT
	VEOL2
)

// c_iflag constants
const (
	IGNBRK  = 0000001
	BRKINT  = 0000002
	IGNPAR  = 0000004
	PARMRK  = 0000010
	INPCK   = 0000020
	ISTRIP  = 0000040
	INLCR   = 0000100
	IGNCR   = 0000200
	ICRNL   = 0000400
	IUCLC   = 0001000
	IXON    = 0002000
	IXANY   = 0004000
	IXOFF   = 0010000
	IMAXBEL = 0020000
	IUTF8   = 0040000
)

// c_oflag constants
const (
	OPOST  = 0000001
	OLCUC  = 0000002
	ONLCR  = 0000004
	OCRNL  = 0000010
	ONOCR  = 0000020
	ONLRET = 0000040
	OFILL  = 0000100
	OFDEL  = 0000200
	NLDLY  = 0000400
	NL0    = 0000000
	NL1    = 0000400
	CRDLY  = 0003000
	CR0    = 0000000
	CR1    = 0001000
	CR2    = 0002000
	CR3    = 0003000
	TABDLY = 0014000
	TAB0   = 0000000
	TAB1   = 0004000
	TAB2   = 0010000
	TAB3   = 0014000
	XTABS  = 0014000
	BSDLY  = 0020000
	BS0    = 0000000
	BS1    = 0020000
	FFDLY  = 0100000
	FF0    = 0000000
	FF1    = 0100000
	VTDLY  = 0040000
	VT0    = 0000000
	VT1    = 0040000
)

// c_cflag constants
const (
	CBAUD      = 0010017
	B0         = 0000000
	B50        = 0000001
	B75        = 0000002
	B110       = 0000003
	B134       = 0000004
	B150       = 0000005
	B200       = 0000006
	B300       = 0000007
	B600       = 0000010
	B1200      = 0000011
	B1800      = 0000012
	B2400      = 0000013
	B4800      = 0000014
	B9600      = 0000015
	B19200     = 0000016
	B38400     = 0000017
	EXTA       = B19200
	EXTB       = B38400
	CSIZE      = 0000060
	CS5        = 0000000
	CS6        = 0000020
	CS7        = 0000040
	CS8        = 0000060
	CSTOPB     = 0000100
	CREAD      = 0000200
	PARENB     = 0000400
	PARODD     = 0001000
	HUPCL      = 0002000
	CLOCAL     = 0004000
	CBAUDEX    = 0010000
	B57600     = 0010001
	B115200    = 0010002
	B230400    = 0010003
	B460800    = 0010004
	B500000    = 0010005
	B576000    = 0010006
	B921600    = 0010007
	B1000000   = 0010010
	B1152000   = 0010011
	B1500000   = 0010012
	B2000000   = 0010013
	B2500000   = 0010014
	B3000000   = 0010015
	B3500000   = 0010016
	B4000000   = 0010017
	__MAX_BAUD = B4000000
	CIBAUD     = 002003600000
	CMSPAR     = 010000000000
	CRTSCTS    = 020000000000
)

// c_lflag constants
const (
	ISIG    = 0000001
	ICANON  = 0000002
	XCASE   = 0000004
	ECHO    = 0000010
	ECHOE   = 0000020
	ECHOK   = 0000040
	ECHONL  = 0000100
	NOFLSH  = 0000200
	TOSTOP  = 0000400
	ECHOCTL = 0001000
	ECHOPRT = 0002000
	ECHOKE  = 0004000
	FLUSHO  = 0010000
	PENDIN  = 0040000
	IEXTEN  = 0100000
)

// ioctl constants
const (
	TCGETS = 0x5401
	TCSETS = 0x5402
)

func (t *Termios) Get(fd int) os.Error {
	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(TCGETS),
		uintptr(unsafe.Pointer(t)))

	if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
		return err
	}

	if r1 != 0 {
		return os.ErrorString("Error")
	}

	return nil
}

func (t *Termios) Set(fd int) os.Error {
	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(TCSETS),
		uintptr(unsafe.Pointer(t)))

	if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
		return err
	}

	if r1 != 0 {
		return os.ErrorString("Error")
	}

	return nil
}
