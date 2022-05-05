package db

import (
	dbConn "cenco-pim/adapter/gorm"
	"cenco-pim/app/entity"
	"cenco-pim/config"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	appConf := config.AppConfig()
	db, _ := dbConn.New(&appConf.Db)
	err := db.AutoMigrate(&entity.FamilyGroup{}, &entity.Family{}, &entity.FamilyExample{}, &entity.Attribute{})
	if err != nil {
		return nil
	}
	return db
}
