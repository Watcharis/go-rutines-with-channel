package firstworkshop

import (
	"fmt"
	"sync"
	"time"
)

//จากตัวอย่างข้างต้น จะบอกถึง การทำงานของ channel กับ sync.WaitGroup

func MultipleGoRutines08() string {
	jobs := make(chan int, 1)

	// done := make(chan bool)

	wg2 := sync.WaitGroup{}

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range jobs {
			fmt.Printf("v ->%+v\n", v)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(5555)
		// done <- true
		// close(done)
	}()

	wg := sync.WaitGroup{}
	for j := 1; j <= 3; j++ {
		wg.Add(1)
		go func(src int) {
			fmt.Printf("src ->%+v\n", src)
			jobs <- src
			wg.Done()
		}(j)
	}
	wg.Wait()
	close(jobs)

	wg2.Wait()

	// <-done

	return "ok"
}
