package data

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
