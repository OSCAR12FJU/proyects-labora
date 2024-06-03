// // // Ejercicio 1: Árbol de búsqueda de prefijos

// // // Diseña e implementa una estructura de datos en Go para un árbol de búsqueda de prefijos (Tree) que almacene un conjunto de palabras. Proporciona funciones para insertar palabras en el árbol y buscar palabras por prefijo.

package main

import "fmt"

//Creamos la estructura de los nodes
type TreeNode struct {
	children map[rune]*TreeNode
	isEnd    bool
}

//Creamos la cabeza principal
type Tree struct {
	root *TreeNode
}

//Esta es una funcion de la estructura la cual instancia directamente a la estructura.}
//Esta funcion se ejecutaria solo una vez durante la inicialización del Tree, esta funcion se utiliza para crear la estructura principal del Tree. Obviamente esta estructua creada no se queda en el aire, se va a ejeuctar desde el main en una variable
func NewTree() *Tree {
	//En la ejecución qeu le damos a este en el bloque siempre va hacer apuntando para que no se creen copias de la estructura sino trabajemos directamente con uno.
	return &Tree{root: &TreeNode{children: make(map[rune]*TreeNode) /*Dentro de esta estructura ingresamos a la unica propiedad que este tiene y lo instanciamos llamando a la estruductura de nodes hijos la cual apuntamos a ella para hacer usp de sus de sus propiedades.*/}}
}

//Este es una unción tambien, pero es un metodo el cual modifica, e inserta.
//Este metodo solamente se va a poder ejecutar cuando la estructura ya este instanciada y este como e sun metodo qeu modica al objeto, se ejecuta directo al tal pasandole un valor el cual ejecute nuestra función.

func (t *Tree) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TreeNode{children: make(map[rune]*TreeNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

//Este metodo es igual al otro realiza una funcionalidad con los objeto ya creados en este, en este caso recorre el arbol buscando atraves de un valor que se le pasa por parametro a este metodo.
func (t *Tree) Search(prefix string) bool {
	node := t.root
	//Guarda en una variable una propiedad del objeto, guardamos esta propiedad ya que esta tiene incluida como valor la estructura de nodo hijo.
	for _, char := range prefix {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return true
}

func main() {
	trie := NewTree()
	words := []string{"apple", "banana", "app", "bat", "ball", "cat"}

	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("apo"))

}

// // Ejercicio 3: Árbol de segmentos

// // Implementa un árbol de segmentos en Go que admita dos operaciones: actualizar un valor en un índice específico del arreglo subyacente y calcular la suma de un rango de valores en el arreglo. Proporciona ejemplos de uso para demostrar la eficiencia de tu implementación.

// package main

// import (
// 	"fmt"
// )

// type SegmentTree struct {
// 	root *SegmentTreeNode
// }

// type SegmentTreeNode struct {
// 	start, end int
// 	sum        int
// 	left, right *SegmentTreeNode
// }

// func NewSegmentTree(arr []int) *SegmentTree {
// 	// Implementa la inicialización del árbol de segmentos aquí
// 	return &SegmentTree{root: }

// }

// func (st *SegmentTree) Update(index, value int) {
// 	// Implementa la operación de actualización aquí
// }

// func (st *SegmentTree) Query(start, end int) int {
// 	// Implementa la operación de consulta (suma de rango) aquí
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	tree := NewSegmentTree(arr)

// 	// Ejemplos de uso: actualizar un valor y calcular la suma de un rango
// 	tree.Update(2, 10)
// 	sum := tree.Query(1, 4)
// 	fmt.Println("Suma del rango:", sum)
// }

// package main

// import(
// 	"fmt"

// )

// type LettersJuego struct{

// }

// type Juego struct{
// 	root   	*LettersJuego

// }

// func main(){

// }
