package models

import "fmt"

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var UsersDB = make(map[uint]User)

func GetUsers() (users []User) {
	users = make([]User, 0, len(UsersDB))
	for _, u := range UsersDB {
		users = append(users, u)
	}
	return
}

func GetUserById(id uint) (user User) {
	user = UsersDB[id]
	return
}

func UpdateUserById(id uint, userToUpdate User) (user User, err error) {
	if user, ok := UsersDB[id]; ok {
		user = userToUpdate
		UsersDB[id] = user
		return user, nil
	}
	return User{}, fmt.Errorf("User with id %d not found", id)
}

func AddUser(user User) (status int) {
	UsersDB[uint(len(UsersDB)+1)] = user
	return 1
}

func DeleteUser(id uint) (status int) {
	if _, ok := UsersDB[id]; ok {
		delete(UsersDB, id)
		return 1
	}
	return 0
}