package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func func1() {
	fmt.Println("func1.....")
}
func func2() {
	fmt.Println("func2.....")
	wg.Done()
}

func func3() {
	fmt.Println("func3.....")
	wg.Done()
}

func sumchan(c chan int, s []int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

var wg sync.WaitGroup

func main() {
	go func1()
	time.Sleep(time.Second)
	gc := runtime.NumGoroutine()
	fmt.Println(gc)

	// synchronized goroutines
	fmt.Println("begin")
	wg.Add(2)

	go func2()
	go func3()

	wg.Wait()
	fmt.Println("end")

	s := []int{1, 5, 2, 8}
	c := make(chan int, 5) // initialize number of channels.
	go sumchan(c, s[:len(s)/2])
	fmt.Println(s[:len(s)/2])
	go sumchan(c, s[len(s)/2:])
	fmt.Println(s[len(s)/2:])
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) // 10 6 16
}
