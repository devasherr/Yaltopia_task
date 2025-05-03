package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"yaltopia_task/models"
)

func Parse() error {
	basePath, err := os.Getwd()
	if err != nil {
		return err
	}
	rawPath := filepath.Join(basePath, "data", "raw")
	parsedPath := filepath.Join("data", "parsed")

	if err := os.MkdirAll(parsedPath, os.ModePerm); err != nil {
		return err
	}
	files := []struct {
		name   string
		parser func([]byte) (any, error)
	}{
		{
			name: "cricket_prematch.json",
			parser: func(data []byte) (any, error) {
				return ParseCricketPrematchJSON(data)
			},
		},
		{
			name: "cricket_result.json",
			parser: func(data []byte) (any, error) {
				return ParseCricketResultJSON(data)
			},
		},
		{
			name: "volleyball_prematch.json",
			parser: func(data []byte) (any, error) {
				return ParseVolleyballPrematchJSON(data)
			},
		},
		{
			name: "volleyball_result.json",
			parser: func(data []byte) (any, error) {
				return ParseVolleyballResultJSON(data)
			},
		},
	}

	for _, file := range files {
		rawData, err := os.ReadFile(filepath.Join(rawPath, file.name))
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", file.name, err)
		}

		parsedData, err := file.parser(rawData)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", file.name, err)
		}

		if err := writeJSON(filepath.Join(parsedPath, file.name), parsedData); err != nil {
			return fmt.Errorf("failed to write %s: %w", file.name, err)
		}
	}

	return nil
}

func writeJSON(filePath string, data any) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(bytes)
	return err
}

func ParseCricketPrematchJSON(jsonData []byte) (*models.PrematchData, error) {
	var data models.PrematchData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing prematch JSON: %v", err)
	}
	return &data, nil
}

func ParseCricketResultJSON(jsonData []byte) (*models.ResultData, error) {
	var data models.ResultData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing result JSON: %v", err)
	}
	return &data, nil
}

func ParseVolleyballPrematchJSON(jsonData []byte) (*models.VolleyPrematchData, error) {
	var data models.VolleyPrematchData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing prematch JSON: %v", err)
	}
	return &data, nil
}

func ParseVolleyballResultJSON(jsonData []byte) (*models.VolleyResultData, error) {
	var data models.VolleyResultData
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing result JSON: %v", err)
	}
	return &data, nil
}
