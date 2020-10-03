package models

import (
	"sort"
	"strings"
)

// Recipe struct para retorno ao endpoint de recipe
type Recipe struct {
	Keywords []string  `json:"keywords"`
	Recipes  []Recipes `json:"recipes"`
}

// Recipes .
type Recipes struct {
	Title       string   `json:"title"`
	Ingredients []string `json:"ingredients"`
	Link        string   `json:"link"`
	Giphy       *string  `json:"gif"`
}

//RecipePuppyResultsToRecipe struct api puppy.Results para a recipe.Recipes
func RecipePuppyResultsToRecipe(puppy RecipePuppyResults) Recipes {
	recipes := Recipes{}

	recipes.Title = puppy.Title
	// Os ingredientes recebidos pelo RecipePuppy são recebidos em String. Organize os ingredientes em um array e ordene esse array por ordem alfabética;
	ingredientesSlice := strings.Split(puppy.Ingredients, ",")
	sort.Strings(ingredientesSlice)
	recipes.Ingredients = ingredientesSlice
	recipes.Link = puppy.Href
	recipes.Giphy = puppy.Giphy

	return recipes
}
