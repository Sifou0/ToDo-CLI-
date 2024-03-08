package main

import "gorm.io/gorm"

func create(db *gorm.DB, item Todo) {
	db.Create(&item)
}

// func delete(db *gorm.DB, id int) {
// 	db.Delete(&Todo{}, id)
// }

func update(db *gorm.DB, item Todo) {
	db.Save(item)
}

func getAll(db *gorm.DB) []Todo {
	var results []Todo
	db.Order("id asc").Find(&results)
	return results
}

func get(db *gorm.DB, id int) Todo {
	var result Todo
	db.First(&result, id)
	return result
}
