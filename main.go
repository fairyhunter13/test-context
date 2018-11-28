package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	timeOut = 999
)

var (
	globalCounter = 0
)

func main() {
	// mapNotConcurrent()
	// loopCounterCase()
	// counterNotWrightCase()
}

func mapNotConcurrent() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a"
		c <- true
	}()

	// time.Sleep(1 * time.Nanosecond)
	m["1"] = "c"
	<-c
	fmt.Println("==========This is the nonconcurrent map section!==========")
	for k, v := range m {
		fmt.Printf("The key %s and value: %s\n", k, v)
	}
}

func loopCounterCase() {
	fmt.Println("==========This is the loop counter section!==========")
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
	}
	wg.Wait()
}

func counterNotWrightCase() {
	fmt.Println("==========This is the loop counter section!==========")
	done := make(chan bool)
	go func() {
		for i := 1; i <= 1000; i++ {
			go func() {
				globalCounter++
			}()
		}
		done <- true
	}()
	<-done
	fmt.Printf("The global counter is %d\n", globalCounter)

}

/* This section is only for context. */

//TestContext
func contextWorking() {
	// startTime := time.Now()
	// ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Millisecond)
	// defer cancel()
	// done := make(chan int)
	// go workLong(ctx, cancel, done)
	// select {
	// case <-ctx.Done():
	// 	time, _ := ctx.Deadline()
	// 	if err := ctx.Err(); err == context.DeadlineExceeded {
	// 		log.Printf("Deadline already reached! Time Deadline: %v", time)
	// 	}
	// 	<-done
	// 	close(done)
	// case <-done:
	// 	log.Printf("The work is done! Deadline is not reached!")
	// }
	// time.Sleep(3000 * time.Millisecond)
}

func workLong(ctx context.Context, cancel context.CancelFunc, done chan<- int) {
	defer func() {
		log.Println("I have escaped with context, or maybe not!")
	}()
	//Reimagine the function to do heavy work.
	time.Sleep(1000 * time.Millisecond)
	log.Println("Work long succeed!")

	//Maybe in here it is not executed.
	log.Println("Shouldn't be happening!")
	done <- 1
	return
}
