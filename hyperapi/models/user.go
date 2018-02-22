package models

import (
	"errors"
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
	Id       string
}

func AddUser(u User) string {
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
    _ = fabric_query_user(uid)
	return UserList["user_11111"], err
}

func GetAllUsers() map[string]*User {
    _ = fabric_query_users()
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Id != "" {
			u.Id = uu.Id
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
