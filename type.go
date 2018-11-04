package main

import "fmt"

type user struct {
	name string
	email string
	ext int
	privileged bool
}

type admin struct {
	person user
	level string
}

func (u user) notify() {
	fmt.Println("Send Notify",u.name,"<"+u.email+">")
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func checktype() {
	//var bill user
	lisa := user{
		name: "lisa",
		email: "lisa@126.com",
		ext: 123,
		privileged: true,
	}
	mike := &user{
		"mike","mike@live.com",123,true,
	}
	lisa.notify()
	mike.notify()

	lisa.changeEmail("lisab@gmail.com")
	lisa.notify()

}
