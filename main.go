package main

import (
	"Practica/clienteServidor/server"
	"Practica/clienteServidor/serverRespaldo"
)

func main( ){
	go server.Run()
	go serverrespaldo.Run()

	select {}

}