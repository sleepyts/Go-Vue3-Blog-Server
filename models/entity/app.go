package entity

import "Go-Vue3-Blog-Server/globalVar"

type App struct {
	Id      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

func (App) TableName() string {
	return "tb_app"
}

func GetApp() []App {
	var apps []App
	globalVar.Db.Find(&apps)
	return apps
}

func AddApp(app App) error {
	return globalVar.Db.Create(&app).Error
}

func DeleteApp(id uint) error {
	return globalVar.Db.Delete(&App{Id: id}).Error
}

func UpdateApp(app App) error {
	return globalVar.Db.Model(&app).Updates(app).Error
}
