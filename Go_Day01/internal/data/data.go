package data

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Ingredient struct {
	ItemName  string  `json:"ingredient_name" xml:"itemname"`
	ItemCount string  `json:"ingredient_count" xml:"itemcount"`
	ItemUnit  *string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	StoveTime   string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Cakes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

// Функция для вывода данных в другом формате
func PrintData(extension string, recipes Cakes) {
	switch extension {
	case ".xml":
		xml, err := сonvertDataToJSON(recipes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(xml))
	case ".json":
		json, err := сonvertDataToXML(recipes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(json))
	}
}

// функция для конвертации и вывода данных в формате XML
func сonvertDataToXML(recipes Cakes) ([]byte, error) {
	xmlData, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("cant convert to XML: %w", err)
	}
	return xmlData, nil
}

// функция для конвертации и вывода данных в формате JSON
func сonvertDataToJSON(recipes Cakes) ([]byte, error) {
	jsonData, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("cant convert to JSON: %w", err)
	}
	return jsonData, nil
}
