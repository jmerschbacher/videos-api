package database

import (
	"github.com/jmerschbacher/videos-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MySQLStarter() *gorm.DB {
	conexao := "root:admin@tcp(127.0.0.1:3306)/db_alura?charset=utf8"
	DB, err := gorm.Open(mysql.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	err = DB.AutoMigrate(&entity.Video{})
	if err != nil {
		return nil
	}

	return DB
}
