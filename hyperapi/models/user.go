package models

import (
	"errors"
	"fmt"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111"}
	UserList["user_11111"] = &u
}

type User struct {
	Id string
}

func AddUser(u User) string {
	return fabric_add_user(u.Id)
}

func GetUser(uid string) (u *User, err error) {
	//_ = fabric_query_user(uid)
	fmt.Println("method not implemented")
	return nil, errors.New("Not implemented")
}

func GetCurrentUser() string {
	fmt.Println("...[GetCurrentUser]...")
	return orgName
}

func GetAllUsers() []string {
	return fabric_query_users()
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	fmt.Println("method not implemented")
	//if u, ok := UserList[uid]; ok {
	//	if uu.Id != "" {
	//		u.Id = uu.Id
	//	}
	//	return u, nil
	//}
	return nil, errors.New("Not implemented")
}

func DeleteUser(uid string) {
	fmt.Println("method not implemented")
	//delete(UserList, uid)
}
