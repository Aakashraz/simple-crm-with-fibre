package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	fmt.Println("assigning request-id ...")
	return context.WithValue(ctx, "request-id", "1234456789")
}

func doSomething(ctx context.Context) {
	requestId := ctx.Value("request-id")
	fmt.Println("assigned request-id: ", requestId)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("CONTEXT EXAMPLE")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fmt.Println("calling error function before Done():  ", ctx.Err())

	//ctx := context.Background()	// this initialization is done in above line
	ctx = enrichContext(ctx)
	go doSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
		fmt.Println("calling error function:  ", ctx.Err())
	}
	time.Sleep(2 * time.Second)
}
