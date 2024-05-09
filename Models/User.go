package models

import "fmt"

type User struct {
	Id       int   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var UsersDB = make(map[int]User)

func GetUsers() (users []User) {
	users = make([]User, 0, len(UsersDB))
	for _, u := range UsersDB {
		users = append(users, u)
	}
	return
}

func GetUserById(id int) (user User) {
	user = UsersDB[id]
	return
}

func UpdateUserById(id int, userToUpdate User) (user User, err error) {
	if user, ok := UsersDB[id]; ok {
		user = userToUpdate
		UsersDB[id] = user
		return user, nil
	}
	return User{}, fmt.Errorf("User with id %d not found", id)
}

func AddUser(user User) (status int) {
	UsersDB[int(len(UsersDB)+1)] = user
	return 1
}

func DeleteUser(id int) (status int) {
	if _, ok := UsersDB[id]; ok {
		delete(UsersDB, id)
		return 1
	}
	return 0
}