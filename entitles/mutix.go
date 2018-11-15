package entitles

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter3 int
	wg3 sync.WaitGroup
	mutux sync.Mutex
)

func MutixMain(){

	wg3.Add(2)

	go intCounter3(1)
	go intCounter3(2)

	wg3.Wait()

	fmt.Println(counter3)

}


func intCounter3 (id int) {

	defer  wg3.Done()

	for count := 0;count < 2;count ++ {
		mutux.Lock()
		{
			value := counter3
			runtime.Gosched()
			value++
			counter3 = value

		}
		mutux.Unlock()
	}
}