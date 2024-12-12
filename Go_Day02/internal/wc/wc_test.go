package wc

import (
	"os"
	"testing"
)

func TestCountAll(t *testing.T) {
	// Создаем директорию testdata, если она не существует
	testDir := "testdata"
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		err := os.Mkdir(testDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}
	}

	// Удаляем тестовые файлы и директорию после завершения тестов
	defer os.RemoveAll(testDir)

	// Создаем тестовые файлы
	createTestFile := func(name, content string) {
		err := os.WriteFile(name, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", name, err)
		}
	}

	// Тест: Успешная обработка файла
	t.Run("CountAll_Success", func(t *testing.T) {
		filePath := "testdata/test1.txt"
		createTestFile(filePath, "Hello world\nThis is a test\nFile for wc\n")
		lineCount, charsCount, wordsCount, err := CountAll(filePath)
		if err != nil {
			t.Fatalf("Error counting in file: %v", err)
		}

		// Проверяем результаты
		if lineCount != 3 {
			t.Errorf("Expected 3 lines, got %d", lineCount)
		}
		if charsCount != 36 {
			t.Errorf("Expected 36 characters, got %d", charsCount)
		}
		if wordsCount != 9 {
			t.Errorf("Expected 9 words, got %d", wordsCount)
		}
	})

	// Тест: Обработка пустого файла
	t.Run("CountAll_EmptyFile", func(t *testing.T) {
		filePath := "testdata/empty.txt"
		createTestFile(filePath, "")
		lineCount, charsCount, wordsCount, err := CountAll(filePath)
		if err != nil {
			t.Fatalf("Error counting in empty file: %v", err)
		}

		// Проверяем результаты
		if lineCount != 0 {
			t.Errorf("Expected 0 lines, got %d", lineCount)
		}
		if charsCount != 0 {
			t.Errorf("Expected 0 characters, got %d", charsCount)
		}
		if wordsCount != 0 {
			t.Errorf("Expected 0 words, got %d", wordsCount)
		}
	})

	// Тест: Обработка несуществующего файла
	t.Run("CountAll_FileNotFound", func(t *testing.T) {
		filePath := "testdata/nonexistent.txt"
		_, _, _, err := CountAll(filePath)
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})

	// Тест: Обработка файла с одной строкой
	t.Run("CountAll_SingleLine", func(t *testing.T) {
		filePath := "testdata/test2.txt"
		createTestFile(filePath, "Single line file")
		lineCount, charsCount, wordsCount, err := CountAll(filePath)
		if err != nil {
			t.Fatalf("Error counting in single line file: %v", err)
		}

		// Проверяем результаты
		if lineCount != 1 {
			t.Errorf("Expected 1 line, got %d", lineCount)
		}
		if charsCount != 16 {
			t.Errorf("Expected 16 characters, got %d", charsCount)
		}
		if wordsCount != 3 {
			t.Errorf("Expected 3 words, got %d", wordsCount)
		}
	})
}
