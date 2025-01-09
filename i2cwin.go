//go:build windows

package i2c

func ioctl(fd, cmd, arg uintptr) error {
	lg.Errorf("i2c operations are not supported on Windows")
	return nil
}
