package main

import (
	"bingo/internal/server"
	"fmt"
)

func main() {

	server := server.NewServer()
	println("Server is running")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
