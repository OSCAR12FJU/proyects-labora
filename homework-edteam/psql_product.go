package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25),
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT  now(),
		updated_at TIMESTAMP,
		CONSTAINT products_id_pk PRIMARY KEY (id)
	)`
)

type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.Db) *PsqlProduct {
	return &PsqlProduct{db}
}

func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err := stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("migrción de producto ejecutada correctamente")
	return nil
}
