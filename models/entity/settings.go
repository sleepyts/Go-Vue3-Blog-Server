package entity

import "Go-Vue3-Blog-Server/globalVar"

type Settings struct {
	AboutMe           string `json:"aboutMe"`
	ICP               string `json:"icp"`
	AboutMePageSongId string `json:"aboutMePageSongId"`
	IndexName         string `json:"indexName"`
	IndexUrl          string `json:"indexUrl"`
	Description       string `json:"description"`
	LogoUrl           string `json:"logoUrl"`
	Announcement      string `json:"announcement"`
}

func (Settings) TableName() string {
	return "tb_settings"
}

func GetSettings() Settings {
	var settings Settings
	err := globalVar.Db.First(&settings).Error
	if err != nil {
		return Settings{}
	}
	return settings
}
