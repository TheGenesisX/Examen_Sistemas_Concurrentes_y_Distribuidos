package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

var clientList = []net.Conn{}

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

		for i := 0; i < len(clientList); i++ {
			if cli != clientList[i]{
				err := gob.NewEncoder(clientList[i]).Encode(mensaje)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)
}
