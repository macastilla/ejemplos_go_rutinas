package main

import (
	"fmt"
	"time"
)

func descargaArchivo() {
	time.Sleep(time.Second * 5)
	fmt.Println("Se descargo el archivo")
}

func lecturaArchivo() {
	time.Sleep(time.Second * 2)
	fmt.Println("Se completo de leer el archivo")
}

func main() {

	go descargaArchivo()
	go lecturaArchivo()

	time.Sleep(time.Second * 6)

	fmt.Println("Finaliza el programa")
}