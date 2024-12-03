package test

import (
	"Go_Day01/internal/data"
	"Go_Day01/internal/parser"
	"os"
	"path/filepath"
	"testing"
)

func TestPrintData(t *testing.T) {
	// Путь к тестовому JSON файлу
	jsonFilePath := filepath.Join("../testdata", "stolen_database.json")

	// Читаем данные из файла
	jsonReader := parser.JSONReader{}
	recipes, err := jsonReader.Read(jsonFilePath)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	// Создаем временный файл для записи результатов
	outputFile, err := os.CreateTemp("", "print_output_*.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(outputFile.Name())

	// Перенаправляем вывод в временный файл
	oldStdout := os.Stdout
	os.Stdout = outputFile
	defer func() { os.Stdout = oldStdout }()

	// Вызов функции для XML
	data.PrintData(".json", recipes)

	// Проверка результата
	outputFile.Close()
	outputContent, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	// Ожидаемый результат для XML
	expectedOutput := `<Cakes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>45 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>2</itemcount>
                <itemunit>mugs</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>8</itemcount>
            </item>
            <item>
                <itemname>Coffee Beans</itemname>
                <itemcount>2.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
            </item>
        </ingredients>
    </cake>
    <cake>
        <name>Moonshine Muffin</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
        </ingredients>
    </cake>
</Cakes>
`

	// Сравниваем результат с ожидаемым
	if string(outputContent) != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, string(outputContent))
	}
}

func TestPrintDataXML(t *testing.T) {
	// Путь к тестовому XML файлу
	xmlFilePath := filepath.Join("../testdata", "original_database.xml")

	// Читаем данные из файла
	xmlReader := parser.XMLReader{}
	recipes, err := xmlReader.Read(xmlFilePath)
	if err != nil {
		t.Fatalf("Failed to read XML file: %v", err)
	}

	// Создаем временный файл для записи результатов
	outputFile, err := os.CreateTemp("", "print_output_*.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(outputFile.Name())

	// Перенаправляем вывод в временный файл
	oldStdout := os.Stdout
	os.Stdout = outputFile
	defer func() { os.Stdout = oldStdout }()

	// Вызов функции для JSON
	data.PrintData(".xml", recipes)

	// Проверка результата
	outputFile.Close()
	outputContent, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	// Ожидаемый результат для JSON
	expectedOutput := `{
    "cake": [
        {
            "name": "Red Velvet Strawberry Cake",
            "time": "40 min",
            "ingredients": [
                {
                    "ingredient_name": "Flour",
                    "ingredient_count": "3",
                    "ingredient_unit": "cups"
                },
                {
                    "ingredient_name": "Vanilla extract",
                    "ingredient_count": "1.5",
                    "ingredient_unit": "tablespoons"
                },
                {
                    "ingredient_name": "Strawberries",
                    "ingredient_count": "7",
                    "ingredient_unit": ""
                },
                {
                    "ingredient_name": "Cinnamon",
                    "ingredient_count": "1",
                    "ingredient_unit": "pieces"
                }
            ]
        },
        {
            "name": "Blueberry Muffin Cake",
            "time": "30 min",
            "ingredients": [
                {
                    "ingredient_name": "Baking powder",
                    "ingredient_count": "3",
                    "ingredient_unit": "teaspoons"
                },
                {
                    "ingredient_name": "Brown sugar",
                    "ingredient_count": "0.5",
                    "ingredient_unit": "cup"
                },
                {
                    "ingredient_name": "Blueberries",
                    "ingredient_count": "1",
                    "ingredient_unit": "cup"
                }
            ]
        }
    ]
}
`

	// Сравниваем результат с ожидаемым
	if string(outputContent) != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, string(outputContent))
	}
}
