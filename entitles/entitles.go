package entitles


type user struct {
	Name string
	Email string
}

type Admin struct {
	user
	Rights int
}