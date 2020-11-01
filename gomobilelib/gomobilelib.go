package gomobilelib

import (
	"log"
)

var indexServer int = 0

var flagInitServer bool = false

var serv *Server

//var jc server.JavaCallBack

func RegisterJavaCallBack(c JavaCallBack) {
	serv.Jc = c
	//jc = c
	//serv.ServerAddJavaCallBack(c)
}

func StartServer() {
	log.Println("indexServer: ", indexServer)

	if flagInitServer == false {
		serv = NewServer()
		serv.ConfigureRouter()
		go func() {
			Start(serv)
		}()
		flagInitServer = true
	}
	indexServer++
}
