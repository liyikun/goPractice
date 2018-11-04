package entitles

import (
	"runtime"
	"sync"
	"fmt"
)

func Mostfuture() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			 for char := 'a'; char < 'a'+26; char++ {
				 fmt.Printf("%c ", char)
				 }
			 }
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Wait to Finish")

	wg.Wait()

	fmt.Println("finish")

}

var primewg sync.WaitGroup

func Primemain()  {

	runtime.GOMAXPROCS(runtime.NumCPU())
	defer fmt.Println(runtime.NumCPU())
	defer fmt.Println("22")
	primewg.Add(2)
	fmt.Println("start printprime")
	go printprime("A")
	go printprime("b")

	fmt.Println("Wait Finish")
	primewg.Wait()

	fmt.Println("Terminating Program")
}

func printprime(prefix string) {
	defer primewg.Done()

		for outer:=2 ; outer < 5000; outer++{
			for inter :=2 ; inter < outer; inter ++ {
				if outer%inter == 0 {
					goto find
				}
			}
			find:
			fmt.Printf("%s,%d\n",prefix,outer)
		}
	fmt.Println("Completed",prefix)
}
