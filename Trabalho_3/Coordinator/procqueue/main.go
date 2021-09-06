package procqueue

import (
	"fmt"
	"net"
	"sync"
)

func InitQueue(pi ProcessInfo) *ProcessQueue {
	q := []ProcessInfo{}
	pq := ProcessQueue{queue: &q}
	pq.isEmpty = *sync.NewCond(&pq.mu)
	return &pq
}

type ProcessInfo struct {
	process uint32
	conn    net.Conn
}

type ProcessQueue struct {
	queue   *[]ProcessInfo
	mu      sync.Mutex
	isEmpty sync.Cond
}

func (pq *ProcessQueue) push(e ProcessInfo) {
	pq.mu.Lock()
	q := append(*pq.queue, e)
	pq.queue = &q
	pq.isEmpty.Signal()
	pq.mu.Unlock()
}

func (pq *ProcessQueue) pop() ProcessInfo {
	pq.mu.Lock()
	for len(*pq.queue) == 0 {
		pq.isEmpty.Wait()
	}
	pi := (*pq.queue)[0]
	q := (*pq.queue)[1:]
	pq.queue = &q
	pq.mu.Unlock()
	return pi
}

func (pq *ProcessQueue) print() string {
	var s string
	pq.mu.Lock()
	s = fmt.Sprintf("%v", (*pq.queue))
	pq.mu.Unlock()
	return s
}
