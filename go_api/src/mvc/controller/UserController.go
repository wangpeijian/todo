package controller

import (
	. "core"
	. "mvc/service"
	"fmt"
	. "common"
)

func init() {

	Router.Register("/find", func(requireHandle RequireHandle) (result Result) {
		return result.SetData(FindAll())
	})

	Router.Register("/save", func(requireHandle RequireHandle) (result Result) {
		name := GetStringValue(requireHandle.Form["name"])
		id :=  GetStringValue(requireHandle.Form["id"])
		return SaveUser(name, id)
	})

	Router.Register("/look", func(requireHandle RequireHandle) (result Result) {
		fmt.Println("name:", GetStringValue(requireHandle.Form["name"]))
		fmt.Println("houseId:", GetStringArray(requireHandle.Form["houseId"]))
		fmt.Println("id:", GetBoolValue(requireHandle.Form["id"]))

		user := User{}
		SetStructEntity(requireHandle, &user)
		fmt.Println(user)

		return result.SetData(user)
	})
}
