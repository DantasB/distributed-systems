package syncprim

import (
	"sync/atomic"
)

var lock uint32 = 0

// Uses the atomic package of go with guarantees atomicity for its functions
// CompareAndSwapUint32(*prt,old,new ) checks if the variable pointed by ptr has value
// equals to old and if true set it to new and return true
//else just return false
func Aquire() {
	for !atomic.CompareAndSwapUint32(&lock, 0, 1) {
	}
}

//Release the lock setting the variable to 0
//Uses atomic write, because it's recommended by the library
func Release() {
	atomic.StoreUint32(&lock, 0)
}
