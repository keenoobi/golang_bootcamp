package test

import (
	"Go_Day01/internal/dbcomparator"
	"Go_Day01/internal/parser"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func TestCompareRecipes(t *testing.T) {
	// Путь к тестовым файлам
	oldFilePath := filepath.Join("../testdata", "original_database.xml")
	newFilePath := filepath.Join("../testdata", "stolen_database.json")

	// Читаем данные из файлов
	xmlReader := parser.XMLReader{}
	oldRecipes, err := xmlReader.Read(oldFilePath)
	if err != nil {
		t.Fatalf("Failed to read old file: %v", err)
	}

	jsonReader := parser.JSONReader{}
	newRecipes, err := jsonReader.Read(newFilePath)
	if err != nil {
		t.Fatalf("Failed to read new file: %v", err)
	}

	// Создаем временный файл для записи результатов
	outputFile, err := os.CreateTemp("", "compare_output_*.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(outputFile.Name())

	// Перенаправляем вывод в временный файл
	oldStdout := os.Stdout
	os.Stdout = outputFile
	defer func() { os.Stdout = oldStdout }()

	// Вызов функции
	dbcomparator.CompareRecipes(oldRecipes, newRecipes)

	// Проверка результата
	outputFile.Close()
	outputContent, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	// Разделяем вывод на строки и сортируем их
	outputLines := strings.Split(string(outputContent), "\n")
	sort.Strings(outputLines)

	// Ожидаемый вывод
	expectedOutput := []string{
		"ADDED cake \"Moonshine Muffin\"",
		"ADDED ingredient \"Coffee Beans\" for cake \"Red Velvet Strawberry Cake\"",
		"CHANGED cooking time for cake \"Red Velvet Strawberry Cake\" - \"45 min\" instead of \"40 min\"",
		"CHANGED unit count for ingredient \"Flour\" for cake \"Red Velvet Strawberry Cake\" - \"2\" instead of \"3\"",
		"CHANGED unit count for ingredient \"Strawberries\" for cake \"Red Velvet Strawberry Cake\" - \"8\" instead of \"7\"",
		"CHANGED unit for ingredient \"Flour\" for cake \"Red Velvet Strawberry Cake\" - \"mugs\" instead of \"cups\"",
		"REMOVED cake \"Blueberry Muffin Cake\"",
		"REMOVED ingredient \"Vanilla extract\" for cake \"Red Velvet Strawberry Cake\"",
		"REMOVED unit \"pieces\" for ingredient \"Cinnamon\" for cake \"Red Velvet Strawberry Cake\"",
		"", // Добавляем пустую строку для соответствия количеству строк
	}
	sort.Strings(expectedOutput)

	// Сравниваем отсортированные строки
	if len(outputLines) != len(expectedOutput) {
		t.Errorf("Expected %d lines, got %d", len(expectedOutput), len(outputLines))
	}

	for i := range outputLines {
		if outputLines[i] != expectedOutput[i] {
			t.Errorf("Expected line:\n%s\nGot:\n%s", expectedOutput[i], outputLines[i])
		}
	}
}
