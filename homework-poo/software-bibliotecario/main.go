package main

import (
	"fmt"
)

type EstadoLibro string
type libro struct {
	ID     int
	titulo string
	autor  string
	genero string
	estado EstadoLibro
}

const (
	Disponible EstadoLibro = "disponible"
	Prestado   EstadoLibro = "prestado"
)

func main() {
	biblioteca := map[int]libro{
		1: {ID: 1, titulo: "El mago de Oz", autor: "steven hoquing", genero: "fantasia", estado: "disponible"},
		2: {ID: 2, titulo: "Harry Potter", autor: "angelica smith", genero: "suspenso", estado: "disponible"},
		3: {ID: 3, titulo: "El Hobbit", autor: "Christopher Jhonson", genero: "ficcion", estado: "disponible"}}

	nuevoLibro := libro{ID: 4, titulo: "Aladin", autor: "Oscar Olla", genero: "infantil", estado: "disponible"}
	nuevoLibro.AgregarLibro(biblioteca, nuevoLibro)
	nuevoLibro.BusquedaLibro(biblioteca, "El Hobbit", "Chistopher")

}

func (l *libro) AgregarLibro(biblioteca map[int]libro, nuevoLibro libro) {

	for _, lib := range biblioteca {
		if lib.titulo == nuevoLibro.titulo {
			fmt.Println("Este libro ya existe en la biblioteca")
			return
		}
	}
	biblioteca[nuevoLibro.ID] = nuevoLibro
	fmt.Println("La nueva tarea fue agregada a la lista:", nuevoLibro)
}

func (l *libro) BusquedaLibro(biblioteca map[int]libro, titulo string, autor string) {

	for _, lib := range biblioteca {
		if lib.titulo == titulo || lib.autor == autor {
			fmt.Println("el libro", lib.titulo, "fue encontrado")
			return
		}
	}
	fmt.Println("No tenemos ese libro en la biblioteca")

}
