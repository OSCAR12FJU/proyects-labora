package main

//prueba tecnica de jugadores de football

import (
	"fmt" //Paquete con metodos para mostrar (en pantalla o console)
	"os"
	"strings" // Paquete con metodos para poder trabajar con los strings
	"time"
)

var (
	adminUsername      = "admin"
	adminPassword      = "root"
	supervisorUsername = "seeker"
	supervisorPassword = "seekr"

	adminActions  = []string{"Crear laborer", "Eliminar laborer", "Generar archivo de texto con mensajes personalizado"}
	laborers      []string
	actionCounter = make(map[string]int)
)

type Player struct {
	Name       string //nombre de los jugadores
	Identifier int    //Identificador que va a tener
	//Ambas son propiedades de la estructura los cuales ya estan con el valor tipado (texto y numero entero positivo)
}

func main() {
	// Lista de jugadores actuales
	playerFootball := playerFootball()
	fizzBuzz := fizzBuzz()

	fmt.Println(playerFootball)
	fmt.Println()
	fmt.Println(fizzBuzz)

	for {
		fmt.Println("Bienvenido a Labora")
		fmt.Println("1. Iniciar sesión como Administrador")
		fmt.Println("2. Iniciar sesión como Supervisor")
		fmt.Println("3. Salir")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			adminLogin()
		case 2:
			supervisorLogin()
		case 3:
			fmt.Println("!Hasta Luego¡")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo")
		}
	}

}

func playerFootball() []string {
	players := []string{"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}
	//Array Literal que contiene todos los nombres de los jugadores

	// Registro de nuevos jugadores
	newPlayers := []string{"New Player 1", "New Player 2", "New Player 3", "New Player 4", "New Player 5", "New Player 6", "New Player 7"}
	//Array Literal que contiene todos los nombres de los jugadores

	// Eliminar los primeros 7 jugadores de la liga anterior
	players = players[7:]
	//Creamos una copia de un slice ya existente , en cual seleccionamos los elementos que necesitamos de ese, pasandole como primer parametro la posición que qeuremos arrancar hasta la que queremos cortar.

	// Agregar nuevos jugadores
	for _, name := range newPlayers {
		players = append(players, name)
	}
	//Aca tommando como base la cantidad de elemento que se encuentra en el array newPlayers vamos dando vueltas segun los elementos que este contenga, y lo estamos haciendo con los slices que estamos reutilzando es crear otros y ahcer qeu la variable apunte hacia esos, volviendo a reasignar la variable y que contenga otro contenido.

	// Mostrar lista completa de jugadores
	fmt.Println("Lista completa de jugadores:")
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1001, sanitizeName(player))
		//Aca estamos haciendo otro bucle for el cual en este en el bucle recorremos el slice que ya creamos con los nuevos jugadores y a este le recorremos y guardamos en variables el indice de cada uno y el valor tambien de cada jugador.
		//Para despues mostrar en un texto concatenado el numero identificador que va a hacer desde 1001
	}
	return players
}

// Función para eliminar caracteres especiales y números del nombre del jugador
func sanitizeName(name string) string {
	// Eliminar caracteres especiales y números
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
			//Lo que se realiza en esta comparación primero es de cada valor el cual recorre el array que le pasamos por parametro esta siendo destructurado y siendo guardado en la la variable r y siendo convertida al mismo tiempo en un rune.
			//Este hace la comparación con estas letras las cuales estas tambien estan pasadas a valore run el cual estas tienen un valor run desde el comienzo final o se si la letra que destructuramos y comparamos es menor a esta es por que esa no es una letra sino un signo.
			//Despues en lo que retornamos es o la letra la cual no se va almacenar en ningun lado ya que no estamos tomando un objeto fijo sino que estamos haciendo modificación directa, y la que no coincida y termine siendo un signo esta va a ejecutar un -1 y esta la va a terminar ignorando y borrando de el string.
		}
		return -1
	}, name)
}

func fizzBuzz() string {
	n := 50
	var resultado string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			resultado += ("FizzBuzz")
		} else if i%3 == 0 {
			resultado += ("Fizz")
		} else if i%5 == 0 {
			resultado += ("Buzz")
		} else {
			resultado += fmt.Sprint(i)
		}
		if i < n {
			resultado += ","
		}
	}
	return resultado

}

/////Pruebas tecnica packages/////

func adminLogin() {
	fmt.Println("Ingrese las credenciales de Admin")
	fmt.Print("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == adminUsername && password == adminPassword {
		fmt.Println("Inicio de sesión exitoso como Admin")
		adminMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo")
	}
}

func supervisorLogin() {
	fmt.Println("Ingrese las credenciales de Supervisor")
	fmt.Print("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == supervisorUsername && password == supervisorPassword {
		fmt.Println("Inicio de sesión exitoso como Supervisor")
		supervisorMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo.")
	}
}

func adminMenu() {
	for {
		fmt.Println("Menú de Admin")
		for i, action := range adminActions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Println("0. Cerrar sesión")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createLaborer()
		case 2:
			deleteLaborer()
		case 3:
			fmt.Println("Accción: Generar archivo de texto con mensajes personalizados")
			generateTextFile()
		case 0:
			fmt.Println("Cerrando sesión de Admin")
			return
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func createLaborer() {
	laborers = append(laborers, fmt.Sprintf("laborer %d", len(laborers)+1))
	fmt.Println("Laborer creado exitosamente.")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}
}

func deleteLaborer() {
	if len(laborers) == 0 {
		fmt.Println("No hay laborers para eliminar.")
		return
	}

	laborers = laborers[:len(laborers)-1]
	fmt.Println("Laborer eliminado exitosamente.")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}
}

func supervisorMenu() {
	for {
		fmt.Println("Manu de Supervisor")
		fmt.Println("1. Crear registro de Admin")
		fmt.Println("0. Cerra sesion")

		var choice int
		fmt.Print("Ingrese su opcio: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Accion: Crear registro de Admin")
			createAdminRecord()
		case 0:
			fmt.Println("Cerrando sesión de Supervisor")
			return
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo")
		}
	}
}

func generateTextFile() {
	fileName := fmt.Sprintf("mensajes_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo: %v\n", err)
		return
	}
	defer file.Close()

	for _, laborer := range laborers {
		_, err := file.WriteString(fmt.Sprintf("%s\n", laborer))
		if err != nil {
			fmt.Printf("Error al escribir en el archivo: %v\n", err)
			return
		}
	}
	fmt.Printf("Archivo %s creado exitosamente. \n", fileName)
}

func createAdminRecord() {
	fileName := fmt.Sprintf("registro_admin_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo de registro: %v\n", err)
		return
	}
	defer file.Close()

	for _, action := range adminActions {
		_, err := file.WriteString(fmt.Sprintf("%s:%d\n", action, actionCounter[action]))
		if err != nil {
			fmt.Printf("Error al escribir en el archivo de registro: %v\n", err)
			return
		}
	}
	fmt.Printf("archivo de registro %s creado exitosamente.\n", fileName)
}
