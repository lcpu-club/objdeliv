// Package memory implements a memory storage driver
// for objdeliv
package memory

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/lcpu-club/objdeliv/storage"
	"github.com/satori/uuid"
)

type MemoryDriver struct {
	data map[uuid.UUID]string
}

func (d *MemoryDriver) GetName() string {
	return "MemoryDriver"
}

func (d *MemoryDriver) NewObject(id uuid.UUID) (io.WriteCloser, error) {
	builder := &strings.Builder{}
	return storage.NewCallbackWriteCloser(builder, func(w io.Writer) error {
		b, ok := w.(*strings.Builder)
		if !ok {
			return fmt.Errorf("invalid writer type: not strings.Builder")
		}
		d.data[id] = b.String()
		return nil
	}), nil
}

func (d *MemoryDriver) GetObject(id uuid.UUID) (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader(d.data[id])), nil
}

func (d *MemoryDriver) IsExist(id uuid.UUID) (bool, error) {
	_, ok := d.data[id]
	return ok, nil
}

func (d *MemoryDriver) ReleaseObject(id uuid.UUID) error {
	delete(d.data, id)
	return nil
}

func (d *MemoryDriver) SetExpire(id uuid.UUID, expire time.Duration) <-chan error {
	return storage.DefaultReleaseTimer(d, id, expire)
}
