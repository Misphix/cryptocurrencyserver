package flowcontrol

import (
	"sync"
	"time"
)

// FlowController implemented token bucket flow control method
type FlowController struct {
	maxNumber, interval uint
	sourceTokens        map[string]uint
	lock                *sync.Mutex
}

// New will new a FlowController
// There is no limitation if one of maxNumber or interval equal to zero
func New(maxNumber uint, interval uint, sources []string) FlowController {
	f := FlowController{maxNumber, interval, make(map[string]uint), &sync.Mutex{}}
	for _, s := range sources {
		f.sourceTokens[s] = maxNumber
	}

	if maxNumber != 0 && interval != 0 {
		go f.addTokens()
	}

	return f
}

// AcquirePermission will check source has token to execute action
func (f FlowController) AcquirePermission(source string) bool {
	if f.maxNumber == 0 || f.interval == 0 {
		return true
	}

	f.lock.Lock()
	defer f.lock.Unlock()
	if f.sourceTokens[source] > 0 {
		f.sourceTokens[source]--
		return true
	}

	return false
}

func (f FlowController) addTokens() {
	f.lock.Lock()
	for key := range f.sourceTokens {
		if f.sourceTokens[key] < f.maxNumber {
			f.sourceTokens[key]++
		}
	}
	f.lock.Unlock()
	time.Sleep(time.Duration(f.interval) * time.Second)
}
