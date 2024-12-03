package dbcomparator

import (
	"Go_Day01/internal/data"
	"fmt"
)

// функция для сравнения данных и вывода результатов
func CompareRecipes(oldRecipes, newREcipes data.Cakes) {
	oldCakesMap := make(map[string]data.Cake)
	newCakesMap := make(map[string]data.Cake)

	for _, cake := range oldRecipes.Cake {
		oldCakesMap[cake.Name] = cake
	}

	for _, cake := range newREcipes.Cake {
		newCakesMap[cake.Name] = cake
	}

	// сравнение рецептов
	for name, oldCake := range oldCakesMap {
		newCake, exists := newCakesMap[name]

		if !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
			continue
		}

		if oldCake.StoveTime != newCake.StoveTime {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, newCake.StoveTime, oldCake.StoveTime)
		}
		compareIngredients(name, oldCake.Ingredients, newCake.Ingredients)
	}
	for name := range newCakesMap {
		if _, exists := oldCakesMap[name]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", name)
		}
	}

}

func compareIngredients(cakeName string, oldIngredients, newIngredients []data.Ingredient) {
	oldIngredientsMap := make(map[string]data.Ingredient)
	newIngredientsMap := make(map[string]data.Ingredient)

	for _, ingredient := range oldIngredients {
		oldIngredientsMap[ingredient.ItemName] = ingredient
	}

	for _, ingredient := range newIngredients {
		newIngredientsMap[ingredient.ItemName] = ingredient
	}

	// сравниваем ингредиенты
	for name, oldIngredient := range oldIngredientsMap {
		newIngredient, exists := newIngredientsMap[name]
		if !exists {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", name, cakeName)
			continue
		}

		// сравниваем количество
		if oldIngredient.ItemCount != newIngredient.ItemCount {
			fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, cakeName, newIngredient.ItemCount, oldIngredient.ItemCount)
		}

		// сравниваем единицы измерения
		if oldIngredient.ItemUnit == nil && newIngredient.ItemUnit != nil {
			fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", *newIngredient.ItemUnit, name, cakeName)
		} else if oldIngredient.ItemUnit != nil && newIngredient.ItemUnit == nil && *oldIngredient.ItemUnit != "" {
			fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", *oldIngredient.ItemUnit, name, cakeName)
		} else if oldIngredient.ItemUnit != nil && newIngredient.ItemUnit != nil && *oldIngredient.ItemUnit != *newIngredient.ItemUnit {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, cakeName, *newIngredient.ItemUnit, *oldIngredient.ItemUnit)
		}
	}

	// проверяем добавлеы ли ингредиенты
	for name := range newIngredientsMap {
		if _, exists := oldIngredientsMap[name]; !exists {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", name, cakeName)
		}
	}
}
