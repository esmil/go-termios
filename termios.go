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

func (t *Termios) Get(fd int) os.Error {
	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(fd), uintptr(getTERMIOS),
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
		uintptr(fd), uintptr(setTERMIOS),
		uintptr(unsafe.Pointer(t)))

	if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
		return err
	}

	if r1 != 0 {
		return os.ErrorString("Error")
	}

	return nil
}
