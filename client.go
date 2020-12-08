package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

var login chan string = make(chan string)

// User ...
type User struct {
	Nombre  string
	Mensaje string
}

var messageChan chan User = make(chan User)
var archiveChan chan User = make(chan User)

func cliente(login chan string, messageChan chan User) {
	var receivedMessage string

	cli, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Sprintln(err)
		return
	}
	defer cli.Close()

	go func() {
		for {
			err := gob.NewDecoder(cli).Decode(&receivedMessage)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(receivedMessage)
		}
	}()

	for {
		select {
		case loginMessage := <-login:
			err = gob.NewEncoder(cli).Encode(loginMessage)
			if err != nil {
				fmt.Sprintln(err)
				return
			}

		case newMessage := <-messageChan:
			chatMessage := newMessage.Nombre + ": " + newMessage.Mensaje
			err = gob.NewEncoder(cli).Encode(chatMessage)
			if err != nil {
				fmt.Sprintln(err)
				return
			}

		case newArchive := <-archiveChan:
			chatMessage := newArchive.Nombre + " sent archive: " + filepath.Base(newArchive.Mensaje)

			err = gob.NewEncoder(cli).Encode(chatMessage)
			if err != nil {
				fmt.Sprintln(err)
				return
			}
			fmt.Print("Test")
			// err = gob.NewEncoder(cli).Encode(fileData)
			// if err != nil {
			// 	fmt.Sprintln(err)
			// 	return
			// }

		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var opc uint64
	var newUser User

	go cliente(login, messageChan)

	fmt.Print("Username: ")
	scanner.Scan()
	nombre := scanner.Text()
	newUser.Nombre = nombre

	login <- "Se ha conectado: " + nombre

	fmt.Println("----------MENU----------")
	fmt.Println("1) Enviar mensaje")
	fmt.Println("2) Enviar documento")
	fmt.Println("0) Salir")

	for {
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			scanner.Scan()
			mensaje := scanner.Text()
			newUser.Mensaje = mensaje

			messageChan <- newUser
		case 2:
			fmt.Print("Ruta absoluta del archivo: ")
			scanner.Scan()
			mensaje := scanner.Text()
			newUser.Mensaje = mensaje

			archiveChan <- newUser

			fileData, err := ioutil.ReadFile("Desktop/test.txt")
			if err != nil {
				fmt.Sprintln(err)
				return
			}
			fmt.Print(string(fileData))
		case 0:
			login <- "Se ha desconectado: " + nombre
			var input string
			fmt.Scanln(&input)

			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}

	var input string
	fmt.Scanln(&input)
}
