package db

import (
	"serve-ressources/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	cfg.DB.User,
	// 	cfg.DB.Password,
	// 	cfg.DB.Host,
	// 	cfg.DB.Port,
	// 	cfg.DB.Name,
	// )

	// fmt.Println(dataSourceName)

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
