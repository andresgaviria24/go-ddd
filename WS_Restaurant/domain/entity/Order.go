package entity

type Order struct {
	IdUser    string `gorm:"NOT NULL;COLUMN:idUser"`
	FirstName string `gorm:"NOT NULL;COLUMN:firstname"`
	LastName  string `gorm:"NOT NULL;COLUMN:lastname"`
	Email     string `gorm:"NOT NULL;COLUMN:email"`
	IdRol     string `gorm:"NOT NULL;COLUMN:idRol"`
	Rol       string `gorm:"NOT NULL;COLUMN:rol"`
	TenantId  string `gorm:"NOT NULL;COLUMN:tenantId"`
}

func (Order) TableName() string {
	return "order"
}
