package main

import (
	"fmt"
	"time"
)

//state for store some data highe level more than wait if data on channel full condition can end
func main() {
	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	for i := range 10{
		jobCh <- i + 1
		fmt.Println(i + 1)
	}
	close(jobCh)

	go double(jobCh, resultCh)
	// for range 2 {
	// 	go double(jobCh, resultCh)
	// }

	for i := range 10 {
		result := <-resultCh
		fmt.Println(result)
	}
}

func double(jobCh <-chan int, resultCh chan<- int) {
	for j := range jobCh {
		time.Sleep(1 * time.Second)
		resultCh <- j * 2
	}
}
