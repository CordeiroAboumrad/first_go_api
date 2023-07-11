package db

import (
	"database/sql"
	"fmt"

	"github.com/CordeiroAboumrad/first_go_api/configs"
	_ "github.com/lib/pq"
)

func openConnection() (*sql.DB, error) {
	config := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		// Mudar para produção
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
