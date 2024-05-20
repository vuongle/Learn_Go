package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	UserId string `json:"user_id"`
	User   User   `json:"user" gorm:"foreignkey:UserId"`
}
