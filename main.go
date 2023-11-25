package main 

import (
	"fmt"
	"API-GO-REST/routes"
	"API-GO-REST/database"
)

func main() {
	database.ConectaComBancoDeDados();
	fmt.Println("Iniciando o servidor Rest com Go");
	routes.HandleRequest()
}