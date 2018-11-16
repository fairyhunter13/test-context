package main

import (
	"context"
	"log"
	"time"
)

const (
	timeOut = 999
)

func main() {
	// startTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Millisecond)
	defer cancel()
	done := make(chan int)
	go workLong(ctx, cancel, done)
	select {
	case <-ctx.Done():
		time, _ := ctx.Deadline()
		if err := ctx.Err(); err == context.DeadlineExceeded {
			log.Printf("Deadline already reached! Time Deadline: %v", time)
		}
		<-done
		close(done)
	case <-done:
		log.Printf("The work is done! Deadline is not reached!")
	}
	time.Sleep(3000 * time.Millisecond)
}

func workLong(ctx context.Context, cancel context.CancelFunc, done chan<- int) {
	defer func() {
		log.Println("I have escaped with context, or maybe not!")
	}()
	//Reimagine the function to do heavy work.
	time.Sleep(1000 * time.Millisecond)
	log.Println("Work long succeed!")

	//Maybe in here it is not executed.
	time.Sleep(1000 * time.Millisecond)
	log.Println("Shouldn't be happening!")
	done <- 1
	return
}
