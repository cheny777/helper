package util

import (
	"errors"
	"sync"
)

var (
	ErrClosed = errors.New("buffer closed")
)

type SafeLoopBuffer struct {
	lock     sync.RWMutex
	size     uint32
	buffer   []interface{}
	writePos uint32
	closed   bool
}

func (b *SafeLoopBuffer) Push(frame interface{}) (uint32, error) {
	var wPos uint32
	if b.isClosed() {
		return wPos, ErrClosed
	}
	b.lock.Lock()
	defer b.lock.Unlock()
	b.buffer[b.writePos%b.size] = frame
	wPos = b.writePos
	b.writePos++
	return wPos, nil
}

func (b *SafeLoopBuffer) Get(index uint32) interface{} {
	if b.isClosed() {
		return ErrClosed
	}
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.buffer[index%b.size]
}

func (b *SafeLoopBuffer) GetWritePos() (uint32, error) {
	if b.isClosed() {
		return 0, ErrClosed
	}
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.writePos, nil
}

func (b *SafeLoopBuffer) isClosed() bool {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.closed
}

func (b *SafeLoopBuffer) Close() error {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.closed = true
	return nil
}

func NewSafeLoopBuffer(size uint32) *SafeLoopBuffer {
	buffer := &SafeLoopBuffer{
		size:     size,
		buffer:   make([]interface{}, size),
		writePos: 0,
	}
	return buffer
}
