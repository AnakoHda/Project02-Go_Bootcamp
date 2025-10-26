package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type MyData struct {
	index int
	time  int
}

func main() {
	//n := flag.Int("N", 0, "first")
	//m := flag.Int("M", 0, "second")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Use 2 arguments")
		return
	}

	n, err1 := strconv.Atoi(args[0])
	m, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Use int arguments")
		return
	}

	data := sortTask(n, m)
	for i := 1; i <= n; i++ {
		fmt.Printf("<%d,%d>\n", data[n-i].index, data[n-i].time)
	}
	//str := mutexTask(n, m)
	//for i := 1; i <= n; i++ {
	//	fmt.Println(str[n-i])
	//}
}

func mutexTask(n int, m int) []string {
	var wg sync.WaitGroup
	wg.Add(n)
	str := make([]string, 0, n)
	var mu sync.Mutex
	for i := 0; i < n; i++ {
		go func(index int, s *[]string, wg *sync.WaitGroup, mu *sync.Mutex) {
			waitTime := rand.Intn(m + 1)
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
			mu.Lock()
			*s = append(*s, fmt.Sprintf("<%d,%v>", index, waitTime))
			mu.Unlock()
			wg.Done()
		}(i, &str, &wg, &mu)
	}
	wg.Wait()
	return str
}

func sortTask(n int, m int) []MyData {
	var wg sync.WaitGroup
	wg.Add(n)

	str := make([]MyData, n)

	for i := 0; i < n; i++ {
		go func(index int, s []MyData, wg *sync.WaitGroup) {
			waitTime := rand.Intn(m + 1)
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
			s[index] = MyData{
				index: index,
				time:  waitTime,
			}
			wg.Done()
		}(i, str, &wg)
	}
	wg.Wait()

	sort.Slice(str, func(i, j int) bool {
		return str[i].time < str[j].time
	})
	return str
}
