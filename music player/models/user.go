package models

var ArrayUser[]User

type User struct{
	ID int 
	Firstname string
	Lastname string
	SavedMusics []Music
}