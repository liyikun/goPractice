package main

import "fmt"

func array()  {
	fmt.Println("hello")



	var array1 [5]int

	array2 := [5]int{10,20,30,40,50}

	array2[2] = 50

	array3 := [...]int{10,20,40}

	array4 := [5]int{1: 20, 4: 30}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

	array5 := [...]*int{1: new(int),2: new(int)}

	*array5[1] = 4
	*array5[2] = 5
	fmt.Println(*array5[1])

	var array7 [3]*string
	array6 := [...]*string{new(string),new(string),new(string)}
	*array6[0] = "test1"
	*array6[1] = "test2"
	*array6[2] = "test3"
	array7 = array6
	fmt.Println(*array7[2])
	*array6[2] = "test4"
	fmt.Println(*array7[2])


	var array8 [4][2]int

	array9 := [4][2]int{{1,2},{3,4},{5,6},{7,8}}

	fmt.Println(array8)
	fmt.Println(array9)

	array9[2][0] = 1
	fmt.Println(array9)

	var array10 [2]int = array9[1]
	var value int = array9[1][0]
	fmt.Println(array10)
	fmt.Println(value)




	}
