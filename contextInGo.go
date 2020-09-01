package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------------")

	ctx1, _ := context.WithCancel(ctx)
	fmt.Println("context:\t", ctx1)
	fmt.Println("context err:\t", ctx1.Err())
	fmt.Printf("context type:\t%T\n", ctx1)
	fmt.Println("--------------")

	ctx2, cancel := context.WithCancel(ctx)
	fmt.Println("context:\t", ctx2)
	fmt.Println("context err:\t", ctx2.Err())
	fmt.Printf("context type:\t%T\n", ctx2)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("--------------")

	cancel()
	fmt.Println("context:\t", ctx2)
	fmt.Println("context err:\t", ctx2.Err())
	fmt.Printf("context type:\t%T\n", ctx2)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("--------------")

	// Now let's see how we can use this in an actual code
	ctx, cancel = context.WithCancel(context.Background())
	fmt.Println("Error check 1:", ctx.Err())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2:", ctx.Err())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("About to cancel context")
	cancel()
	fmt.Println("Context is cancelled")

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 3:", ctx.Err())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
