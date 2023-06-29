package ctxbind

import (
	"context"
	"sync"
	"time"
)

const backgroundID = 0

var (
	contextMu  sync.RWMutex
	lastID     = backgroundID
	contextMap = make(map[int]context.Context)
	cancelMap  = make(map[int]context.CancelFunc)
)

func BackgroundContext() int {
	return backgroundID
}

func WithCancel(parentID int) (id int) {
	contextMu.Lock()
	defer contextMu.Unlock()

	var alreadyUsed bool
	for alreadyUsed || id == backgroundID {
		id = generateId()
		_, alreadyUsed = contextMap[id]
	}

	contextMap[id], cancelMap[id] = context.WithCancel(getContext(parentID))
	return id
}

func GetContext(id int) context.Context {
	if id == backgroundID {
		return context.Background()
	}

	contextMu.RLock()
	defer contextMu.RUnlock()

	return getContext(id)
}

func getContext(id int) context.Context {
	if id == backgroundID {
		return context.Background()
	}

	return contextMap[id]
}

func Deadline(id int) (deadline time.Time, ok bool) { return GetContext(id).Deadline() }
func Done(id int) <-chan struct{}                   { return GetContext(id).Done() }
func Err(id int) error                              { return GetContext(id).Err() }
func Value(id int, key any) any                     { return GetContext(id).Value(key) }
func CancelContext(id int)                          { GetCancelFunction(id)() }

func Unregister(id int) {
	if id == backgroundID {
		return
	}

	contextMu.Lock()
	defer contextMu.Unlock()

	delete(contextMap, id)
	delete(cancelMap, id)
}

func Destroy() {
	contextMu.Lock()
	defer contextMu.Unlock()

	for id := range contextMap {
		if id == backgroundID {
			continue
		}
		delete(contextMap, id)
	}

	for id := range cancelMap {
		delete(cancelMap, id)
	}
}

func GetCancelFunction(id int) context.CancelFunc {
	contextMu.RLock()
	defer contextMu.RUnlock()

	return cancelMap[id]
}

func generateId() int {
	// TODO: add mutex locked assert
	lastID++
	return lastID
}
