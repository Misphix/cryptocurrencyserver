package flowcontrol

import (
	"sync"
	"time"
)

var sourceTokens = make(map[string]uint)
var maxNumber, interval uint
var lock = sync.Mutex{}

// Init will initialize flow control necessary parameter
func Init(maxNum uint, inter uint, sources []string) {
	maxNumber = maxNum
	interval = inter
	for _, s := range sources {
		sourceTokens[s] = maxNumber
	}
	go addTokens()
}

// HasPermission will check source has token to execute action
func HasPermission(source string) bool {
	lock.Lock()
	defer lock.Unlock()
	if sourceTokens[source] > 0 {
		sourceTokens[source]--
		return true
	}

	return false
}

func addTokens() {
	lock.Lock()
	for key := range sourceTokens {
		if sourceTokens[key] < maxNumber {
			sourceTokens[key]++
		}
	}
	lock.Unlock()
	time.Sleep(time.Duration(interval) * time.Second)
}
