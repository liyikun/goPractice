package entitles

import (
	"sync"
	"sync/atomic"
	"runtime"
	"fmt"
	"time"
)

var (
	counter int64

	wg sync.WaitGroup

	wg2 sync.WaitGroup

	shutdown int64
)

func Listingmain()  {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()

	fmt.Println("final result is ",counter)

	wg2.Add(2)


	go doWork("A")
	go doWork("B")


	time.Sleep(1 * time.Millisecond)
	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown, 1)
	wg2.Wait()
}

func doWork(name string)  {

	defer wg2.Done()

	for {
		fmt.Println("work is",name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("println %s shutdown\n", name)
			break;
		}
	}
}

func intCounter(id int)  {

	defer wg.Done()

	for count:=0 ; count < 2; count ++ {
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}
}
