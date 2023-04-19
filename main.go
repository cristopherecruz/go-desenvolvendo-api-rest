package main

import (
	"github.com/cristopherecruz/go-desenvolvendo-api-rest/database"
	"github.com/cristopherecruz/go-desenvolvendo-api-rest/routes"
)

func main() {

	database.ConectarBancoDados()

	routes.HandleRequest()
}
