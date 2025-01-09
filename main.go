package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// func printNumbers() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 		time.Sleep(15 * time.Millisecond)
// 	}
// }

func worker(id int) {
	// workTime := time.Duration(rand.Intn(100)) * time.Millisecond
	// time.Sleep(workTime)
	fmt.Printf("Worker %d done\n", id)
	// ch <- fmt.Sprintf("Worker %d done in %v", id, workTime)
}

// func fast_worker(id int, ch chan string) {
// 	workTime := 1 * time.Millisecond
// 	time.Sleep(workTime)
// 	fmt.Println("Fast worker done waiting for receiver")
// 	ch <- fmt.Sprintf("fast worker done")
// }

func main() {

	results := make(chan string)

	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		go worker(i)
	}

	for i := 1; i <= numWorkers; i++ {
		// result := <-results
		// fmt.Println(result)
		fmt.Println("received")

	}

	// go fast_worker(1, results)
	time.Sleep(1 * time.Second)
	// result := <-results
	// fmt.Println(result)

	close(results)

	// Allow time for goroutines to complete
	select {}
}
