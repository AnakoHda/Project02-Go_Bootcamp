package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	//k := flag.Uint("N", 0, "first")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Use 1 arguments")
		return
	}
	k, err1 := strconv.ParseUint(args[0], 10, 64)
	if err1 != nil {
		fmt.Println("Use uint arguments")
		return
	}

	doneSignal := make(chan struct{})

	go func(interval uint64) {
		for i := 1; ; i++ {
			select {
			case <-doneSignal:
				return
			default:
				t := uint64(i) * interval
				time.Sleep(time.Duration(interval) * time.Second)
				fmt.Printf("<Tick %d,%d>\n", i, t)
			}

		}
	}(k)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Termination")
	// если необходимо завершить задачу в горутине до конца
	//doneSignal <- struct{}{}

}
