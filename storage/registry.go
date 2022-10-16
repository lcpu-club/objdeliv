package storage

import (
	"fmt"
)

type DriverConfigure map[string]interface{}
type DriverFactory func(conf DriverConfigure) (Driver, error)

var drivers map[string]DriverFactory

var ErrDriverNotFound error = fmt.Errorf("Driver not found")

func NewDriver(name string, conf DriverConfigure) (Driver, error) {
	driverFactory, ok := drivers[name]
	if !ok {
		return nil, ErrDriverNotFound
	}
	return driverFactory(conf)
}

func RegisterDriver(name string, factory DriverFactory) {
	drivers[name] = factory
}

func RemoveDriver(name string) {
	delete(drivers, name)
}
