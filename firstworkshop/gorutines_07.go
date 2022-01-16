package firstworkshop

import (
	"fmt"
	"sync"
)

func MultipleGoRutines07() ([]string, []string, bool) {
	fmt.Println("---------------------- MultipleGoRutines07 ----------------------")
	round := 4

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	t := []string{}
	f := []string{}
	cc := make(chan bool)

	mock := 1
	isBool := false

	// wg2 := new(sync.WaitGroup)

	for i := 0; i < round; i++ {

		wg.Add(1)
		go func() {
			mu.Lock()
			defer func() {
				wg.Done()
				mu.Unlock()
			}()

			if mock > 2 {
				fmt.Println("in t")
				t = append(t, "hello")
				// isBool = true
				cc <- true
			} else {
				fmt.Println("in f")
				f = append(f, "world")
				mock += 1
				// isBool = false
				cc <- false
			}
		}()
	}

	cc2 := make(chan bool)
	// wg2.Add(1)
	go func(isBool bool) {
		// wg2.Done()
		b := isBool
		for checkBool := range cc {
			fmt.Println("checkBool ->", checkBool)
			b = checkBool
		}

		fmt.Println("isBool ->", b)
		cc2 <- isBool
		// if !isBool {
		// 	cc2 <- false
		// 	close(cc2)
		// } else {
		// 	cc2 <- true
		// 	close(cc2)
		// }
	}(isBool)

	wg.Wait()
	close(cc)
	// wg2.Wait()

	resultBoolean := <-cc2

	return t, f, resultBoolean
}
