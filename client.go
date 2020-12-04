package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

var canal chan string = make(chan string)

func cliente(canal chan string) {
	cli, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Sprintln(err)
		return
	}

	for {
		mensaje := <-canal
		err = gob.NewEncoder(cli).Encode(mensaje)
		if err != nil {
			fmt.Sprintln(err)
		}
	}
	// cli.Close()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var opc uint64

	go cliente(canal)

	fmt.Print("Username: ")
	scanner.Scan()
	nombre := scanner.Text()

	canal <- "Se ha conectado: " + nombre

	for {
		fmt.Println("1) Enviar mensaje")
		fmt.Println("2) Enviar documento")
		fmt.Println("3) Mostrar chat")
		fmt.Println("4) Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			fmt.Print("Mensaje: ")
			scanner.Scan()
			mensaje := scanner.Text()

			canal <- nombre + ": " + mensaje
		}
	}

	var input string
	fmt.Scanln(&input)
}
