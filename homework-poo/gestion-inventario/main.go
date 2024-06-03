package main

import (
	"fmt"
)

// Estructura del inventario
type Product struct {
	ID       int
	nombre   string
	precio   float64
	cantidad int
}

func main() {
	inventario := map[int]Product{
		1: {ID: 1, nombre: "arroz", precio: 120.50, cantidad: 3},
		2: {ID: 2, nombre: "fideos", precio: 230.50, cantidad: 9},
		3: {ID: 3, nombre: "lentejas", precio: 320.20, cantidad: 5},
	}
	nuevoProducto := Product{ID: 4, nombre: "aceite", precio: 320.50, cantidad: 4}
	nuevoProducto.AgregarPruduct(inventario, nuevoProducto)
	nuevoProducto.CambiarCantidad(inventario, "aceite", 9)
	nuevoProducto.EliminarProducto(inventario, "arroz")
	nuevoProducto.MostrarInventario(inventario)

}

//Funcion para agregar productos

// En la funcion de aca vamos a determinar primero qeu esta funcion es parte de la estructura
// Despues vamos a declarar el metodo el cual el va a recibir el inventario y el nuevo producto que le determinamos que va seguir los tipos de la estruct
// el inventario lo pedimos para poder hacer la validación de si este existe dentro del map.
func (p *Product) AgregarPruduct(inventario map[int]Product, nuevoProducto Product) {
	//Si este producto que lo buscamos por su ID existe que no se agregue
	// el ok nos retorna el false o true de la condicion
	if _, ok := inventario[nuevoProducto.ID]; ok {
		fmt.Println("El producto ya existe")
		return
	}
	inventario[nuevoProducto.ID] = nuevoProducto
	fmt.Println("Producto agregado al inventario", nuevoProducto)
}

//Funcion para actualizar cantidad de los productos

func (p *Product) CambiarCantidad(inventario map[int]Product, nombreProducto string, nuevaCantidad int) {
	for id, product := range inventario {
		if product.nombre == nombreProducto {
			product.cantidad = nuevaCantidad
			inventario[id] = product
			fmt.Println("Nueva cantidad cambiada del inventario")
			return
		}
	}
	fmt.Println("El producto no existe en el inventario")
}

// Funcion para eliminar prodcutos
func (p *Product) EliminarProducto(inventario map[int]Product, nombreProducto string) {
	var ID int
	var encontrado bool
	for id, product := range inventario {
		if product.nombre == nombreProducto {
			ID = id
			encontrado = true
			break
		}
	}
	if encontrado {
		delete(inventario, ID)
		fmt.Println("Ya se elimino del inventario el producto:", inventario[ID].nombre)
	} else {

		fmt.Println("El producto indicado no existe")
	}

}

func (p *Product) MostrarInventario(inventario map[int]Product) {
	for _, product := range inventario {
		fmt.Printf("Nombre: %s, Precio: %.2f, Cantidad: %d\n", product.nombre, product.precio, product.cantidad)
	}
}
