package main

import (
	"log"

	"github.com/xPeacosta/go-react-social/bd"
	"github.com/xPeacosta/go-react-social/handlers"
)

func main() {
	if !bd.CheckConnection() {
		log.Fatal("No existe conexión a la Base de Datos")
		return
	}
	handlers.Manejadores()
}
