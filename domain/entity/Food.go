package entity

type Food struct {
	Id    string  `gorm:"NOT NULL;COLUMN:id"`
	Name  string  `gorm:"NOT NULL;COLUMN:name"`
	Price float32 `gorm:"NOT NULL;COLUMN:price"`
}

func (Food) TableName() string {
	return "food"
}
