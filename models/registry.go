package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: User{}},
		{Model: Category{}},
	}
}

func RegistryDatabase(db *gorm.DB) {

	for _, model := range RegisterModel() {
		err := db.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println("Success to migrate Database")

}
