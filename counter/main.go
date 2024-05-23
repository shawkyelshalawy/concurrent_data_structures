package main

import (
	"fmt"
	"sync"
)

const NUMCPUS = 4

type Counter struct {
	global    int
	glock     sync.Mutex
	local     [NUMCPUS]int
	llock     [NUMCPUS]sync.Mutex
	threshold int
}

func (c *Counter) Init(threshold int) {
	c.threshold = threshold
	c.global = 0
	for i := 0; i < NUMCPUS; i++ {
		c.local[i] = 0
	}
}

func (c *Counter) Update(threadID int, amt int) {
	cpu := threadID % NUMCPUS
	c.llock[cpu].Lock()
	c.local[cpu] += amt
	if c.local[cpu] >= c.threshold {

		c.glock.Lock()
		c.global += c.local[cpu]
		c.glock.Unlock()
		c.local[cpu] = 0
	}
	c.llock[cpu].Unlock()
}

func (c *Counter) Get() int {
	c.glock.Lock()
	val := c.global
	c.glock.Unlock()
	return val
}

func main() {
	var counter Counter
	counter.Init(5)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				counter.Update(threadID, 1)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("Approximate Counter Value:", counter.Get())
}
