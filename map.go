package main

import "fmt"

func testmap()  {
	//dict := make(map[string]int)
	//
	//dict1 := map[string]string{"ff":"aa","44":"33","ss1":"313"}

	dict2 := map[int][]int{}
	dict2[1] = make([]int,10)
	fmt.Println(dict2)

	color := map[string]string{}

	color["RED"] = "aa"
	fmt.Println(color)

	value,exists := color["RED"]
	_,exists2 := color["FF"]


	fmt.Println(value,exists,exists2)

	colors := map[string]string{
		"RED":"#111",
		"BLUE":"#222",
		"BLACK":"#333",
	}
	for key,value := range colors{
		fmt.Println(key,value)
	}
	delete(colors,"RED")
	for key,value := range colors{
		fmt.Println(key,value)
	}
	}
