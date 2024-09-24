package entity

import "Go-Vue3-Blog-Server/globalVar"

type Category struct {
	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (Category) TableName() string {
	return "tb_category"
}

func GetCategoryList() ([]Category, error) {
	var categoryList []Category
	err := globalVar.Db.Find(&categoryList).Error
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}
