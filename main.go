package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func startHTTP(allUsers map[string]net.Conn) {
	http.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		Users(res, req, allUsers)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	//Create User Map
	allUsers := make(map[string]net.Conn, 0)

	port := os.Getenv("PORT")

	server, serverError := net.Listen("tcp", ":"+"2222")

	if serverError != nil {
		fmt.Println("Error: ", serverError)
	} else {
		go startHTTP(allUsers)
		go fmt.Println("Server is Up and Running at port ", port)
		for {
			conn, connError := server.Accept()
			if connError != nil {
				go fmt.Println("Error: ", connError)
			} else {
				go SocketHandler(conn, allUsers)
			}
		}
	}
}
