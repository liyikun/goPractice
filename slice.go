package main

import "fmt"

func slice()  {

	slice := make([]string,5)
	slice1 := make([]string,3,5)

	fmt.Println(slice)
	fmt.Println(slice1)

	slice2 := []string{99:""}

	slice3 := []int{1,2,3}

	fmt.Println(slice2)
	fmt.Println(slice3)


	slice4 := []int{1,2,3,4,5}
	slice5 := slice4[1:3]

	slice5[1] = 10

	fmt.Println(slice4)
	fmt.Println(slice5)

	slice5 = append(slice5,60,70)
	fmt.Println(slice4)
	fmt.Println(slice5)

	source := []int{10,20,30,40,50}
	//slice6 := source[1:4:5]

	//for _,value := range slice6 {
	//	fmt.Println(value)
	//}
	//
	//for index:=1;index < len(slice6);index ++ {
	//	fmt.Println(slice6[index])
	//}

	foo(source)
	fmt.Println(source)
}

func foo(slice []int)  {
	slice[1] = 1000
}
