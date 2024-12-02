package parser

import (
	"Go_Day01/internal/data"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type DBReader interface {
	Read(filename string) (data.Cakes, error)
}

type JSONReader struct{}

func (j JSONReader) Read(filename string) (data.Cakes, error) {
	var cakes data.Cakes

	bytesValue, err := os.ReadFile(filename)
	if err != nil {
		return cakes, fmt.Errorf("error: cant open or read file: %w", err)
	}

	err = json.Unmarshal(bytesValue, &cakes)
	if err != nil {
		return cakes, fmt.Errorf("error: cant parse JSON: %w", err)
	}

	return cakes, nil
}

type XMLReader struct{}

func (x XMLReader) Read(filename string) (data.Cakes, error) {
	var cakes data.Cakes

	bytesValue, err := os.ReadFile(filename)
	if err != nil {
		return cakes, fmt.Errorf("error: cant open or read file: %w", err)
	}

	err = xml.Unmarshal(bytesValue, &cakes)
	if err != nil {
		return cakes, fmt.Errorf("error: cant parse XML: %w", err)
	}

	return cakes, nil
}

func GetReader(filename string) (DBReader, string, error) {
	extension := filepath.Ext(filename)

	switch extension {
	case ".json":
		return JSONReader{}, extension, nil
	case ".xml":
		return XMLReader{}, extension, nil
	default:
		return nil, "", fmt.Errorf("unsupported file extension or file doesnt exis")
	}
}
