package main

import "fmt"

type notifier interface {
	notify()
}

type users struct {
	name string
	email string
}

func (u *users) notify() {
	fmt.Printf("Send user email to %s<%s>\n",
		u.name,u.email)
}

type admins struct {
	name string
	email string
}

func (u *admins) notify() {
	fmt.Printf("Send admin email to %s<%s>\n",
		u.name,u.email)
}
//
//func (u *boos2) notify() {
//	fmt.Printf("Send admin email to %s<%s>\n",
//		u.name,u.email)
//}


func (u *admins) ChangeName(name string) {
	u.name = name
}


type boss struct {
	admins
	level string
}

type boos2 struct {
	level string
	info admins
}

func SendNotify(n notifier) {
	n.notify()
}


func PolyMain() {
	bill := users{
		"bill",
		"bill@123.com",
	}
	liyikun := admins{
		"liyikun",
		"lyk@163.com",
	}

	SendNotify(&bill)
	SendNotify(&liyikun)

	vc := boss{
		admins: admins{name:"ss",email:"ss@qq.com"},
		level: "lv",
	}

	SendNotify(&vc)
	//ss := boos2{
	//	level:"lv2",
	//	info: admins{name:"ss",email:"ss@qq.com"},
	//}
	//
	//ss.info.ChangeName("ss2")
	////vc.info.ChangeName("vc2")
	//
	//vc.notify()


}
