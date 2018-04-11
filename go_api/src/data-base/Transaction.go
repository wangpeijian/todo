package data_base

import (
	"github.com/go-xorm/xorm"
	. "core"
)

func Transaction(engine *xorm.Engine, result *Result, cb func(session *xorm.Session)) {
	session := engine.NewSession()
	err := session.Begin()

	defer func() {
		defer session.Close()

		if err := recover(); err != nil {
			session.Rollback()
			res := Result{}.SetError(err.(string))
			*result = res
		}

	}()

	cb(session)

	err = session.Commit()
	if err != nil {
		panic(ERROR)
	}
}
