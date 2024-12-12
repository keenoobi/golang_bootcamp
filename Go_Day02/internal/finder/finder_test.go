package finder

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestWalkDir(t *testing.T) {
	// Создаем временную директорию для тестов
	testDir := "testdata"
	if err := os.Mkdir(testDir, 0755); err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Создаем тестовые файлы и директории
	createTestFile := func(name string) {
		file, err := os.Create(name)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", name, err)
		}
		file.Close()
	}

	createTestDir := func(name string) {
		if err := os.Mkdir(name, 0755); err != nil {
			t.Fatalf("Failed to create test directory %s: %v", name, err)
		}
	}

	createTestSymlink := func(target, link string) {
		if err := os.Symlink(target, link); err != nil {
			t.Fatalf("Failed to create symlink %s -> %s: %v", link, target, err)
		}
	}

	createTestFile(filepath.Join(testDir, "file1.txt"))
	createTestFile(filepath.Join(testDir, "file2.log"))
	createTestDir(filepath.Join(testDir, "subdir"))
	createTestFile(filepath.Join(testDir, "subdir", "file3.txt"))
	createTestSymlink(filepath.Join(testDir, "file1.txt"), filepath.Join(testDir, "symlink1"))
	createTestSymlink("nonexistent", filepath.Join(testDir, "broken_symlink")) // Создаем ссылку на несуществующий файл

	// Тест: Поиск файлов с расширением .txt
	t.Run("WalkDir_FindFilesWithExtension", func(t *testing.T) {
		var output bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := WalkDir(testDir, false, true, false, "txt")
		if err != nil {
			t.Fatalf("Error walking directory: %v", err)
		}

		w.Close()
		os.Stdout = oldStdout
		output.ReadFrom(r)

		expected := filepath.Join(testDir, "file1.txt") + "\n" +
			filepath.Join(testDir, "subdir", "file3.txt") + "\n"
		if output.String() != expected {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expected, output.String())
		}
	})

	// Тест: Поиск директорий
	t.Run("WalkDir_FindDirectories", func(t *testing.T) {
		var output bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := WalkDir(testDir, true, false, false, "")
		if err != nil {
			t.Fatalf("Error walking directory: %v", err)
		}

		w.Close()
		os.Stdout = oldStdout
		output.ReadFrom(r)

		expected := testDir + "\n" +
			filepath.Join(testDir, "subdir") + "\n"
		if output.String() != expected {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expected, output.String())
		}
	})

	// Тест: Поиск символических ссылок
	t.Run("WalkDir_FindSymlinks", func(t *testing.T) {
		var output bytes.Buffer
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		err := WalkDir(testDir, false, false, true, "")
		if err != nil {
			t.Fatalf("Error walking directory: %v", err)
		}

		w.Close()
		os.Stdout = oldStdout
		output.ReadFrom(r)

		expected := filepath.Join(testDir, "broken_symlink") + " -> [broken]\n" + filepath.Join(testDir, "symlink1") + " -> " + filepath.Join(testDir, "file1.txt") + "\n"
		if output.String() != expected {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expected, output.String())
		}
	})

	// Тест: Обработка ошибок доступа
	t.Run("WalkDir_PermissionError", func(t *testing.T) {
		// Создаем директорию с ограниченными правами
		restrictedDir := filepath.Join(testDir, "restricted")
		createTestDir(restrictedDir)
		os.Chmod(restrictedDir, 0000) // Устанавливаем права на доступ

		err := WalkDir(testDir, true, true, true, "")
		if err != nil {
			t.Fatalf("Error walking directory: %v", err)
		}

		// Восстанавливаем права для удаления директории
		os.Chmod(restrictedDir, 0755)
	})
}
