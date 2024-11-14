package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

// Combine and minify all css files
func BundleAndMinifyAllJS(distPath string) {
	fmt.Println("💨 Starting JS minification process")
	file := "./src/JS/tokens.js"

	minifiedFilePath  := distPath + "/icons.min.js"
	fmt.Println("🤏 Minified all CSS files")

	// Concatenate CSS files
	var concatenatedJS strings.Builder
	// for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		concatenatedJS.Write(data)
		concatenatedJS.WriteString("\n") // Ensure there's a newline between files
	// }

	fmt.Println("📝 Now writing buffer to file")
	// Create a minifier instance
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	// Minify the concatenated CSS
	minifiedCSS := new(bytes.Buffer)
	if err := m.Minify("text/javascript", minifiedCSS, strings.NewReader(concatenatedJS.String())); err != nil {
		log.Fatal(err)
	}

	// Write the minified CSS to a file
	if err := os.WriteFile(minifiedFilePath, minifiedCSS.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("📝 Wrote buffer to one file file")
	fmt.Println("✅ Minification complete. Output saved to ", minifiedFilePath)

	if err := CalculateFileSize(minifiedFilePath); err != nil {
		log.Fatal(err)
	}
}