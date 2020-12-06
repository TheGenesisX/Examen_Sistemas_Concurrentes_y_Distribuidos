package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

var clientList = []net.Conn{}
var chat = []string{}

func servidor() {
	serv, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Sprintln(err)
		return
	}

	for {
		cli, err := serv.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		clientList = append(clientList, cli)
		go clientHandler(cli)
	}
}

func clientHandler(cli net.Conn) {
	var mensaje string

	for {
		err := gob.NewDecoder(cli).Decode(&mensaje)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(mensaje)
		chat = append(chat, mensaje)
		logoutFlag := strings.Contains(mensaje, "Se ha desconectado:")

		if logoutFlag == true {
			for i := 0; i < len(clientList); i++ {
				if cli == clientList[i] {
					clientList = append(clientList[:i], clientList[i+1:]...)
				}
			}
		}

		for i := 0; i < len(clientList); i++ {
			if cli != clientList[i] {
				err := gob.NewEncoder(clientList[i]).Encode(mensaje)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func saveChat() {
	file, err := os.Create("chatBackup.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, message := range chat {
		file.WriteString(message + "\n")
	}
}

func main() {
	var opc uint64
	go servidor()

	fmt.Println("Iniciando servidor.")

	fmt.Println("----------MENU----------")
	fmt.Println("1) Guardar chat.")
	fmt.Println("0) Salir.")
	fmt.Println("Chat en vivo: ")

	for {
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			saveChat()
			fmt.Println("Copia del chat guardada con exito.")
		case 0:
			fmt.Println("Cerrando servidor.")
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}

	var input string
	fmt.Scanln(&input)
}
