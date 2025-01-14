package rotater

import (
	"archive/tar"
	"compress/gzip"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestArchiveLog(t *testing.T) {
	// Создаем временный лог-файл
	logFile := "testdata/test.log"
	os.Mkdir("testdata", 0755)
	os.WriteFile(logFile, []byte("Test log content"), 0644)

	// Устанавливаем MTIME для тестового файла
	os.Chtimes(logFile, time.Now(), time.Unix(1600785299, 0))

	// Автоматически удаляем тестовые файлы после завершения тестов
	t.Cleanup(func() {
		os.Remove(logFile)
		os.RemoveAll("testdata")
	})

	// Тест: Успешная архивация файла
	t.Run("ArchiveLog_Success_SameDirectory", func(t *testing.T) {
		err := ArchiveLog(logFile, "")
		if err != nil {
			t.Fatalf("Error archiving log file: %v", err)
		}

		// Проверяем, что архив создан в той же директории
		archiveName := filepath.Join(filepath.Dir(logFile), "test_1600785299.tar.gz")
		if _, err := os.Stat(archiveName); os.IsNotExist(err) {
			t.Fatalf("Archive file %s was not created", archiveName)
		}

		// Удаляем архив
		os.Remove(archiveName)
	})

	// Тест: Успешная архивация файла в указанной директории
	t.Run("ArchiveLog_Success_SpecifiedDirectory", func(t *testing.T) {
		// Создаём временную директорию для архива
		archiveDir := "testdata/archive"
		os.MkdirAll(archiveDir, 0755)
		defer os.RemoveAll(archiveDir)

		err := ArchiveLog(logFile, archiveDir)
		if err != nil {
			t.Fatalf("Error archiving log file: %v", err)
		}

		// Проверяем, что архив создан в указанной директории
		archiveName := filepath.Join(archiveDir, "test_1600785299.tar.gz")
		if _, err := os.Stat(archiveName); os.IsNotExist(err) {
			t.Fatalf("Archive file %s was not created", archiveName)
		}
	})

	// Тест: Ошибка, если директория не существует
	t.Run("ArchiveLog_DirectoryNotExist", func(t *testing.T) {
		archiveDir := "testdata/nonexistent"
		err := ArchiveLog(logFile, archiveDir)
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}

		expectedError := "failed to create archive file"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error to contain '%s', got '%v'", expectedError, err)
		}
	})

	// Тест: Ошибка, если файл не существует
	t.Run("ArchiveLog_FileNotExist", func(t *testing.T) {
		err := ArchiveLog("nonexistent.log", "")
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}

		expectedError := "failed to get file info"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error to contain '%s', got '%v'", expectedError, err)
		}
	})

	// Тест: Ошибка, если нет прав на запись в директорию
	t.Run("ArchiveLog_NoWritePermission", func(t *testing.T) {
		// Создаём директорию без прав на запись
		archiveDir := "testdata/readonly"
		os.MkdirAll(archiveDir, 0444) // Только чтение
		defer os.RemoveAll(archiveDir)

		err := ArchiveLog(logFile, archiveDir)
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}

		expectedError := "failed to create archive file"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error to contain '%s', got '%v'", expectedError, err)
		}
	})
}

func TestAddFileToArchive(t *testing.T) {
	// Создаем временный лог-файл
	logFile := "testdata/test.log"
	os.Mkdir("testdata", 0755)
	os.WriteFile(logFile, []byte("Test log content"), 0644)

	// Автоматически удаляем тестовые файлы после завершения тестов
	t.Cleanup(func() {
		os.Remove(logFile)
		os.RemoveAll("testdata")
	})

	// Открываем лог-файл
	file, err := os.Open(logFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	// Получаем информацию о файле
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}

	// Тест: Успешная запись файла в архив
	t.Run("AddFileToArchive_Success", func(t *testing.T) {
		// Создаем временный файл архива
		archiveFile, err := os.Create("test.tar.gz")
		if err != nil {
			t.Fatalf("Failed to create archive file: %v", err)
		}
		defer archiveFile.Close()

		// Создаем gzip writer
		gzipWriter := gzip.NewWriter(archiveFile)
		defer gzipWriter.Close()

		// Создаем tar writer
		tarWriter := tar.NewWriter(gzipWriter)
		defer tarWriter.Close()

		// Добавляем файл в архив
		err = addFileToArchive(tarWriter, file, fileInfo)
		if err != nil {
			t.Fatalf("Error adding file to archive: %v", err)
		}

		// Удаляем архив
		os.Remove("test.tar.gz")
	})
}
