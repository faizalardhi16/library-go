package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Transaction{}},
		{Model: Book{}},
		{Model: Reservation{}},
	}
}

func InitializedDB(server *gorm.DB) {

	//db Migration
	for _, model := range RegisterModels() {

		err := server.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println("Success Migrate the Table!")

}
