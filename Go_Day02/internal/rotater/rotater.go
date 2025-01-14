package rotater

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ArchiveLog(logFile, archiveDir string) error {
	fileInfo, err := os.Stat(logFile)
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	// Проверка, что архивируется имено .log-файл
	if filepath.Ext(logFile) != ".log" {
		return fmt.Errorf("error: file %s is not a .log file", logFile)
	}

	baseName := filepath.Base(strings.TrimSuffix(logFile, filepath.Ext(logFile)))
	timestamp := fileInfo.ModTime().Unix()

	archiveName := fmt.Sprintf("%s_%d.tar.gz", baseName, timestamp)

	if archiveDir == "" {
		archiveDir = filepath.Dir(logFile)
	}

	archiveName = filepath.Join(archiveDir, archiveName)

	archiveFile, err := os.Create(archiveName)
	if err != nil {
		return fmt.Errorf("failed to create archive file: %v", err)
	}
	defer archiveFile.Close()

	gzipWriter := gzip.NewWriter(archiveFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	logFileHandle, err := os.Open(logFile)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer logFileHandle.Close()

	err = addFileToArchive(tarWriter, logFileHandle, fileInfo)
	if err != nil {
		return fmt.Errorf("failed to add file to archive: %v", err)
	}

	return nil
}

func addFileToArchive(tarWriter *tar.Writer, file *os.File, fileInfo os.FileInfo) error {
	header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
	if err != nil {
		return fmt.Errorf("failed to create tar header: %v", err)
	}

	header.Name = fileInfo.Name()

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("failed to write tar header: %v", err)
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return fmt.Errorf("failed to copy file content to archive: %v", err)
	}

	return nil
}
