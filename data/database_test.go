package dataBase

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"

	"example/web-service-gin/models"

	"github.com/stretchr/testify/assert"
)

// Helper function to capture log output
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	f()
	return buf.String()
}

func TestLoadAlbums(t *testing.T) {
	// Create a temporary JSON file with test data
	testData := []models.Album{
		{ID: "1", Title: "Test Album 1", Artist: "Test Artist 1", Price: 10.99},
		{ID: "2", Title: "Test Album 2", Artist: "Test Artist 2", Price: 20.99},
	}
	jsonData, err := json.Marshal(testData)
	assert.NoError(t, err)

	tmpFile, err := os.CreateTemp("", "albums_test_*.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(jsonData)
	assert.NoError(t, err)
	tmpFile.Close()

	// Call LoadAlbums with the temporary file
	LoadAlbums(tmpFile.Name())

	// Verify that the Albums variable is populated correctly
	assert.Equal(t, testData, Albums)
}

func TestLoadAlbums_FileReadError(t *testing.T) {
	// Capture log output
	output := captureOutput(func() {
		LoadAlbums("non_existent_file.json")
	})

	// Verify log output
	assert.Contains(t, output, "Failed to read file")
}

func TestLoadAlbums_JSONUnmarshalError(t *testing.T) {
	// Create a temporary file with invalid JSON data
	tmpFile, err := os.CreateTemp("", "albums_test_*.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = io.WriteString(tmpFile, "invalid json")
	assert.NoError(t, err)
	tmpFile.Close()

	// Capture log output
	output := captureOutput(func() {
		LoadAlbums(tmpFile.Name())
	})

	// Verify log output
	assert.Contains(t, output, "Failed to unmarshal JSON")
}
