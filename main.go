package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	timeOut = 999
)

func main() {
	// mapNotConcurrent()
	// loopCounterCase()
	sharedVariable()
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
	for k, v := range m {
		fmt.Printf("The key %s and value: %s", k, v)
	}
}

func loopCounterCase() {
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

func sharedVariable() {
	errorChannel := ParallelWrite([]byte("Test"))
	for err := range errorChannel {
		if err != nil {
			fmt.Printf("There is an error! Error: %s", err)
		}
	}
}

// ParallelWrite writes data to file1 and file2, returns the errors.
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	defer close(res)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
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
