package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"os"
	//"os/exec"
	//"strconv"
)

type alumno struct {
	nombre        string
	calificacion1 string
	calificacion2 string
	calificacion3 string
}

var estudiantes []alumno

func main() {

	var seleccion int
	var flag bool = false

	for flag == false {

		menu := ` 	
				[1] capturar alumno
				[2] Guardar
				[3] Leer
				[4] salir
				Elige una opcion `

		//limpiar_pantalla()
		fmt.Print(menu)
		fmt.Scanln(&seleccion)

		switch seleccion {
		case 1:
			for i := 0; i < 5; i++ {
				a := alumno{}
				fmt.Println("Captura el nombre del estudiante: ")
				fmt.Scanln(&a.nombre)
				fmt.Print("capture calificacion 1: ")
				fmt.Scanln(&a.calificacion1)
				fmt.Print("capture calificacion 2: ")
				fmt.Scanln(&a.calificacion2)
				fmt.Print("capture calificacion 3: ")
				fmt.Scanln(&a.calificacion3)

				//for j := 0; j < 3; j++ {
				//fmt.Printf("captura calificaciones %d:  ", j+1)
				//var temp string
				//fmt.Scanln(&temp)2
				//tempint, _ := strconv.Atoi(temp)
				//calificaciones[j+(3*i)] = tempint
				//}

				estudiantes = append(estudiantes, a)
			}
		case 2:
			fmt.Print(estudiantes)
			guardar()
		case 3:

			// Read test.txt
			data, err := ioutil.ReadFile("martin.txt")
			if err != nil {
				panic(err)
			}

			// Should output: `The file contains: Hello, world!`
			fmt.Printf("%s", data)

		}

		//pausa()

	}
	//func limpiar_pantalla() {
	//	cmd := exec.Command("cmd", "/c", "cls")
	//	cmd.Stdout = os.Stdout
	//	cmd.Run()
	//
	//func pausa() {
	//var pausar string
	//fmt.Println("Press enter to continue. ")
	//fmt.Scanln(&pausar)
	//}
}

func guardar() {
	f, _ := os.Create("martin.txt")

	defer f.Close()
	//w := bufio.NewWriter(f)

	data := ""

	for i := 0; i < len(estudiantes); i++ {
		data += estudiantes[i].nombre + estudiantes[i].calificacion1 + estudiantes[i].calificacion2 + estudiantes[i].calificacion3

	}
	fmt.Println(data)
	_, err := f.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}
}
