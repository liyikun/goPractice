package entitles

import (
	"math/rand"
	"time"
	"fmt"
)

var (
	numberGoroutines = 4
	taskLoad = 10
)

func MoreChan()  {
	tasks := make(chan string,10)

	wg.Add(numberGoroutines)

	for gr:=1;gr<=numberGoroutines;gr ++ {
		go worker(tasks,gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d",post)
	}

	close(tasks)

	wg.Wait()

}

func init() {
	rand.Seed(time.Now().Unix())
}


func worker(tasks chan string,worker int) {

	defer wg.Done()

	for {

		task, ok := <-tasks

		if !ok {
			fmt.Printf("Worker: %d ;shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d :Started %s\n", worker, task)

		sleep := rand.Float64() * 10000
		fmt.Println("sleep:",sleep)
		fmt.Println("duration",time.Duration(sleep))
		time.Sleep(time.Duration(sleep)*time.Millisecond)

		fmt.Printf("Worker: %d :Is Finished %s \n", worker, task)
	}
}