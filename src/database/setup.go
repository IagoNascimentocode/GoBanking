package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/models"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Falied to connect to database")
	}

	database.AutoMigrate(&models.Account{})

	DB = database
}
