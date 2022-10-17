package firstworkshop

import (
	"fmt"
	"sync"
	"time"
)

func Test09() error {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	a := make(chan int, 1)
	b := make(chan int, 1)
	count := 0
	// fmt.Println("a")
	// a <- 1
	// fmt.Println("b")
	// b <- 2

	wg.Add(1)
	go func() {

		defer wg.Done()

		fmt.Println("line 32")
		a <- 1
		fmt.Println("a")
	}()

	fmt.Println("---- before time sleep ----")
	time.Sleep(3 * time.Second)
	fmt.Println("---- after time sleep ----")

	o := <-a
	fmt.Println("---- before close ----", o)
	close(a)

	wg.Add(1)
	go func(o int) {

		defer func() {
			fmt.Println("-- defer gorutine 2 --")
			wg.Done()
		}()

		fmt.Println("before 222")
		b <- 2
		mu.Lock()
		count = o
		mu.Unlock()
		fmt.Println("o ->", o)
	}(o)

	fmt.Println("before wait")

	wg.Wait()
	close(b)

	count += <-b

	fmt.Println("before final")
	fmt.Println("count ->", count)
	return nil
}
