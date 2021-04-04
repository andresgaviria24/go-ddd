package entity

type Users struct {
	Id       string `gorm:"TYPE:char(40);PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	UserName string `gorm:"NOT NULL;COLUMN:userName"`
	Password string `gorm:"NULL;COLUMN:password"`
}

func (Users) TableName() string {
	return "users"
}
