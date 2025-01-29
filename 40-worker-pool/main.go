package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// fmt.Println("id: ", id, "working on ", j)
		time.Sleep(time.Second)
		fmt.Println("id: ", id, "work done!", j)

		results <- j * 2
	}
}

func main() {
	const numJobs = 5

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := range numJobs {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}

}
