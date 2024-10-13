package entity

import "Go-Vue3-Blog-Server/globalVar"

type Admin struct {
	Id       uint
	Username string
	Password string
}

func (Admin) TableName() string {
	return "tb_admin"
}

func IsContain(admin Admin) bool {
	return globalVar.Db.Where("username =?", admin.Username).First(&Admin{}).RowsAffected > 0
}
