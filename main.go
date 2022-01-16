package main

import (
	"fmt"
	"sync"
	"time"
	"watcharis/go-routines-chanel/firstworkshop"
)

func main() {
	start := time.Now()
	wg := new(sync.WaitGroup)
	strChan := make(chan string)

	// channel goroutines สามารถเเบ่งออกได้เป็น 2 ประเภท คือ
	// - buffers ---> intChan := make(chan int)
	// - unbuffers ---> testChan := make(chan int, 10)
	// buffers คือ channel ที่ไม่มีการกำหนดขนาดในการส่งข้อมูล ทำให้ สามารถ ส่ง ค่า เข้า  channel ได้ทีละครั้ง
	// unbuffers คือ channel ที่มีการกำหนดขนาดในการส่งข้อมูล ทำให้ สามารถ ส่ง ค่า เข้า  channel ได้ทีละหลายค่า
	// ลูกศรเป็นตัวกำหนด พฤติกรรม การทำงานของ channel
	//  test := make(chan int)
	//  test <- int เป็นการเขียนข้อมูล ลง channel || send only
	//  <-test int	เป็นการอ่านข้อมูล จาก channel || recieve only

	intChan := make(chan int)      //unbuffers
	testChan := make(chan int, 10) //buffers
	testTwoChan := make(chan int)

	// deadlock will be occurred at this line
	wg.Add(1)
	go func(strChan <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Printf("received str: %s\n", <-strChan)
	}(strChan, wg)

	strChan <- "hello, world!"

	wg.Add(1)
	go func(intChan <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Printf("received int: %d\n", <-intChan)
	}(intChan, wg)

	intChan <- 20

	//ตัวอย่าง channel unbuffers
	wg.Add(1)
	go func(testChan <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for t := range testChan {
			fmt.Printf("channel unbuffers received int: %d\n", t)
		}
		//
		// fmt.Printf("received int: %d\n", <-testChan)
		// fmt.Printf("received int: %d\n", <-testChan)
		// fmt.Printf("received int: %d\n", <-testChan)

		// data := <-testChan
		// checkArray := make(map[int]int, 0)
		// for i := 0; i < data; i++ {
		// 	checkArray[i+1] = i + 1
		// }
		// fmt.Println("checkArray ->", checkArray)

	}(testChan, wg)

	// ทำการส่งค่าเข้า testChan หลายๆค่า
	// testChan <- 9
	// testChan <- 1
	// testChan <- 2
	// testChan <- 3
	// testChan <- 4
	// testChan <- 5
	// testChan <- 6
	// testChan <- 7
	// testChan <- 8
	// testChan <- 0
	tc := []int{1, 2, 4, 5, 6, 7, 8, 9}
	for _, d := range tc {
		testChan <- d
	}

	go firstworkshop.Publisher(testTwoChan)

	go firstworkshop.Subscriber(testTwoChan, wg)

	go firstworkshop.ChanelOne(intChan, wg)
	// channel จะใช้การ for loop ในการนำข้อมุลมาใช้ หากมี การส่งค่าเข้า channel
	for inchanel := range intChan {
		fmt.Println("inchanel ->", inchanel)
	}

	in := firstworkshop.Gen(2, 3, 4)
	fmt.Println("in ->", in)

	c1 := firstworkshop.Sq(in)

	c2 := firstworkshop.Sq(in)

	// channel จะใช้การ for loop ในการนำข้อมุลมาใช้ หากมี การส่งค่าเข้า channel
	// for j := range c1 {
	// 	fmt.Println("c1 ->", j)
	// }
	// for k := range c2 {
	// 	fmt.Println("c2 ->", k)
	// }

	responseFirstClassFunctions := func() []int {
		data := []int{}
		for n := range firstworkshop.Merge(c1, c2) {
			fmt.Println("Merge n --->", n)
			data = append(data, n)
		}
		return data
	}()
	fmt.Println("responseFirstClassFunctions ->>>", responseFirstClassFunctions)

	//higher order function
	fff, sss := firstworkshop.Sphere(11)
	t1 := fff(1.24)
	t2 := sss("hello world")

	fmt.Println("t1 ->", t1)
	fmt.Println("t2 ->", t2)

	m := firstworkshop.MultipleCalculate()
	a := m(2)
	b := a(2)
	fmt.Println("b ->", b)

	// responseMessage := make(chan string)
	// wg.Add(1)
	// go func(method string, url string, responseMessage <-chan string, wg *sync.WaitGroup) {
	// 	requestFlask, _ := firstworkshop.Request(method, url, nil)
	// 	defer wg.Done()

	// 	for flaskResponse := range requestFlask {
	// 		fmt.Println("flaskResponse ->", flaskResponse.Body.Message+" "+<-responseMessage)
	// 	}

	// }("GET", "http://localhost:5005", responseMessage, wg)

	// responseMessage <- "Omicron"

	// wg.Add(1)
	// go func(method string, url string, responseMessage chan<- string, wg *sync.WaitGroup) {
	// 	requestFlask, _ := firstworkshop.Request(method, url, nil)
	// 	defer wg.Done()

	// 	message := make(chan string)
	// 	go func() {
	// 		for flaskResponse := range requestFlask {
	// 			message <- flaskResponse.Body.Message
	// 			close(message)
	// 		}
	// 	}()

	// 	for rawMessage := range message {
	// 		responseMessage <- rawMessage
	// 		close(responseMessage)
	// 	}

	// }("GET", "http://localhost:5005", responseMessage, wg)

	// for flaskResponse := range responseMessage {
	// 	fmt.Println("flaskResponse ->", flaskResponse)
	// }

	wg.Wait()

	fmt.Printf("time use: %+v\n", time.Since(start).Seconds())

	testMultipleGoRutines := firstworkshop.MultipleGoRutines()
	fmt.Println("testMultipleGoRutines ->", testMultipleGoRutines)

	// testRaceConditions := firstworkshop.RaceCondition01()
	// fmt.Println("testRaceConditions ->", testRaceConditions)

	// testCheckSpeedDelareMap := firstworkshop.CheckSpeedDelareMap()
	// fmt.Println("testCheckSpeedDelareMap ->", testCheckSpeedDelareMap)

	testCheckSwitchCase := firstworkshop.CheckSwitchCase(79)
	fmt.Println("testCheckSwitchCase ->", testCheckSwitchCase)

	// getBirthDate := firstworkshop.FindBirthDay()
	// fmt.Println("getBirthDate ->", getBirthDate)

	// firstworkshop.MultipleGoRutines04()
	// firstworkshop.SyncMapTutorial()

	// recieveDataFromChannel := firstworkshop.MultipleGoRutines06()
	// for text := range recieveDataFromChannel {
	// 	fmt.Println("text ->", text)
	// }

	s, t, bo := firstworkshop.MultipleGoRutines07()
	fmt.Println("s ->", s)
	fmt.Println("t ->", t)
	fmt.Println("bo ->", bo)

	s0 := firstworkshop.MultipleGoRutines08()
	fmt.Println("s0 ->", s0)

	// closure function สามารถ สเเตก ค่า ที่ ถุกเก็บไว้ใน memory ได้
	// จาก func one.XX() มีการประกาศ i = 0 เเล้ว ทำการ return สเเตกที่ i += 1
	// ทุกๆครั้ง ที่เรียก add() จะทำการ บวกค่า 1
	// 		add := firstworkshop.XX()
	// 		fmt.Println(add())
	// 		fmt.Println(add())
	// 		fmt.Println(add())
	// 		return
	fmt.Println("say say say")

	fmt.Println("world world world")
}
