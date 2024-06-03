package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100),
		created_at TIMESTAMP NOT NULL DEFAULT  now(),
		updated_at TIMESTAMP,
		CONSTAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
)

type PsqlInvoiceHeader struct {
	db *sql.DB
}

func NewPsqlInvoiceHeader(db *sql.Db) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
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
