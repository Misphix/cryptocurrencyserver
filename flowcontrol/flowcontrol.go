package flowcontrol

import (
	"sync"
	"time"
)

// FlowController is an API provider side flow control system.
// It implemented token bucket flow control method
type FlowController struct {
	maxNumber uint
	interval  uint
	tokens    uint
	lock      *sync.Mutex
}

// New will new a FlowController
// There is no limitation if one of maxNumber or interval equal to zero
func New(maxNumber uint, interval uint) *FlowController {
	f := &FlowController{
		maxNumber: maxNumber,
		interval:  interval,
		tokens:    maxNumber,
		lock:      &sync.Mutex{},
	}

	if maxNumber != 0 && interval != 0 {
		go f.addTokens()
	}

	return f
}

// AcquirePermission will check if tokens is enough to execute action
func (f *FlowController) AcquirePermission() bool {
	if f.maxNumber == 0 || f.interval == 0 {
		return true
	}

	f.lock.Lock()
	defer f.lock.Unlock()
	if f.tokens > 0 {
		f.tokens--
		return true
	}

	return false
}

func (f *FlowController) addTokens() {
	time.Sleep(time.Duration(f.interval) * time.Second)
	f.lock.Lock()
	if f.tokens < f.maxNumber {
		f.tokens++
	}

	f.lock.Unlock()
}
