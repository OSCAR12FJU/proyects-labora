package main

import (
	"fmt"
)

type Tareas struct {
	ID          int
	nombre      string
	descripcion string
	responsable string
	estado      EstadoTarea
}
type EstadoTarea string

const (
	Pendiente  EstadoTarea = "pendiente"
	EnProgreso EstadoTarea = "en progreso"
	Completada EstadoTarea = "completada"
)

func main() {

	tareasActuales := map[int]Tareas{
		1: {ID: 1, nombre: "Diseño UI UX", descripcion: "Portada de landing", responsable: "Martin Chau", estado: "pendiente"},
		2: {ID: 2, nombre: "Desarrollo de BD", descripcion: "Estructura para recibir la data de los clientes", responsable: "Oscar Flores", estado: "pendiente"},
		3: {ID: 3, nombre: "Diseño Frontend", descripcion: "Armado de la landing", responsable: "Nahuel Tisera", estado: "pendiente"}}

	nuevaTarea := Tareas{ID: 4, nombre: "Cloud", descripcion: "Almacenamiento en el cloud", responsable: "German", estado: "pendiente"}
	nuevaTarea.AgregarTarea(tareasActuales, nuevaTarea)
	nuevaTarea.CambiarResponsable(tareasActuales, "Diseño Frontend", "Oscar Flores")
	nuevaTarea.CambiarEstado(tareasActuales, "Desarrollo de BD", "completada")
	nuevaTarea.MostrarTareasP(tareasActuales)

	// for _, tarea := range tareasActuales {
	// 	fmt.Println(tarea, "\n")
	// }
}

//Funcion para agregar tarea

func (t *Tareas) AgregarTarea(tareasActuales map[int]Tareas, nuevaTarea Tareas) {

	for _, tarea := range tareasActuales {
		if tarea.nombre == nuevaTarea.nombre {
			fmt.Println("Esta tarea ya existe en la lista")
			return
		}
	}
	tareasActuales[nuevaTarea.ID] = nuevaTarea
	fmt.Println("La nueva tarea fue agregada a la lista:", nuevaTarea)
}

// Funcion para cambiar asignación de tareas
func (t *Tareas) CambiarResponsable(tareasActuales map[int]Tareas, tareaCambiar string, nuevoResponsable string) {
	for _, tarea := range tareasActuales {
		if tarea.nombre == tareaCambiar {
			tarea.responsable = nuevoResponsable
			fmt.Println("Nuevo responsable asignado a la tarea:", tarea.nombre)
			return
		}
	}
	fmt.Println("La tarea", tareaCambiar, "no se encontró en la lista actual")
}

func (t *Tareas) CambiarEstado(tareasActuales map[int]Tareas, tareaCambiar string, estadoTarea EstadoTarea) {
	for _, tarea := range tareasActuales {
		if tarea.nombre == tareaCambiar {
			tarea.estado = estadoTarea
			fmt.Println("\n Nuevo estado asignado a la tarea:", tarea.estado)
			return
		}
	}
	fmt.Println("El estado", tareaCambiar, "no se pudo cambiar")
}

func (t *Tareas) MostrarTareasP(tareasActuales map[int]Tareas) {
	for _, tarea := range tareasActuales {
		if tarea.estado == "pendiente" {
			fmt.Printf("Nombre: %s, Descripción: %s, Responsable: %s Estado: %s\n", tarea.nombre, tarea.descripcion, tarea.responsable, tarea.estado)
		}
	}
}
