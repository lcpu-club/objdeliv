package local

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/lcpu-club/objdeliv/storage"
)

func New(path string) *LocalDriver {
	return &LocalDriver{
		path: path,
	}
}

func Factory(conf storage.DriverConfigure) (storage.Driver, error) {
	pathOrigin, ok := conf["path"]
	if !ok {
		return nil, fmt.Errorf("invalid driver configure for LocalDriver: `path` needed")
	}
	path, ok := pathOrigin.(string)
	if !ok {
		return nil, fmt.Errorf("invalid driver configure for LocalDriver: `path` is not string")
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, fs.FileMode(0666))
	}
	if err != nil {
		return nil, err
	}
	return New(path), nil
}

func init() {
	storage.RegisterDriver("LocalDriver", Factory)
}
