package sysfs

import (
	"os"
	"syscall"
)

func openChild(dir *os.File, name string, flags int, mode uint32) (*os.File, error) {
	dirFd := int(dir.Fd())
	childFd, err := syscall.Openat(dirFd, name, flags, mode)
	if err != nil {
		return nil, err
	}
	return os.NewFile(uintptr(childFd), name), nil
}

func openChildRO(dir *os.File, name string) (*os.File, error) {
	return openChild(dir, name, syscall.O_RDONLY, 0666)
}

func openChildDir(dir *os.File, name string) (*os.File, error) {
	return openChild(dir, name, syscall.O_DIRECTORY, 0666)
}
