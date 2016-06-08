package common

import "github.com/jinzhu/gorm"

var dbModelData = &DBModel{
	list: []interface{}{},
}

type DBModel struct {
	list []interface{}
}

func (self *DBModel) add(model interface{}) {
	self.list = append(self.list, model)
}

func RegisterDBModel(model interface{}) {
	dbModelData.add(model)
}

func AutoMigrate(db *gorm.DB) {
	//db.AutoMigrate(dbModelData.list)
	for _, value := range dbModelData.list {
		db.AutoMigrate(value)
	}
}
