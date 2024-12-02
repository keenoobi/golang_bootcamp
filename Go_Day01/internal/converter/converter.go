package converter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"Go_Day01/internal/data"
)

// функция для конвертации и вывода данных в формате XML
func PrintXML(recipes data.Cakes) {
	xmlData, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Println("error: cant convert to XML:", err)
		return
	}
	fmt.Println(string(xmlData))
}

// функция для конвертации и вывода данных в формате JSON
func PrintJSON(recipes data.Cakes) {
	jsonData, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Println("error: cant convert to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
