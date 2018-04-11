package controller

import (
	. "core"
	. "common"
	. "mvc/service"
	. "mvc/dao"
)


func init() {

	/**发送激活邮件*/
	Router.Register("/sendRegisterMail", func(requireHandle RequireHandle) (result Result) {
		user := Sys_user{}
		SetStructEntity(requireHandle, &user)

		//校验用户是否重复
		if CheckUserExist(user){
			return result.SetError(USER_HAS_EXIST)
		}

		//注册新用户
		userId := RegisterUser(user)

		if userId == ""{
			return result.SetError(ERROR)
		}

		//发送邮件
		body := "<a href='http://localhost:8888/activateAccount?id=" + userId + "'>点击链接激活账号</a>"
		err := SendToMail(user.Username, "激活账号", body, true)

		//邮件发送失败
		if err != nil{
			return result.SetError(ERROR)
		}

		return result.SetError(SUCCESS)
	})

	/**点击激活链接激活账号*/
	Router.Register("/activateAccount", func(requireHandle RequireHandle) (result Result) {

		userId := requireHandle.Query["id"][0]

		//校验用户是否已经激活
		if CheckUserActivated(userId){
			return result.SetError(USER_HAS_ACTIVATED)
		}

		if userId == "" || !ActivateAccount(userId){
			return result.SetError(ERROR)
		}

		return result.SetError(SUCCESS)
	})

}
