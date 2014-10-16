package main

import "fmt"

func fibonacci() func() int {
	pre := 0
	fib := 1
	return func() int {
		tmp := fib
		fib = pre + fib
		pre = tmp
		return fib
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
