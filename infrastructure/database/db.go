package database

import (
	"github.com/jmerschbacher/videos-api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func PostgreSQLStarter() *gorm.DB {
	conexao := "host=localhost user=root password=root dbname=db_alura port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	err = DB.AutoMigrate(
		&entity.Video{},
		&entity.Categoria{})
	if err != nil {
		log.Panic("Erro ao criar tabelas via Gorm")
	}

	return DB
}
