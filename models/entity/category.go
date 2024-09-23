package entity

type Category struct {
	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

func (Category) TableName() string {
	return "tb_category"
}
