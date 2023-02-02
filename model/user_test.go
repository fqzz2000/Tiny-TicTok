package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestUserDAO_QueryUserByID(t *testing.T) {
	Init_DB(true)
	var dbs []UserDB
	DB.Find(&dbs)
	fmt.Printf("%+v\n", dbs)
	us := NewUserDAO().QueryUserById(2)
	fmt.Printf("%+v\n", us.UserName)
	us = NewUserDAO().QueryUserById(1)
	fmt.Printf("%+v\n", us.UserName)
}

func TestUserDAO_AddNewUser(t *testing.T) {
	Init_DB(true)
	usr := UserDB{
		UserName: "testusr",
		UserPswd: "abcd",
	}
	var usrs []UserDB
	DB.Find(&usrs)
	oldLen := len(usrs)
	NewUserDAO().AddNewUser(&usr)
	DB.Find(&usrs)
	newLen := len(usrs)
	assert.Equal(t, newLen - 1, oldLen)
	DB.Where("user_name = ?", "testusr").Delete(&UserDB{})
}

func TestUserDAO_QueryNameExists(t *testing.T) {
	Init_DB(true)
	assert.Equal(t, false, NewUserDAO().QueryNameExists("hello"))
	assert.Equal(t, true, NewUserDAO().QueryNameExists("Kiyo Wu"))
}
