package syncprim

import (
	"sync/atomic"
)

var lock uint32 = 0

func Aquire() {
	for !atomic.CompareAndSwapUint32(&lock, 0, 1) {
	}
}

func Release() {
	lock = 0
}
