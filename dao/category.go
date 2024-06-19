package dao

type Category struct {
	ID   int    `json:"id"`
	Name string `gorm:"size:20;not null;unique" json:"name"`
}
