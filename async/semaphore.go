package async

func NewSemaphore(size int) Semaphore {
	if size <= 0 {
		panic("invalid size, should be grater then 0")
	}
	return make(Semaphore, size)
}

type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) AcquireSelect() chan<- struct{} {
	return s
}

func (s Semaphore) AcquireAll() {
	for i := 0; i < cap(s); i++ {
		s.Acquire()
	}
}

func (s Semaphore) Release() {
	<-s
}

func (s Semaphore) ReleaseSelect() <-chan struct{} {
	return s
}

func (s Semaphore) ReleaseAll() {
	for i := 0; i < cap(s); i++ {
		s.Release()
	}
}
