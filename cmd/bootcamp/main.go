package main

import (
	"context"
	"fmt"
)

func main() {

	numbers := make(chan int)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "stream", numbers)

	go func(context.Context) {
		stream := ctx.Value("stream").(chan int)
		for num := range stream {
			fmt.Println("Got:", num)
		}
	}(ctx)

	go func(context.Context) {
		stream := ctx.Value("stream").(chan int)
		for i := 0; i <= 5; i++ {
			stream <- i
		}
		close(stream)
	}(ctx)

	_, _ = fmt.Scanln()
}
