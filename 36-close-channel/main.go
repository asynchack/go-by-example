package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("receive more job:", job)
			} else {
				fmt.Println("no more job!")
				done <- true
				return
			}
		}
	}()

	for i := range 3 {
		jobs <- i
	}
	close(jobs)
	fmt.Println("all job sent")

	<-done

	val, ok := <-jobs
	fmt.Println(val, ok)
}
