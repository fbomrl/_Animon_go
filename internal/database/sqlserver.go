package database

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func SqlServer(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir conexão: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar (ping): %w", err)
	}
	return db, nil
}
