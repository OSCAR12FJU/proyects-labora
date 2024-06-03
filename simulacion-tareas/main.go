// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // )

// // // // // En este caso vamos a necesitar recibir por parametro el canal y no por que vamosa hacer uso de su valor sino para enviarle y asignarle el valor
// // // // func enviarNumeros(c chan int) {
// // // // 	for i := 1; i <= 5; i++ {
// // // // 		c <- i
// // // // 		//Aca inicia el envio de los valores por el gourutines en el cual tomamos del un bucle el valor de un elemento y a este lo vamos pasamos a un canal
// // // // 	}
// // // // 	close(c)
// // // // } //Cerrar el canal despues de enviar todos los valores}

// // // // func recibirNumeros(c chan int) {
// // // // 	for numero := range c {
// // // // 		fmt.Println("Numeros recibidos:", numero)
// // // // 		//Este es un paso despues en la cual recibimos ese valor que nos llego de la funcion anterior y la destructuramos para mostrarla.
// // // // 	}
// // // // } //Recibir y mostrar números del canal

// // // //	func main() {
// // // //		canal := make(chan int) //Crear una canal de tipo int
// // // //		//iniciar gourutines para enviar y recibir númeors a traves del canal
// // // //		go enviarNumeros(canal)
// // // //		recibirNumeros(canal)
// // // //		//En este bloque antes de realizar la ejecucución de ambas funciones, creamos el canal el cual va a estar vacio pero declarado y va hacer usado para el almacenamiento.
// // // //	}

// // // package main

// // // import (
// // // 	"fmt"
// // // 	"math/rand"
// // // 	"sync"
// // // 	"time"
// // // )

// // // type Corredor struct {
// // // 	Nombre             string
// // // 	Velocidad          int
// // // 	DistanciaRecorrida int
// // // }

// // // func avanzarCompetidor(competidor Corredor, canalGanadores chan string, wg *sync.WaitGroup) {
// // // 	defer wg.Done()

// // // 	for {
// // // 		distanciaAvance := rand.Intn(10) + 1
// // // 		competidor.DistanciaRecorrida += distanciaAvance

// // // 		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)+50))

// // // 		fmt.Printf("%s ha avanzado %d unidades\n", competidor.Nombre, distanciaAvance)

// // // 		if competidor.DistanciaRecorrida >= 100 {
// // // 			canalGanadores <- competidor.Nombre
// // // 			break
// // // 		}
// // // 	}
// // // }

// // // func main() {
// // // 	rand.Seed(time.Now().UnixNano())

// // // 	canalGanadores := make(chan string)

// // // 	competidores := []Corredor{
// // // 		{"Competidor 1", rand.Intn(5) + 1, 0},
// // // 		{"Competidor 2", rand.Intn(5) + 1, 0},
// // // 		{"Competidor 3", rand.Intn(5) + 1, 0},
// // // 	}

// // // 	var wg sync.WaitGroup

// // // 	for _, competidor := range competidores {
// // // 		wg.Add(1)
// // // 		go avanzarCompetidor(competidor, canalGanadores, &wg)
// // // 	}
// // // 	wg.Wait()
// // // 	close(canalGanadores)
// // // 	ganador := <-canalGanadores
// // // 	fmt.Printf("\n%s ha ganado la carrera! \n", ganador)
// // // }

// // package main

// // import (
// // 	"fmt"
// // 	"math/rand"
// // 	"sync"
// // 	"time"
// // )

// // type Tareas struct {
// // 	Nombre    string
// // 	Proceso   int
// // 	Terminado int
// // }

// // func TrabajoTareas(t *Tareas, tarea chan string, wg *sync.WaitGroup) {
// // 	defer wg.Done()

// // 	for {
// // 		tiempoTarea := rand.Intn(10) + 1
// // 		t.Proceso += tiempoTarea

// // 		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)+50))

// // 		fmt.Printf("%s ha avanzado %d unidades\n", t.Nombre, tiempoTarea)

// // 		if t.Proceso >= 100 {
// // 			canalTarea <- t.Nombre
// // 			break
// // 		}
// // 	}

// // }

// // func main() {

// // 	rand.Seed(time.Now().UnixNano())
// // 	canalTarea := make(chan string)

// // 	tareas := []Tareas{
// // 		{"Tarea 1", rand.Intn(5) + 1, 0},
// // 		{"Tarea 2", rand.Intn(5) + 1, 0},
// // 		{"Tarea 3", rand.Intn(5) + 1, 0},
// // 	}

// // 	var wg sync.WaitGroup

// // 	for _, tarea := range tareas {
// // 		wg.Add(1)
// // 		go TrabajoTareas(tarea, canalTarea, &wg)
// // 	}
// // 	wg.Wait()
// // 	close(canalTarea)
// // 	finalizada := <-canalTarea
// // 	fmt.Printf("\n%s se ha finalizado \n", finalizada)

// // }
// package main

// import "fmt"

// func Search(l []string , prefix string) bool {

// 	var letEncontrado string

// 	for _, char := range prefix {
// 		if letters[char] == nil {
// 			return false
// 		}
// 		letEncontrado = letters[char]
// 	}
// 	return true
// }

// func main(){
// 	letters := []string{"app", "pala","saco","arco"}
// 	fmt.Println(trie.Search(letters, "ap"))
// }
