package async

import "sync/atomic"

type Counter struct {
	noCopy noCopy //nolint:structcheck
	val    uint32
}

func (c *Counter) Increment() uint32 {
	return atomic.AddUint32(&c.val, 1)
}

func (c *Counter) Decrement() uint32 {
	return atomic.AddUint32(&c.val, ^uint32(0))
}

func (c *Counter) Reset() {
	atomic.StoreUint32(&c.val, 0)
}

func (c *Counter) Get() uint32 {
	return atomic.LoadUint32(&c.val)
}
