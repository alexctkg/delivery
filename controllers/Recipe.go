package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"

	service "delivery/services"
)

func IndexRecipe(c *gin.Context) {

	ingredientes := c.DefaultQuery("i", "")
	if ingredientes == "" {
		c.JSON(400, gin.H{"errors": []string{"A solicitação deve conter um ingrediente"}})
		c.Abort()
		return
	}

	//A API deve receber como parâmetro um conjunto de ingredientes (máximo 3)
	ingredientesSlice := strings.Split(ingredientes, ",")
	if len(ingredientesSlice) > 3 {
		c.JSON(400, gin.H{"errors": []string{"A solicitação deve conter um número máximo de 3 ingredientes"}})
		c.Abort()
		return
	}

	code, response, err := service.RecipePuppyRequest(ingredientes)
	if err != nil {
		c.JSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	c.JSON(*code, response)
	return
}
