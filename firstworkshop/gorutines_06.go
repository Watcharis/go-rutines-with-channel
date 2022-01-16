package firstworkshop

import (
	"sync"
)

// Gorutines การเขียน ข้อมูลเข้า ที่ channel เดียวกัน
// การเขียนข้อมูลเข้า channel ด้วย gorutines หลายๆตัว โดยที่ gorutines นั้นๆ ส่งข้อมูลเข้า ที่ channel เดียวกัน
// โดยพื้นฐานเเล้ว การที่ goretines หลายๆตัวส่งข้อมูลเข้า ที่ channel เดียวกัน มันจะเกิด race-condition ตามมาอย่างแน่นอน

// จากการจะส่งค่า เข้า channel โดยคำนึงถึงเงื่อนไข ของ race-condition และ สามารถทำให้ gorutines สามารถส่งค่าเข้า channel ได้ครบโดยไม่เกิด panic

// จากตัวอย่าง นี่คือ การเขียน gorutines เข้า channel เดียวกัน
func MultipleGoRutines06() <-chan string {
	c := make(chan string, 3)

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	// สร้าง main gorutines เพื่อ ควบคุมการทำงานของ gorutines ที่จะส่งข้อมูลเข้า channel
	// ใน case นี้ main gorutines จะ ทำการ ปิด channel เมื่อ gorutines ตัวอื่นๆทำงานเสร็จ
	go func(wg *sync.WaitGroup, mu *sync.Mutex, c chan string) { // main gorutines
		wg.Add(1)
		go func(c chan string, mu *sync.Mutex) { //gorutines 01
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			i := 5
			for j := 0; j < i; j++ {
				k := "hello world"
				c <- k
			}
		}(c, mu)

		wg.Add(1)
		go func(c chan string, mu *sync.Mutex) { //gorutines 02
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			a := 5
			for b := 0; b < a; b++ {
				g := "strong hold"
				c <- g
			}
		}(c, mu)

		wg.Add(1)
		go func(c chan string, mu *sync.Mutex) { //gorutines 03
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			a := 5
			for b := 0; b < a; b++ {
				g := "flame guard"
				c <- g
			}
		}(c, mu)

		defer close(c)
		wg.Wait()
	}(wg, mu, c)

	return c
}
