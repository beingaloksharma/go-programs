// package main

// import "fmt"

// type User struct {
// 	Name     string
// 	Email    string
// 	Mobno    int64
// 	Password string
// }

// func (u *User) Set(data User) {
// 	u.Name = data.Name
// 	u.Email = data.Email
// 	u.Mobno = data.Mobno
// 	u.Password = data.Password
// }

// func (u *User) Get() {
// 	fmt.Printf("Name - %s \nEmail - %s \nMoblie Number - %d \nPassword - %s \n", u.Name, u.Email, u.Mobno, u.Password)
// }

// func main() {
// 	//Initilize User
// 	user := User{
// 		Name:     "Alok Kumar Sharma",
// 		Email:    "alok@email.com",
// 		Mobno:    1234567890,
// 		Password: "Pass@1234",
// 	}
// 	// Set
// 	user.Set(user)
// 	user.Get()
// }
