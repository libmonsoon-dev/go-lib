package async

import "sync/atomic"

type IdGenerator struct {
	noCopy  noCopy //nolint:structcheck
	counter int32
}

func (c *IdGenerator) Generate() int32 {
	for {
		val := atomic.LoadInt32(&c.counter)
		if atomic.CompareAndSwapInt32(&c.counter, val, val+1) {
			return val
		}
	}
}

func (c *IdGenerator) Reset() {
	atomic.StoreInt32(&c.counter, 0)
}

func (c *IdGenerator) Current() int32 {
	return atomic.LoadInt32(&c.counter)
}
