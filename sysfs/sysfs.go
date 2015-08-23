package sysfs

import (
	"os"
	"sync"
)

type SysFS struct {
	rootDir *os.File
	mutex   sync.Mutex
}

// Open looks for a sysfs mounted at /sys and returns a SysFS object referring
// to it.
//
// Returns an error if the sysfs filesystem can't be opened for some reason.
func Open() (*SysFS, error) {
	return OpenAt("/sys")
}

// OpenAt is like Open but allows accessing a sysfs mounted somewhere other
// than /sys. Non-standard mount paths are not worth supporting in most
// applications since all reasonable distributions mount sysfs at /sys.
func OpenAt(rootPath string) (*SysFS, error) {
	rootDir, err := os.Open(rootPath)
	if err != nil {
		return nil, err
	}

	return &SysFS{
		rootDir: rootDir,
	}, nil
}

func (fs *SysFS) Close() error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()
	return fs.rootDir.Close()
}

func (fs *SysFS) OpenDevice(path string) (*Device, error) {
	devDir, err := openChildDir(fs.rootDir, path[1:])
	if err != nil {
		return nil, err
	}

	// TODO: See if it's actually a device directory, by checking if we
	// can find a "subsystem" symlink inside.

	return &Device{
		dir: devDir,
	}, nil
}
