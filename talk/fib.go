package main

import "fmt"

func fib(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		c <- a
	}
	close(c)
}

func main() {
	c := make(chan int)

	go fib(10, c)

	for v := range c {
		fmt.Println(v)
	}
}
