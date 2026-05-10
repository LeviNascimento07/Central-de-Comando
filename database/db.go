package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	host := "m23-0t.h.filess.io"
	port := "3307"
	user := "sgi_marketdead"
	password := os.Getenv("DB_PASSWORD")
	dbname := "sgi_marketdead"

	log.Printf("Conectando com usuário: %s e senha: %s", user, password)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexao com banco: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	log.Println("Banco de dados conectado com sucesso")
}
