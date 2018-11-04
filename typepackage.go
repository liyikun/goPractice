package main

import (
	"./entitles"
	"fmt"
)

func typepackage()  {
	a := entitles.Admin{
		Rights: 10,
	}

	a.Name = "lyk"
	a.Email = "ee@ev.com"

	fmt.Printf("user is %v\n",a)

}
