package model

type Role struct{
	Id string
	Name string
}

type UserObject  struct {
	Name string
	Marks int
	Role Role
}

type Student interface{
	speak_name()
}

