package storage

import (
	"time"

	"github.com/satori/uuid"
)

func DefaultReleaseTimer(driver Driver, id uuid.UUID, duration time.Duration) <-chan error {
	ch := make(chan error)
	go func() {
		<-time.NewTimer(duration).C
		ch <- driver.ReleaseObject(id)
	}()
	return ch
}
