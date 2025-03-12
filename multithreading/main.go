package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started jobs %d\n", id, j)
		time.Sleep(1 * time.Second)
		results <- j * 2
		fmt.Printf("Worker %d finished jobs %d\n", id, j)
	}
}

func producer(id int, data chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(id) * time.Second)
		data <- id
		fmt.Printf("Producer %d produced %d\n", id, i)
	}
}

func consumer(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range data {
		fmt.Printf("Consumer %d consumed %d\n", id, d)
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	data := make(chan int)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 3; i++ {
		go producer(i, data)
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go consumer(i, data, &wg)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result:", result)
	}

	time.Sleep(10 * time.Second)
	close(data)
}
