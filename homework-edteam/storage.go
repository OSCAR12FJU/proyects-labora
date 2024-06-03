import (
	"fmt"
	"database/sql"
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("MySQL", "root:Fuerzaabasto1@@tcp(127.0.0.1:3306)/nuevoshema")
		if err != nil {
			log.Fatalf("No se pudo hacer la conexión: %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo establecer la conexión: %v", err)
		}
		fmt.Println("Conectado Correctamente")
	})
}
func Pool() *sql.DB {
	return db
}