package main

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	userArray := []User{
		User{
			Name:    "Bob",
			Rating:  5.5,
			Premium: true,
		},
		User{
			Name:    "Alice",
			Rating:  4.1,
			Premium: false,
		},
		User{
			Name:    "Max",
			Rating:  1.5,
			Premium: true,
		},
	}

	userArray = append(userArray, User{
		Name:    "Max",
		Rating:  7.9,
		Premium: true,
	})


	//fmt.Println("Before")
	//
	//for i := range userArray {
	//	fmt.Println(userArray[i])
	//}
	//
	//for i := range userArray {
	//	if userArray[i].Premium {
	//		userArray[i].Rating += 1.0
	//	}
	//}
	//
	//fmt.Println("After")
	//
	//for i := range userArray {
	//	fmt.Println(userArray[i])
	//}
}
