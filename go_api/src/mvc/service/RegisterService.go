package service

import (
	. "common"
	. "core"
	. "data-base"
	. "mvc/dao"
	"github.com/go-xorm/xorm"
)

func CheckUserExist(user Sys_user) bool {
	userList :=  make([]Sys_user, 0)
	Engine.Where("username = ?", user.Username).Find(&userList)
	return len(userList) > 0
}

func RegisterUser(user Sys_user) string {
	res := Result{}
	user.Id = GetUUID()
	user.Status = 0

	Transaction(Engine, &res, func(session *xorm.Session) {
		_, err := session.Insert(&user)
		if err != nil {
			panic(ERROR)
		}
	})

	if res.Code == 0 {
		return user.Id
	}else{
		return ""
	}
}

func CheckUserActivated(userId string) bool {
	userList :=  make([]Sys_user, 0)
	Engine.Where("id = ? and status = 1", userId).Find(&userList)
	return len(userList) > 0
}

func ActivateAccount(userId string) bool {
	result, err := Engine.Exec("update sys_user set status = 1 where id = ? and status = 0", userId)
	if err != nil {
		return false
	} else {
		rows, _ := result.RowsAffected()

		if int64(0) == rows {
			return false
		}

		return true
	}
}
