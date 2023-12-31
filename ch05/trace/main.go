package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ...lots of work...
	fmt.Println("Inside bigSlowOperation")
	time.Sleep(10 * time.Millisecond) // simulate slow operations by sleeping
}

func trace(msg string) func() {
	fmt.Println("Inside trace")
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()

	_ = double(4)
	fmt.Println(triple(5))
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
