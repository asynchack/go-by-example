package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 读请求容器，要读的key，以及接收指回来的chan
type readOp struct {
	key  int
	resp chan int
}

// 写请求容器，要写的key，val，以及写成功的响应flag
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
func main() {
	// 记录读写请求次数，
	var readOps uint64
	var writeOps uint64
	// 读写交流的chan 无buffer

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// stateful goroutine，数据只有一份，只有一个goroutine管理

	go func() {
		// 数据存储容器
		datas := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- datas[read.key]
			case write := <-writes:
				datas[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 写
	for i := 0; i < 10; i++ {
		go func() {
			for {
				writeOp := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- writeOp
				<-writeOp.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}
	// 读

	for i := 0; i < 100; i++ {
		go func() {
			for {
				readOp := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- readOp
				<-readOp.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("reads: ", atomic.LoadUint64(&readOps))
	fmt.Println("writes: ", atomic.LoadUint64(&writeOps))
}
