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

// termios types as defined in bits/termios.h
type CC_t byte
type Speed_t uint
type TCFlag_t uint

const NCCS = 32

type Termios struct {
	C_iflag  TCFlag_t
	C_oflag  TCFlag_t
	C_cflag  TCFlag_t
	C_lflag  TCFlag_t
	C_line   CC_t
	C_cc     [NCCS]CC_t
	C_ispeed Speed_t
	C_ospeed Speed_t
}

// c_cc control characters
const (
	VINTR    = 0
	VQUIT    = 1
	VERASE   = 2
	VKILL    = 3
	VEOF     = 4
	VTIME    = 5
	VMIN     = 6
	VSWTC    = 7
	VSTART   = 8
	VSTOP    = 9
	VSUSP    = 10
	VEOL     = 11
	VREPRINT = 12
	VDISCARD = 13
	VWERASE  = 14
	VLNEXT   = 15
	VEOL2    = 16
)

// c_iflag constants
const (
	IGNBRK  = TCFlag_t(0000001)
	BRKINT  = TCFlag_t(0000002)
	IGNPAR  = TCFlag_t(0000004)
	PARMRK  = TCFlag_t(0000010)
	INPCK   = TCFlag_t(0000020)
	ISTRIP  = TCFlag_t(0000040)
	INLCR   = TCFlag_t(0000100)
	IGNCR   = TCFlag_t(0000200)
	ICRNL   = TCFlag_t(0000400)
	IUCLC   = TCFlag_t(0001000)
	IXON    = TCFlag_t(0002000)
	IXANY   = TCFlag_t(0004000)
	IXOFF   = TCFlag_t(0010000)
	IMAXBEL = TCFlag_t(0020000)
	IUTF8   = TCFlag_t(0040000)
)

// c_oflag constants
const (
	OPOST  = TCFlag_t(0000001)
	OLCUC  = TCFlag_t(0000002)
	ONLCR  = TCFlag_t(0000004)
	OCRNL  = TCFlag_t(0000010)
	ONOCR  = TCFlag_t(0000020)
	ONLRET = TCFlag_t(0000040)
	OFILL  = TCFlag_t(0000100)
	OFDEL  = TCFlag_t(0000200)
	NLDLY  = TCFlag_t(0000400)
	NL0    = TCFlag_t(0000000)
	NL1    = TCFlag_t(0000400)
	CRDLY  = TCFlag_t(0003000)
	CR0    = TCFlag_t(0000000)
	CR1    = TCFlag_t(0001000)
	CR2    = TCFlag_t(0002000)
	CR3    = TCFlag_t(0003000)
	TABDLY = TCFlag_t(0014000)
	TAB0   = TCFlag_t(0000000)
	TAB1   = TCFlag_t(0004000)
	TAB2   = TCFlag_t(0010000)
	TAB3   = TCFlag_t(0014000)
	XTABS  = TCFlag_t(0014000)
	BSDLY  = TCFlag_t(0020000)
	BS0    = TCFlag_t(0000000)
	BS1    = TCFlag_t(0020000)
	FFDLY  = TCFlag_t(0100000)
	FF0    = TCFlag_t(0000000)
	FF1    = TCFlag_t(0100000)
	VTDLY  = TCFlag_t(0040000)
	VT0    = TCFlag_t(0000000)
	VT1    = TCFlag_t(0040000)
)

// c_cflag constants
const (
	CBAUD      = TCFlag_t(0010017)
	B0         = TCFlag_t(0000000)
	B50        = TCFlag_t(0000001)
	B75        = TCFlag_t(0000002)
	B110       = TCFlag_t(0000003)
	B134       = TCFlag_t(0000004)
	B150       = TCFlag_t(0000005)
	B200       = TCFlag_t(0000006)
	B300       = TCFlag_t(0000007)
	B600       = TCFlag_t(0000010)
	B1200      = TCFlag_t(0000011)
	B1800      = TCFlag_t(0000012)
	B2400      = TCFlag_t(0000013)
	B4800      = TCFlag_t(0000014)
	B9600      = TCFlag_t(0000015)
	B19200     = TCFlag_t(0000016)
	B38400     = TCFlag_t(0000017)
	EXTA       = B19200
	EXTB       = B38400
	CSIZE      = TCFlag_t(0000060)
	CS5        = TCFlag_t(0000000)
	CS6        = TCFlag_t(0000020)
	CS7        = TCFlag_t(0000040)
	CS8        = TCFlag_t(0000060)
	CSTOPB     = TCFlag_t(0000100)
	CREAD      = TCFlag_t(0000200)
	PARENB     = TCFlag_t(0000400)
	PARODD     = TCFlag_t(0001000)
	HUPCL      = TCFlag_t(0002000)
	CLOCAL     = TCFlag_t(0004000)
	CBAUDEX    = TCFlag_t(0010000)
	B57600     = TCFlag_t(0010001)
	B115200    = TCFlag_t(0010002)
	B230400    = TCFlag_t(0010003)
	B460800    = TCFlag_t(0010004)
	B500000    = TCFlag_t(0010005)
	B576000    = TCFlag_t(0010006)
	B921600    = TCFlag_t(0010007)
	B1000000   = TCFlag_t(0010010)
	B1152000   = TCFlag_t(0010011)
	B1500000   = TCFlag_t(0010012)
	B2000000   = TCFlag_t(0010013)
	B2500000   = TCFlag_t(0010014)
	B3000000   = TCFlag_t(0010015)
	B3500000   = TCFlag_t(0010016)
	B4000000   = TCFlag_t(0010017)
	__MAX_BAUD = B4000000
	CIBAUD     = TCFlag_t(002003600000)
	CMSPAR     = TCFlag_t(010000000000)
	CRTSCTS    = TCFlag_t(020000000000)
)
// c_lflag constants
const (
	ISIG    = TCFlag_t(0000001)
	ICANON  = TCFlag_t(0000002)
	XCASE   = TCFlag_t(0000004)
	ECHO    = TCFlag_t(0000010)
	ECHOE   = TCFlag_t(0000020)
	ECHOK   = TCFlag_t(0000040)
	ECHONL  = TCFlag_t(0000100)
	NOFLSH  = TCFlag_t(0000200)
	TOSTOP  = TCFlag_t(0000400)
	ECHOCTL = TCFlag_t(0001000)
	ECHOPRT = TCFlag_t(0002000)
	ECHOKE  = TCFlag_t(0004000)
	FLUSHO  = TCFlag_t(0010000)
	PENDIN  = TCFlag_t(0040000)
	IEXTEN  = TCFlag_t(0100000)
)

// ioctl constants
const (
	TCGETS = 0x5401
	TCSETS = 0x5402
)

func Get(fd int, dst *Termios) os.Error {
	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(TCGETS),
		uintptr(unsafe.Pointer(dst)))

	if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
		return err
	}

	if r1 != 0 {
		return os.ErrorString("Error")
	}

	return nil
}

func Set(fd int, src *Termios) os.Error {
	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(TCSETS),
		uintptr(unsafe.Pointer(src)))

	if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
		return err
	}

	if r1 != 0 {
		return os.ErrorString("Error")
	}

	return nil
}
