//go:build !windows

package i2c

import "golang.org/x/sys/unix"

func ioctl(fd, cmd, arg uintptr) error {
	_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	if err != 0 {
		return err
	}
	return nil
}
