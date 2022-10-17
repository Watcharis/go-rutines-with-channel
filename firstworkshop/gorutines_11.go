package firstworkshop

import (
	"fmt"
	"sync"
)

func Test11() {

	wg := new(sync.WaitGroup)
	ch := make(chan int, 2)

	wg.Add(1)
	go func() {

		defer func() {
			fmt.Println("----- gorutine 1 -----")
			wg.Done()
		}()

		n := 10
		for i := 0; i < n; i++ {
			fmt.Println("a ->", i)
			ch <- i
		}
	}()

	wg.Add(1)
	go func() {

		defer func() {
			fmt.Println("----- gorutine 2 -----")
			wg.Done()
		}()

		n := 10
		for i := 0; i < n; i++ {
			fmt.Println("b ->", i)
			ch <- i
		}
	}()

	wg.Add(1)
	go func() {

		defer func() {
			fmt.Println("----- channel done -----")
			wg.Done()
		}()

		for value := range ch {
			fmt.Println("value ->", value)
		}
	}()

	wg.Wait()

	// select {
	// case <-ctx.Done():
	// 	fmt.Println("------ context done ------")
	// 	return
	// }
}
