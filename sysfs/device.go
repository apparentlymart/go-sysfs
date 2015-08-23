package sysfs

import (
	"os"
	"sync"
)

type Device struct {
	dir   *os.File
	mutex sync.Mutex
}
