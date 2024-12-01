package data

import (
	"fmt"
	"strings"
)

type Ingredient struct {
	IngredientName  string  `json:"ingredient_name" xml:"itemname"`
	IngredientCount string  `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  *string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Cakes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

// Метод String для структуры Ingredient
func (i Ingredient) String() string {
	if i.IngredientUnit == nil {
		return fmt.Sprintf("%s %s", i.IngredientName, i.IngredientCount)
	}
	return fmt.Sprintf("%s %s %s", i.IngredientName, i.IngredientCount, *i.IngredientUnit)
}

// Метод String для структуры Cake
func (c Cake) String() string {
	ingredients := make([]string, len(c.Ingredients))
	for i, ingredient := range c.Ingredients {
		ingredients[i] = ingredient.String()
	}
	return fmt.Sprintf("Название: %s\nВремя: %s\nИнгредиенты:\n%s\n", c.Name, c.Time, strings.Join(ingredients, "\n"))
}

// Метод String для структуры Cakes
func (cs Cakes) String() string {
	cakes := make([]string, len(cs.Cake))
	for i, cake := range cs.Cake {
		cakes[i] = cake.String()
	}
	return strings.Join(cakes, "\n")
}
