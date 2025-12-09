package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/sashabaranov/go-openai"
)

// GetProjectRoot returns the project root directory
func GetProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	// Go up from app/utils/ to project root
	return filepath.Join(filepath.Dir(b), "..", "..")
}

// LogToFile writes content to a log file in the logs directory
func LogToFile(filename string, content []byte) error {
	logsDir := filepath.Join(GetProjectRoot(), "logs")
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return err
	}
	filePath := filepath.Join(logsDir, filename)
	return os.WriteFile(filePath, content, 0644)
}

// ReadFile reads a file from the project root
func ReadFile(relativePath string) ([]byte, error) {
	filePath := filepath.Join(GetProjectRoot(), relativePath)
	return os.ReadFile(filePath)
}

// WriteFile writes content to a file in the project
func WriteFile(relativePath string, content []byte) error {
	filePath := filepath.Join(GetProjectRoot(), relativePath)
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filePath, content, 0644)
}

// EmbeddingLog represents the log structure for embedding requests
type EmbeddingLog struct {
	Timestamp string                   `json:"timestamp"`
	Model     string                   `json:"model"`
	Input     string                   `json:"input"`
	Response  openai.EmbeddingResponse `json:"response"`
}

func LogEmbedding(model openai.EmbeddingModel, input string, response openai.EmbeddingResponse) error {
	// Log to file with timestamp, input and model
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	embeddingLog := EmbeddingLog{
		Timestamp: timestamp,
		Model:     string(model),
		Input:     input,
		Response:  response,
	}

	logJSON, err := json.MarshalIndent(embeddingLog, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal embedding log: %w", err)
	}

	filename := fmt.Sprintf("embedding_%s.json", timestamp)
	if err := LogToFile(filename, logJSON); err != nil {
		return fmt.Errorf("failed to write log file: %w", err)
	}
	return nil
}
