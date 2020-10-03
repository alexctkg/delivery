package controllers

import (
	"delivery/models"
	"strings"

	"github.com/gin-gonic/gin"

	service "delivery/services"
)

// IndexRecipe godoc
// @Tags Recipe
// @Summary Index Recipe
// @Description Gera uma lista de receitas dado ingredientews
// @Param i query string false "Array de ngredientes (separados por virsgula)"
// @Produce json
// @Success 200 {object} []models.Recipe
// @Failure 400 {object} models.DefaultError
// @Router /recipe/ [GET]
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
		c.JSON(400, gin.H{"errors": []string{"Ops! Algo de errado aconteceu."}})
		c.Abort()
		return
	}

	if code != nil {
		if *code != 200 {
			c.JSON(400, gin.H{"errors": []string{"Ops! Algo de errado aconteceu."}})
			c.Abort()
			return
		}
	} else {
		if *code != 200 {
			c.JSON(400, gin.H{"errors": []string{"Ops! Algo de errado aconteceu. 1"}})
			c.Abort()
			return
		}
	}

	//convertendo map para estrutura para facilitar manipulação
	recipePuppyStruct, err := models.MapToRecipePuppy(response)
	if err != nil {
		c.JSON(400, gin.H{"errors": []string{"Houve uma falha na operação."}})
		c.Abort()
		return
	}

	recipe := models.Recipe{}
	recipe.Keywords = ingredientesSlice

	for _, worker := range recipePuppyStruct.Results {
		//Para obter o gif no Giphy, utilize o título da receita recebido pelo RecipePuppy;
		// code, response, err := service.GiphyRequest(worker.Title)
		// if err != nil {
		// 	c.JSON(400, gin.H{"errors": []string{"Ops! Algo de errado aconteceu."}})
		// 	c.Abort()
		// 	return
		// }

		// spew.Dump(response)

		// if code != nil {
		// 	if *code != 200 {
		// 		c.JSON(400, gin.H{"errors": []string{"Ops! Algo de errado aconteceu."}})
		// 		c.Abort()
		// 		return
		// 	}
		// } else {
		// 	if *code != 200 {
		// 		c.JSON(400, gin.H{"errors": []string{"Ops! Algo de errado aconteceu."}})
		// 		c.Abort()
		// 		return
		// 	}
		// }

		recipeResouce := models.RecipePuppyResultsToRecipe(worker)
		recipe.Recipes = append(recipe.Recipes, recipeResouce)

	}
	c.JSON(200, recipe)
	c.Abort()
	return
}
