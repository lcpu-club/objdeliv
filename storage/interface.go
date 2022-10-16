// Package storage provides utilities and facades to
// interact with storage engines
package storage

import (
	"fmt"
	"io"
	"time"

	"github.com/satori/uuid"
)

var ErrNotExist error = fmt.Errorf("object not exist")
var ErrAlreadyExist error = fmt.Errorf("object already exist")

// Driver defines storage driver interface
type Driver interface {
	// GetName returns the name of the interface
	GetName() string
	// NewObject creates a new object in the storage
	NewObject(id uuid.UUID) (io.WriteCloser, error)
	// GetObject retrieves an object in the storage
	GetObject(id uuid.UUID) (io.ReadCloser, error)
	// IsExist returns if an object exists
	IsExist(id uuid.UUID) (bool, error)
	// ReleaseObject releases an object in the storage
	ReleaseObject(id uuid.UUID) error
	// SetExpire sets an expire duration for an exact object
	SetExpire(id uuid.UUID, expire time.Duration) <-chan error
}
