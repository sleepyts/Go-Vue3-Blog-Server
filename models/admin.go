package model

type Admin struct {
	Id       uint
	Username string
	Password string
}

func (Admin) TableName() string {
	return "tb_admin"
}