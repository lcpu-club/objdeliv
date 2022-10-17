package memory

import (
	"github.com/lcpu-club/objdeliv/storage"
	"github.com/satori/uuid"
)

func New() *MemoryDriver {
	return &MemoryDriver{
		data: make(map[uuid.UUID]string),
	}
}

func Factory(conf storage.DriverConfigure) (storage.Driver, error) {
	return New(), nil
}

func init() {
	storage.RegisterDriver("MemoryDriver", Factory)
}
