package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//"github.com/rs/cors"
)

var db *sql.DB

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Error en cargar el arch .env")
	}
	//Realizando la conexión a la base de datos, esto lo hacemos llamando a la función que lo va a realizar.
	var err error
	db, err = createConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Probamos la conexión creando la base de datos, las tablas con las columnas y filas las creamos desde el Workbench.
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS `system_books`")
	if err != nil {
		panic(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/create-book", createBooksNew).Methods("POST")
	router.HandleFunc("/get-books", getBooks).Methods("GET")

	http.ListenAndServe(":8080", router)

}

// Función de la conexión a la BD
func createConnection() (*sql.DB, error) {
	conectionString := "root:Fuerzaabasto1@@tcp(127.0.0.1:3306)/system_books"
	db, err := sql.Open("mysql", conectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Conectado correctamente")
	return db, nil
}

// Estructuras de referencia de las tablas que tenemos creadas en MySQL
type Books struct {
	ID     int    `json:"book_id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Status *bool  `json:"status"`
}

// Función Controller para crear un libro tomando las propiedades de la estructura y subir los valores a la BD.
func createBooksNew(w http.ResponseWriter, r *http.Request) {
	var books Books
	json.NewDecoder(r.Body).Decode(&books)

	result, err := db.Exec(`INSERT INTO books (name, author) VALUES (?,?)`, books.Name, books.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	books.ID = int(lastInsertID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getSearchBooks(w http.ResponseWriter, r *http.Request) {
}

// Función Controller para buscar los objeto ya creado de la BD.
func getBooks(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT book_id, name, author, status FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Books

	for rows.Next() {
		var book Books
		err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
