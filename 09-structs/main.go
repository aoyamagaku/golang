package main

import "fmt"

type User struct {
	Name        string  //""
	Age         int     // 0
	PhoneNumber string  // ""
	IsClose     bool    // false
	Rating      float64 // 1.0 -- 100.0
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isclose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}
	if age <= 0 || age >= 150 {
		return User{}
	}
	if phoneNumber == "" {
		return User{}
	}
	if rating < 1.0 || rating > 100.0 {
		return User{}
	}
	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isclose,
		Rating:      rating,
	}
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhoneNumber(newNumber string) {
	if newNumber != "" {
		u.PhoneNumber = newNumber
	}
}

func (u *User) CloseAccount() {
	u.IsClose = true
}

func (u *User) OpenAccount() {
	u.IsClose = false
}

func (u *User) RatingUP(rating float64) {
	if u.Rating+rating <= 100.0 {
		u.Rating += rating
	}

}

func (u *User) RatingDOWN(rating float64) {
	if u.Rating-rating >= 1.0 {
		u.Rating -= rating
	}
}
func main() {
	user := NewUser(
		"Danila",
		32,
		"+7 (999) 33 22 11",
		true,
		64,
	)

	fmt.Println("User:", user)
}
