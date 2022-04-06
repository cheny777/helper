package util

import "sync"

type FifoQueue struct {
	m      sync.Mutex
	buffer []interface{}
}

func (fq *FifoQueue) Put(x interface{}) {
	fq.m.Lock()
	defer fq.m.Unlock()
	fq.buffer = append(fq.buffer, x)
}

func (fq *FifoQueue) Pop() interface{} {
	fq.m.Lock()
	defer fq.m.Unlock()
	if len(fq.buffer) == 0 {
		return nil
	}
	x := fq.buffer[0]
	fq.buffer = fq.buffer[1:]
	return x
}

func (fq *FifoQueue) Len() int {
	fq.m.Lock()
	defer fq.m.Unlock()
	return len(fq.buffer)
}

func (fq *FifoQueue) Clean() {
	fq.m.Lock()
	defer fq.m.Unlock()
	fq.buffer = fq.buffer[:0]
}
