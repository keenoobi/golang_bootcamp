package test

import (
	"Go_Day01/internal/fscomparator"
	"os"
	"testing"
)

func TestCompareFiles(t *testing.T) {
	// Тестовые данные
	oldFilename := "./../testdata/snapshot1.txt"
	newFilename := "./../testdata/snapshot2.txt"

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
	err = fscomparator.CompareFiles(oldFilename, newFilename)
	if err != nil {
		t.Errorf("Error comparing files: %v", err)
	}

	// Проверка результата
	outputFile.Close()
	outputContent, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	expectedOutput := "ADDED /etc/systemd/system/very_important/stash_location.jpg\n"
	if string(outputContent) != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, string(outputContent))
	}
}
