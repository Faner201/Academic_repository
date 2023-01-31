package main

import (
	"fmt"
	"time"
)

type Addres struct {
	Street string
	City   string
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	DataBrith time.Time
	Addres    Addres
}

func NewUser(firstName, lastName, email string, dataBrith time.Time, addres Addres) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		DataBrith: dataBrith,
		Addres:    addres,
	}
}

func (u *User) Show() {
	fmt.Printf("%s - %s address: %+v\n", u.FirstName, u.LastName, u.Addres)
}

func (u *User) SetFirstName(name string) {
	u.FirstName = name
}

type Admin struct {
	IsAdmin bool
	User
}

func (a *Admin) Show() {
	fmt.Printf("admin: %s %s\n", a.FirstName, a.LastName)
}

func main() {
	addres := Addres{
		"Some city",
		"Some street",
	}

	user := NewUser(
		"Adria",
		"Lopatka",
		"lopatkaAd@gmail.com",
		time.Now(),
		addres,
	)

	admin := Admin{
		true,
		*user,
	}

	admin.Show()

	user.SetFirstName("Polka")
	user.Show()
}
