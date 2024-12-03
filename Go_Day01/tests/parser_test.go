package test

import (
	"Go_Day01/internal/parser"
	"path/filepath"
	"testing"
)

func TestJSONReader_Read(t *testing.T) {
	// Путь к тестовому JSON файлу
	jsonFilePath := filepath.Join("../testdata", "stolen_database.json")

	// Создаем экземпляр JSONReader
	jsonReader := parser.JSONReader{}

	// Читаем файл
	cakes, err := jsonReader.Read(jsonFilePath)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	// Проверяем, что данные корректно прочитаны
	if len(cakes.Cake) == 0 {
		t.Errorf("Expected non-empty cakes list, got empty")
	}

	// Пример проверки конкретного элемента
	if cakes.Cake[0].Name != "Red Velvet Strawberry Cake" {
		t.Errorf("Expected cake name 'Red Velvet Strawberry Cake', got %s", cakes.Cake[0].Name)
	}
}

func TestXMLReader_Read(t *testing.T) {
	// Путь к тестовому XML файлу
	xmlFilePath := filepath.Join("../testdata", "original_database.xml")

	// Создаем экземпляр XMLReader
	xmlReader := parser.XMLReader{}

	// Читаем файл
	cakes, err := xmlReader.Read(xmlFilePath)
	if err != nil {
		t.Fatalf("Failed to read XML file: %v", err)
	}

	// Проверяем, что данные корректно прочитаны
	if len(cakes.Cake) == 0 {
		t.Errorf("Expected non-empty cakes list, got empty")
	}

	// Пример проверки конкретного элемента
	if cakes.Cake[0].Name != "Red Velvet Strawberry Cake" {
		t.Errorf("Expected cake name 'Red Velvet Strawberry Cake', got %s", cakes.Cake[0].Name)
	}
}

func TestGetReader(t *testing.T) {
	// Тестовые данные
	testCases := []struct {
		filename string
		expected parser.DBReader
		hasError bool
	}{
		{"../testdata/stolen_database.json", parser.JSONReader{}, false},
		{"../testdata/original_database.xml", parser.XMLReader{}, false},
		{"../testdata/unknown_extension.txt", nil, true},
	}

	for _, tc := range testCases {
		reader, _, err := parser.GetReader(tc.filename)

		if tc.hasError {
			if err == nil {
				t.Errorf("Expected error for file %s, got nil", tc.filename)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for file %s: %v", tc.filename, err)
			}

			if _, ok := reader.(parser.JSONReader); ok && (tc.expected != parser.JSONReader{}) {
				t.Errorf("Expected JSONReader for file %s, got %T", tc.filename, reader)
			}

			if _, ok := reader.(parser.XMLReader); ok && (tc.expected != parser.XMLReader{}) {
				t.Errorf("Expected XMLReader for file %s, got %T", tc.filename, reader)
			}
		}
	}
}
