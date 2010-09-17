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

// termios struct as defined in termios.h
const NCCS = 20

type Termios struct {
	IFlag  uint       // input flags
	OFlag  uint       // output flags
	CFlag  uint       // control flags
	LFlag  uint       // local flags
	CC     [NCCS]byte // control chars
	ISpeed uint       // input speed
	OSpeed uint       // output speed
}

// control characters
const (
	VEOF     = iota // ICANON
	VEOL            // ICANON
	VEOL2           // ICANON together with IEXTEN
	VERASE          // ICANON
	VWERASE         // ICANON together with IEXTEN
	VKILL           // ICANON
	VREPRINT        // ICANON together with IEXTEN
	VERASE2         // ICANON
	VINTR           // ISIG
	VQUIT           // ISIG
	VSUSP           // ISIG
	VDSUSP          // ISIG together with IEXTEN
	VSTART          // IXON, IXOFF
	VSTOP           // IXON, IXOFF
	VLNEXT          // IEXTEN
	VDISCARD        // IEXTEN
	VMIN            // !ICANON
	VTIME           // !ICANON
	VSTATUS         // ICANON together with IEXTEN
)

// _POSIX_VDISABLE	0xff

// input flag constants
const (
	IGNBRK  = 0x00000001 // ignore BREAK condition
	BRKINT  = 0x00000002 // map BREAK to SIGINTR
	IGNPAR  = 0x00000004 // ignore (discard) parity errors
	PARMRK  = 0x00000008 // mark parity and framing errors
	INPCK   = 0x00000010 // enable checking of parity errors
	ISTRIP  = 0x00000020 // strip 8th bit off chars
	INLCR   = 0x00000040 // map NL into CR
	IGNCR   = 0x00000080 // ignore CR
	ICRNL   = 0x00000100 // map CR to NL (ala CRMOD)
	IXON    = 0x00000200 // enable output flow control
	IXOFF   = 0x00000400 // enable input flow control
	IXANY   = 0x00000800 // any char will restart after stop
	IMAXBEL = 0x00002000 // ring bell on input queue full
)

// output flag constants
const (
	OPOST  = 0x00000001 // enable following output processing
	ONLCR  = 0x00000002 // map NL to CR-NL (ala CRMOD)
	TABDLY = 0x00000004 // tab delay mask
	TAB0   = 0x00000000 // no tab delay and expansion
	TAB3   = 0x00000004 // expand tabs to spaces
	OXTABS = TAB3
	ONOEOT = 0x00000008 // discard EOT's (^D) on output
	OCRNL  = 0x00000010 // map CR to NL on output
	ONOCR  = 0x00000020 // no CR output at column 0
	ONLRET = 0x00000040 // NL performs CR function
)

// control flag constants
const (
	CIGNORE    = 0x00000001 // ignore control flags
	CSIZE      = 0x00000300 // character size mask
	CS5        = 0x00000000 // 5 bits (pseudo)
	CS6        = 0x00000100 // 6 bits
	CS7        = 0x00000200 // 7 bits
	CS8        = 0x00000300 // 8 bits
	CSTOPB     = 0x00000400 // send 2 stop bits
	CREAD      = 0x00000800 // enable receiver
	PARENB     = 0x00001000 // parity enable
	PARODD     = 0x00002000 // odd parity, else even
	HUPCL      = 0x00004000 // hang up on last close
	CLOCAL     = 0x00008000 // ignore modem status lines
	CCTS_OFLOW = 0x00010000 // CTS flow control of output
	CRTSCTS    = (CCTS_OFLOW | CRTS_IFLOW)
	CRTS_IFLOW = 0x00020000 // RTS flow control of input
	CDTR_IFLOW = 0x00040000 // DTR flow control of input
	CDSR_OFLOW = 0x00080000 // DSR flow control of output
	CCAR_OFLOW = 0x00100000 // DCD flow control of output
	MDMBUF     = CCAR_OFLOW
)


// local flag constants
const (
	ECHOKE     = 0x00000001 // visual erase for line kill
	ECHOE      = 0x00000002 // visually erase chars
	ECHOK      = 0x00000004 // echo NL after line kill
	ECHO       = 0x00000008 // enable echoing
	ECHONL     = 0x00000010 // echo NL even if ECHO is off
	ECHOPRT    = 0x00000020 // visual erase mode for hardcopy
	ECHOCTL    = 0x00000040 // echo control chars as ^(Char)
	ISIG       = 0x00000080 // enable signals INTR, QUIT, [D]SUSP
	ICANON     = 0x00000100 // canonicalize input lines
	ALTWERASE  = 0x00000200 // use alternate WERASE algorithm
	IEXTEN     = 0x00000400 // enable DISCARD and LNEXT
	EXTPROC    = 0x00000800 // external processing
	TOSTOP     = 0x00400000 // stop background jobs from output
	FLUSHO     = 0x00800000 // output being flushed (state)
	NOKERNINFO = 0x02000000 // no kernel output from VSTATUS
	PENDIN     = 0x20000000 // XXX retype pending input (state)
	NOFLSH     = 0x80000000 // don't flush after interrupt
)

// speed constants
const (
	B0      = 0
	B50     = 50
	B75     = 75
	B110    = 110
	B134    = 134
	B150    = 150
	B200    = 200
	B300    = 300
	B600    = 600
	B1200   = 1200
	B1800   = 1800
	B2400   = 2400
	B4800   = 4800
	B9600   = 9600
	B19200  = 19200
	B38400  = 38400
	B7200   = 7200
	B14400  = 14400
	B28800  = 28800
	B57600  = 57600
	B76800  = 76800
	B115200 = 115200
	B230400 = 230400
	B460800 = 460800
	B921600 = 921600
	EXTA    = 19200
	EXTB    = 38400
)

// ioctl constants
const (
	TIOCGETA = 0x402C7413
	TIOCSETA = 0x802C7414

	getTERMIOS = TIOCGETA
	setTERMIOS = TIOCSETA
)
