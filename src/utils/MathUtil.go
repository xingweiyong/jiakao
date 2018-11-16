package utils

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randowMutex sync.Mutex
)

func GetRandowInt(start, end int) int {
	randowMutex.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end-start)
	randowMutex.Unlock()
	return n
}
