package dao

type Sys_user struct {
	Id   string `xorm:"pk"`
	Username string
	Password string
	Nickname string
	Status int
}