package main

import (
	"log"
	// "fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// connectionString := "root:Fuerzaabasto1@@tcp(127.0.0.1:3306)/nuevoshema"

	// db, err := sql.Open("mysql", connectionString)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("Conexión establecida correctamente")
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	invoiceheader.NewService(storageInvoiceHeader)
	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}
}
