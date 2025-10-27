package main

import (
	"flag"
	"fmt"
	"strconv"
)

func generator(from, to int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := from; i <= to; i++ {
			//time.Sleep(time.Duration(i) * time.Second)
			out <- i
		}
	}()
	return out
}

func square(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for u := range input {
			//time.Sleep(time.Duration(v) * time.Second)
			out <- (u * u)
		}
	}()
	return out
}
func main() {
	//k := flag.Int("N", 0, "first")
	//n := flag.Int("M", 0, "second")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Use 2 arguments")
		return
	}
	k, err1 := strconv.Atoi(args[0])
	n, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Use int arguments")
		return
	}

	for i := range square(generator(k, n)) {
		fmt.Println(i)
	}
}
