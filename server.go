package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

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
	}
}

func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)
}
