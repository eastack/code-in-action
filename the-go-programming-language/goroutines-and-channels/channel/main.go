package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	defer close(out)
	for x := 0; x < 100; x++ {
		out <- x
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Drained")
}

func squarer(out chan<- int, in <-chan int) {
	defer close(out)
	for x := range in {
		out <- x * x
	}
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
