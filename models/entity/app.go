package entity

type App struct {
	Id      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

func (App) TableName() string {
	return "tb_app"
}
