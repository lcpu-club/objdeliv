// Package local implements a local storage driver
// for objdeliv
package local

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/lcpu-club/objdeliv/storage"
	"github.com/satori/uuid"
)

type LocalDriver struct {
	path string
}

func (d *LocalDriver) GetName() string {
	return "LocalDriver"
}

func (d *LocalDriver) calculatePath(id uuid.UUID) string {
	return filepath.Join(d.path, id.String()+".dat")
}

func (d *LocalDriver) NewObject(id uuid.UUID) (io.WriteCloser, error) {
	path := d.calculatePath(id)
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.FileMode(0666))
}

func (d *LocalDriver) GetObject(id uuid.UUID) (io.ReadCloser, error) {
	path := d.calculatePath(id)
	return os.OpenFile(path, os.O_RDONLY, os.FileMode(0666))
}

func (d *LocalDriver) IsExist(id uuid.UUID) (bool, error) {
	_, err := os.Stat(d.calculatePath(id))
	if os.IsNotExist(err) {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

func (d *LocalDriver) ReleaseObject(id uuid.UUID) error {
	return os.Remove(d.calculatePath(id))
}

func (d *LocalDriver) SetExpire(id uuid.UUID, expire time.Duration) <-chan error {
	return storage.DefaultReleaseTimer(d, id, expire)
}
