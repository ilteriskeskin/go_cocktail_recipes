package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oleiade/reflections"
)

type Cocktail struct {
	Drinks []struct {
		StrDrink        string `json:"strDrink"`
		StrCategory     string `json:"strCategory"`
		StrGlass        string `json:"strGlass"`
		StrInstructions string `json:"strInstructions"`
		StrIngredient1  string `json:"strIngredient1"`
		StrIngredient2  string `json:"strIngredient2"`
		StrIngredient3  string `json:"strIngredient3"`
		StrIngredient4  string `json:"strIngredient4"`
		StrIngredient5  string `json:"strIngredient5"`
		StrIngredient6  string `json:"strIngredient6"`
		StrIngredient7  string `json:"strIngredient7"`
		StrIngredient8  string `json:"strIngredient8"`
		StrIngredient9  string `json:"strIngredient9"`
		StrIngredient10 string `json:"strIngredient10"`
		StrIngredient11 string `json:"strIngredient11"`
		StrIngredient12 string `json:"strIngredient12"`
		StrIngredient13 string `json:"strIngredient13"`
		StrIngredient14 string `json:"strIngredient14"`
		StrIngredient15 string `json:"strIngredient15"`
	} `json:"drinks"`
}

func main() {

	var cocktailName string
	var cocktailIngredientsName []string
	var cocktailIngredients []string
	var cocktail Cocktail
	var i int
	defaultIngredient := "StrIngredient"

	for i = 1; i <= 15; i++ {
		cocktailIngredientsName = append(cocktailIngredientsName, defaultIngredient+fmt.Sprintf("%d", i))
	}

	fmt.Print("Enter cocktail name: ")
	fmt.Scanln(&cocktailName)
	fmt.Println("---------------------")

	url := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + cocktailName

	r, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &cocktail); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for drinkIndex, drink := range cocktail.Drinks {
		cocktailIngredients = []string{}

		for _, fieldName := range cocktailIngredientsName {

			value, _ := reflections.GetField(drink, fieldName)
			if value == "" {
				break
			}
			strValue := fmt.Sprintf("%v", value)
			cocktailIngredients = append(cocktailIngredients, strValue)
		}
		fmt.Println("ID: ", drinkIndex+1)
		fmt.Println("Name: ", drink.StrDrink)
		fmt.Println("Category: ", drink.StrCategory)
		fmt.Println("Glass: ", drink.StrGlass)
		fmt.Println("Instructions: ", drink.StrInstructions)
		for ingredientIndex, ingredient := range cocktailIngredients {
			fmt.Println("Ingredient", ingredientIndex+1, ": ", ingredient)
		}
		fmt.Println("---------------------")
	}
}
