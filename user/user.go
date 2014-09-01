package user

import (
    "code.google.com/p/go-uuid/uuid"
)

var userDataByName = make(map[string]User)
var userDataByID = make(map[string]User)

type User struct {
    ID string `json:"id"`
    Name string `json:"name"`
}

func GenUserID() string {
    return uuid.New()
}

func FindUser(userName string) User {
    return userDataByName[userName]
}

func (user *User) Signup(name string) {
    user.ID = GenUserID()
    user.Name = name
    user.Save()
}

func (user *User) Save() {
    userDataByName[user.Name] = *user
    userDataByID[user.ID] = *user
}

func (user *User) ToMap() map[string]string {
    return map[string]string {"id": user.ID, "name": user.Name}
}
