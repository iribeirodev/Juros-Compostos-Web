package main

import (
	"iribeirodev/juros-compostos-web/routes"
	"log"
)

// @title Juros Compostos API
// @version 1.0
// @description API para simulação de juros compostos e parcelas.

// @BasePath /api/v1

func main() {
	router := routes.IniciarRotas()
	log.Println("Servidor rodando na porta 8080")
	router.Run(":8080")
}
