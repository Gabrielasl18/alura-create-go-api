package database 

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
)

var (
	DB *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port= 5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if(err != nil) {
		log.Panic("Erro ao conectar com o Banco de Dados")
	}
}
