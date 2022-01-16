package firstworkshop

import (
	"fmt"
	"sync"
	"time"
)

func MultipleGoRutines04() {

	var m sync.Mutex

	waitChan := make(chan bool, 3)

	go func() {
		var i int = 1
		for true {
			fmt.Printf("%d\n", i)
			i += 3
			time.Sleep(time.Duration(1) * time.Second)
			m.Lock()
			waitChan <- true
			m.Unlock()
		}
	}()

	go func() {
		var i int = 2

		for true {
			fmt.Printf("%d\n", i)
			i += 3
			time.Sleep(time.Duration(1) * time.Second)
			m.Lock()
			waitChan <- true
			m.Unlock()
		}
	}()

	go func() {
		var i int = 3

		for true {
			fmt.Printf("%d\n", i)
			i += 3
			time.Sleep(time.Duration(1) * time.Second)
			m.Lock()
			waitChan <- true
			m.Unlock()
		}
	}()

	for b := range waitChan {
		fmt.Printf("channel return : %+v\n", b)
	}
}
