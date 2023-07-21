package semaphore

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

type Semaphore chan struct{}

func (s Semaphore) Acquire() { s <- struct{}{} }

func (s Semaphore) AcquireAll() {
	for i := 0; i < cap(s); i++ {
		s.Acquire()
	}
}

func (s Semaphore) Release() { <-s }

func (s Semaphore) ReleaseAll() {
	for i := 0; i < cap(s); i++ {
		s.Release()
	}
}

func (s Semaphore) Wait() {
	s.AcquireAll()
	s.ReleaseAll()
}
