package service

import (
	. "common"
	. "core"
	. "data-base"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id   string `xorm:"pk"`
	Name string
}

func FindAll() []User {
	everyone := make([]User, 0)
	err := Engine.Find(&everyone)

	if err == nil{

	}

	return everyone
}

func SaveUser(name string, id string) Result {
	user := User{id, name}

	res := Result{}.SetError(SUCCESS)

	Transaction(Engine, &res , func(session *xorm.Session){
		_, err := session.Insert(&user)
		if err != nil {
			panic(ERROR)
		}
	})

	return res
}