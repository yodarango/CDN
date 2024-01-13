package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)


func main(){
disPath := "../CSS/dist"
bundleAndMinifyAllCSS(disPath)
bundleAndMinifySingleCSSFile("../CSS/ds.css", disPath)
bundleAndMinifySingleCSSFile("../CSS/tokens.css", disPath)
bundleAndMinifySingleCSSFile("../CSS/utils.css", disPath)
}

// Combine and minify a specific css file by path
func bundleAndMinifySingleCSSFile(file string, distPath string){

	fmt.Println("💨 Starting minification process")
	originalFilePath := strings.Replace(file, ".css", "", -1)
	originalFilePath = strings.Replace(originalFilePath, "../CSS/", "", -1)
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

	if err := calculateFileSize(minifiedFilePath); err != nil {
		log.Fatalf("Error calculating file size: %v", err)
	}

	calculateFileSize(minifiedFilePath)
}
// Combine and minify all css files
func bundleAndMinifyAllCSS(distPath string) {
	fmt.Println("💨 Starting minification process")
	// List of CSS files to concatenate and minify
	var files []string

	// get all the css files 
	cssDir, err := os.Open("../CSS")

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

		fileName := "../CSS/" + cssFile
		
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

	if err := calculateFileSize(minifiedFilePath); err != nil {
		log.Fatal(err)
	}
}
// output the size of the file path
func calculateFileSize(filePath string) error {
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