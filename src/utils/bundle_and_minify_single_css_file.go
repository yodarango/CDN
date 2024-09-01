package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

// Combine and minify a specific css file by path
func BundleAndMinifySingleCSSFile(file string, distPath string){

	fmt.Println("💨 Starting minification process")
	originalFilePath := strings.Replace(file, ".css", "", -1)
	originalFilePath = strings.Replace(originalFilePath, "./src/CSS/", "", -1)
	minifiedFilePath  := distPath + "/" + originalFilePath + ".min.css"
	fmt.Println("❌ removed file ending")

		// Read the content of the original CSS file
	originalCSSContent, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading CSS file: %v", err)
	}

	// Create a minifier instance
	m := minify.New()
	m.AddFunc("text/css", css.Minify)

	// Minify the concatenated CSS
	fmt.Println("🤏 Starting minification process")
	minifiedCSS := new(bytes.Buffer)
	if err := m.Minify("text/css", minifiedCSS, bytes.NewReader(originalCSSContent)); err != nil {
		log.Fatalf("Error minifying resource: %v", err)
	}
	fmt.Println("✅ Ended minification process")
	fmt.Println("📝 Writing buffer to file")

	// Write the minified CSS to a file
	if err := os.WriteFile(minifiedFilePath, minifiedCSS.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Minification complete. Output saved to ", minifiedFilePath)

	if err := CalculateFileSize(minifiedFilePath); err != nil {
		log.Fatalf("Error calculating file size: %v", err)
	}

	CalculateFileSize(minifiedFilePath)
}