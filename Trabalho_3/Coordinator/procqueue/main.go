package procqueue

import (
	"fmt"
	"net"
	"sync"
)

//Initializes a thread-safe ProcessQueue
func InitQueue() *ProcessQueue {
	q := []ProcessInfo{}
	pq := ProcessQueue{queue: &q}
	pq.processCount = make(map[uint32]int)
	pq.isEmpty = *sync.NewCond(&pq.mu)
	return &pq
}

//Node of queue
//Contains the tcp connection of the process that requested access and it's number
type ProcessInfo struct {
	Process uint32
	Conn    net.Conn
}

//Thread-safe ProcessQueue
// Contains a golang slice to be the queue
// a mutex and a condition variable to control concurrency
// a map to keep how many times a process has exited the queue
type ProcessQueue struct {
	queue        *[]ProcessInfo
	mu           sync.Mutex
	isEmpty      sync.Cond
	processCount map[uint32]int
}

// Lock the mutex to access the queue
// append to the queue
// signal to the condition variable to awake if the is a go routine waiting to pop
// unlock the lock
func (pq *ProcessQueue) Push(e ProcessInfo) {
	pq.mu.Lock()
	q := append(*pq.queue, e)
	pq.queue = &q
	pq.isEmpty.Signal()
	pq.mu.Unlock()
}

// Lock the mutex to access the queue
// checks if the queue is empty and, if it's the case, waits on the condition variable until there is something to pop
// increments the number of times a process has exited the queue and been granteed acess to critical region
// pop the queue
// unlock the lock
func (pq *ProcessQueue) Pop() ProcessInfo {
	pq.mu.Lock()
	for len(*pq.queue) == 0 {
		pq.isEmpty.Wait()
	}
	pi := (*pq.queue)[0]
	q := (*pq.queue)[1:]
	pq.queue = &q
	pq.processCount[pi.Process] += 1
	pq.mu.Unlock()
	return pi
}

// Lock the mutex to access the queue
// gets the amount of times a process has exited the queue and been granteed acess to critical region
// unlock the lock
// returns the amount
func (pq *ProcessQueue) Count(processNumber uint32) int {
	var result int
	pq.mu.Lock()
	result = pq.processCount[processNumber]
	pq.mu.Unlock()
	return result
}

// Lock the mutex to access the queue
// gets a string of the queue
// unlock the lock
// returns the string
func (pq *ProcessQueue) Print() string {
	var s string
	pq.mu.Lock()
	s = fmt.Sprintf("%v", (*pq.queue))
	pq.mu.Unlock()
	return s
}
