package persistence

import (
	"ws_restaurant/infrastructure/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbHelper struct {
	FoodRepository repository.FoodRepository
	db             *gorm.DB
}

func InitDbHelper() (*DbHelper, error) {

	dsn := "userName:password@tcp(localhost)/restaurant?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
	return &DbHelper{
		FoodRepository: &FoodRepositoryImpl{db},
		db:             db,
	}, nil
}
