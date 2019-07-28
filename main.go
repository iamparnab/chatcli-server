package main

import (
	chathandlers "chat_handlers"
	"fmt"
	"net"
	"os"
)

func main() {
	//Create User Map
	allUsers := make(map[string]net.Conn, 0)

	port := os.Getenv("PORT")
	if port == "" {
		port = "2222"
	}
	server, serverError := net.Listen("tcp", ":"+port)

	if serverError != nil {
		fmt.Println("Error: ", serverError)
	} else {
		go fmt.Println("Server is Up and Running at port ", port)
		for {
			conn, connError := server.Accept()
			if connError != nil {
				go fmt.Println("Error: ", connError)
			} else {
				go chathandlers.SocketHandler(conn, allUsers)
			}
		}
	}
}
