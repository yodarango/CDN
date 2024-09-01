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

// Combine and minify all css files
func BundleAndMinifyAllCSS(distPath string) {
	fmt.Println("💨 Starting minification process")
	// List of CSS files to concatenate and minify
	var files []string

	// get all the css files 
	cssDir, err := os.Open("./src/CSS")

	if err != nil {
		log.Fatalf("Error opening CSS directory %v", err)
	}

	fmt.Println("📂 Reading all contents of CSS dir")
	cssFiles, err :=  cssDir.Readdirnames(0)
	if err != nil {
		log.Fatalf("Error reading CSS directory %v", err)
	}

	fmt.Println("📖 Reading contents for each file")
	for _, cssFile := range cssFiles {

		fileName := "./src/CSS/" + cssFile
		
		if strings.Contains(cssFile, ".css") {
			files = append(files, fileName)
		}
		
	}

	 minifiedFilePath  := distPath + "/fullds.min.css"
	fmt.Println("🤏 Minified all CSS files")

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

	fmt.Println("📝 Now writing buffer to file")
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

	fmt.Println("📝 Wrote buffer to one file file")
	fmt.Println("✅ Minification complete. Output saved to ", minifiedFilePath)

	if err := CalculateFileSize(minifiedFilePath); err != nil {
		log.Fatal(err)
	}
}