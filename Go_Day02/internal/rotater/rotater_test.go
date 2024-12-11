package rotater

import (
	"archive/tar"
	"compress/gzip"
	"os"
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

	// Тест: Успешная архивация файла
	t.Run("ArchiveLog_Success", func(t *testing.T) {
		err := ArchiveLog(logFile, "")
		if err != nil {
			t.Fatalf("Error archiving log file: %v", err)
		}

		// Проверяем, что архив создан
		archiveName := "test.log_1600785299.tar.gz"
		if _, err := os.Stat(archiveName); os.IsNotExist(err) {
			t.Fatalf("Archive file %s was not created", archiveName)
		}

		// Удаляем архив
		os.Remove(archiveName)
	})

	// Тест: Файл не существует
	t.Run("ArchiveLog_FileNotFound", func(t *testing.T) {
		err := ArchiveLog("nonexistent.log", "")
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})

	// Тест: Некорректная директория архивации
	t.Run("ArchiveLog_InvalidArchiveDir", func(t *testing.T) {
		err := ArchiveLog(logFile, "/nonexistent/directory")
		if err == nil {
			t.Fatalf("Expected error, got nil")
		}
	})

	// Удаляем тестовые файлы
	os.Remove(logFile)
	os.RemoveAll("testdata")
}

func TestAddFileToArchive(t *testing.T) {
	// Создаем временный лог-файл
	logFile := "testdata/test.log"
	os.Mkdir("testdata", 0755)
	os.WriteFile(logFile, []byte("Test log content"), 0644)

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

	// Удаляем тестовые файлы
	os.Remove(logFile)
	os.RemoveAll("testdata")
}
