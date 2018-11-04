package entitles

import (
	"fmt"
	"time"
)

func  TestRun()  {


	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()

}

func Runner(baton chan int)  {
	var newRunner int

	runner := <- baton

	fmt.Printf("Runner %d Running with Baton\n", runner)

	if runner !=4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To Line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Println("Runner %d is Finish\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d",runner,newRunner)

	baton <- newRunner
}
