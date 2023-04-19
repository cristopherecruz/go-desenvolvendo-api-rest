package controllers

import (
	"github.com/cristopherecruz/go-desenvolvendo-api-rest/database"
	"github.com/cristopherecruz/go-desenvolvendo-api-rest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.String(http.StatusOK, "Home Page")
}

func BuscarTodasPersonalidades(c *gin.Context) {
	var p []models.Personalidade
	dados := database.Db
	dados.Find(&p)

	c.JSON(http.StatusOK, p)
}

func BuscarPersonalidadePorId(c *gin.Context) {

	var p models.Personalidade

	id := c.Param("id")

	encontrou := false

	database.Db.First(&p, id)

	if &p != nil {
		c.JSON(http.StatusOK, p)
		encontrou = true
	}

	if !encontrou {
		c.String(http.StatusNotFound, "Personalidade não encontrada!")
	}
}

func CriarPersonalidade(c *gin.Context) {

	var novaPersonalidade models.Personalidade

	err := c.ShouldBindJSON(&novaPersonalidade)
	if err != nil {
		panic(err.Error())
	}

	database.Db.Create(&novaPersonalidade)

	c.JSON(http.StatusCreated, novaPersonalidade)
}

func DeletarPersonalidade(c *gin.Context) {

	var p models.Personalidade

	id := c.Param("id")

	database.Db.Delete(&p, id)

	c.String(http.StatusNoContent, "")
}

func EditarPersonalidade(c *gin.Context) {

	var personalidade models.Personalidade

	id := c.Param("id")

	err := c.ShouldBindJSON(&personalidade)
	if err != nil {
		panic(err.Error())
	}

	var p models.Personalidade

	database.Db.First(&p, id)

	p.Nome = personalidade.Nome
	p.Historia = personalidade.Historia

	database.Db.Save(&p)

	c.JSON(http.StatusOK, p)
}
