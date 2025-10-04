package db

import (
	"fmt"
	"log"

	"github.com/Sa-Leonardo/WorkPoint/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
    host := "localhost"
    port := 5454
    user := "postgres"
    password := "1234"
    dbname := "postgres"

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        host, user, password, dbname, port,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco: %v", err)
    }

    // Auto-migrate das tabelas do sistema de ponto
    err = db.AutoMigrate(
        &models.Empresa{},
        &models.Usuario{},
        &models.RegistroPonto{},
        &models.Justificativa{},
    )
    if err != nil {
        log.Fatalf("Erro ao migrar tabelas: %v", err)
    }

    fmt.Println("Conectado ao banco com GORM e tabelas migradas!")
    return db
}
