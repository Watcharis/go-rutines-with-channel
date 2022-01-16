package firstworkshop

import (
	"fmt"
	"sync"
)

func MultipleGoRutines() int {
	var i int
	fmt.Printf("<<--------------- MultipleGoRutines ------------------>>\n")

	// Initialize a waitgroup variable
	wg := new(sync.WaitGroup) // var wg sync.WaitGroup
	m := new(sync.Mutex)      // var m sync.Mutex

	hundred := make(chan int)
	hundredTwo := make(chan int)

	set := func(n int) int {
		i += n
		return i
	}

	// `Add(1) signifies that there is 1 task that we need to wait for
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		// Calling `wg.Done` indicates that we are done with the task we are waiting fo
		defer wg.Done()

		// The `Lock` method of the mutex blocks if it is already locked
		// if not, then it blocks other calls until the `Unlock` method is called
		m.Lock()

		i = set(5)

		// Defer `Unlock` until this method returns
		m.Unlock()
	}(wg, m)

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		m.Lock()
		i = set(6)
		m.Unlock()
	}(wg, m)

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		m.Lock()
		i = set(7)
		m.Unlock()
	}(wg, m)

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		defer wg.Done()
		m.Lock()
		i = set(8)
		hundred <- i
		close(hundred)
		m.Unlock()
	}(wg, m)

	//	การโยนค่าจาก channel เข้า gorutines ควรทำการ pass ค่าโดยใช้ parameter
	//	เพื่อป้องกันไม่ให้ gorutines นั้นๆนำค่าจาก address เดิมมาใช้เผื่อค่ามีการเปลี่ยงเเปลง
	dataHundred := <-hundred
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex, i int) {
		defer wg.Done()
		fmt.Println("dataHundred ->", i)
	}(wg, m, dataHundred)

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex, hundred <-chan int) {
		defer wg.Done()

		// ******
		//  sync.Mutex
		// 		การใช้ Mutex lock การทำงานของ gorutines เพื่อจำกัดการเข้า ถึง resource ที่ถูกใช้ร่วมกัน
		//      โดย Mutex จะทำการ lock() gorutines นั้น ที่กำลังทำงานและเข้าถึง resource นั้นอยู่ จนกว่าจะเสร็จแล้ว unlock()
		// 		มัน ถึงจะยอมให้ gorutines ตัวถัดไปเริ่มทำงานและเข้าใช้ resource นั้น
		// ******

		// ตัวอย่าง ต่อไปนี้ ไม่จำเป็นต้อง ใช้ Mutex เพราะ ใช้ channel ในการ pass ค่า สู่ variable i
		// ทำให้ gorutines ไม่ต้องแย่งกันเพื่อ access ค่าให้ variable i

		// การ lock() ใช้เฉพาะตอนที่ มีการเข้า ถึง resouce เดียวกัน เช่น
		//   - การ access data เข้า variable i ของ gorutines
		//   - การเขียน ค่า เข้า channel
		stack := 0
		for data := range hundred {
			stack += data
		}

		i += stack

		wg.Add(1)
		go func(wg *sync.WaitGroup, m *sync.Mutex, i int) {
			defer wg.Done()
			m.Lock()
			t := i
			hundredTwo <- t
			m.Unlock()
		}(wg, m, i)

		wg.Add(1)
		go func(wg *sync.WaitGroup, m *sync.Mutex) {
			defer wg.Done()
			data := <-hundredTwo

			fmt.Println("data ->", data)
			fmt.Println("last i :", i)

			i += data
		}(wg, m)

	}(wg, m, hundred)

	// timeout := time.Tick(3 * time.Second)

	// `wg.Wait` blocks until `wg.Done` is called the same number of times
	// as the amount of tasks we have (in this case, 1 time)
	wg.Wait()
	return i
}
