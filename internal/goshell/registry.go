package goshell

import (
	"os"
	"sync"
)

type RegistryFile struct {
	File  *os.File
	Mutex sync.Mutex
}
