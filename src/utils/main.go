package main

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

// Combine and minify all css files
func main() {
	// List of CSS files to concatenate and minify
	files := []string{"../CSS/tokens.css", "../CSS/utils.css", "../CSS/ds.css"}
	 minifiedFilePath  := "../CSS/fullds.min.css"

	// Concatenate CSS files
	var concatenatedCSS strings.Builder
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		concatenatedCSS.Write(data)
		concatenatedCSS.WriteString("\n") // Ensure there's a newline between files
	}

	// Create a minifier instance
	m := minify.New()
	m.AddFunc("text/css", css.Minify)

	// Minify the concatenated CSS
	minifiedCSS := new(bytes.Buffer)
	if err := m.Minify("text/css", minifiedCSS, strings.NewReader(concatenatedCSS.String())); err != nil {
		log.Fatal(err)
	}

	// Write the minified CSS to a file
	if err := os.WriteFile(minifiedFilePath, minifiedCSS.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

	log.Println("Minification complete. Output saved to ", minifiedFilePath)

	if err := CalculateFileSize(minifiedFilePath); err != nil {
		log.Fatal(err)
	}
}


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
	log.Printf("The size of '%s' is %.2f KB", filePath, sizeInKB)

	return nil
}