package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/models"
)

var DB *gorm.DB

func ConnectDataBase() {

	connection, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Falied to connect to database")
	}

	connection.AutoMigrate(&models.Account{})

	DB = connection
}
