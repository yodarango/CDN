package utils

import (
	"fmt"
	"os"
)

// output the size of the file path
func CalculateFileSize(filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file statistics
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// File size in bytes
	sizeInBytes := fileInfo.Size()

	// Convert size to kilobytes (1 KB = 1024 Bytes)
	sizeInKB := float64(sizeInBytes) / 1024.0

	// Log the size
	fmt.Printf("🐘 The size of '%s' is %.2f KB \n \n", filePath, sizeInKB)

	return nil
}