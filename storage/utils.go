package storage

import (
	"io"
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

type CallbackWriteCloser struct {
	w        io.Writer
	callback func(w io.Writer) error
}

func (cwc *CallbackWriteCloser) Write(p []byte) (n int, err error) {
	return cwc.w.Write(p)
}

func (cwc *CallbackWriteCloser) Close() error {
	return cwc.callback(cwc.w)
}

func NewCallbackWriteCloser(w io.Writer, callback func(w io.Writer) error) io.WriteCloser {
	return &CallbackWriteCloser{
		w:        w,
		callback: callback,
	}
}
