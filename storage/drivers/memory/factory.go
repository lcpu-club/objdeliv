package memory

import (
	"github.com/lcpu-club/objdeliv/storage"
)

func New() *MemoryDriver {
	return &MemoryDriver{}
}

func Factory(conf storage.DriverConfigure) (storage.Driver, error) {
	return New(), nil
}

func init() {
	storage.RegisterDriver("MemoryDriver", Factory)
}
