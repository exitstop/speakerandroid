package main

import "github.com/exitstop/speakerandroid/internal/server"

func main() {
	serv := server.NewServer()
	serv.ConfigureRouter()
	server.Start(serv)
}
