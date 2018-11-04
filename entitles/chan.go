package entitles

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

var wgchan sync.WaitGroup

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func TestChanMain() {
	court := make(chan int)

	wgchan.Add(2)

	go player("Fish",court)
	go player("Venny",court)

	court <-1

	wgchan.Wait()
	}

func player(name string,court chan int )  {

	defer wgchan.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Win\n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Players %s Missed\n", name)

			close(court)
			return
		}

		ball++

		fmt.Printf("Player %s Hit %d\n",name,ball)

		court <- ball
	}

}



