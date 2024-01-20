package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// does not seem to be working right now
type Icon struct {
    ClassName string
    URL       string
}

func LoopOverIcons(cssFilePath, outputFilePath string) {
    file, err := os.Open(cssFilePath)
    if err != nil {
        log.Fatalf("Error opening CSS file: %v", err)
    }
    defer file.Close()

    icons := make([]Icon, 0)
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "{")
        if len(parts) >= 2 {
            className := strings.TrimSpace(parts[0])
            if strings.HasPrefix(className, ".icon-") && strings.Contains(className, "background-image") {
                parts = strings.Split(parts[1], ";")
                for _, part := range parts {
                    if strings.Contains(part, "background-image:") {
                        urlStart := strings.Index(part, "url(")
                        urlEnd := strings.Index(part, ")")
                        if urlStart != -1 && urlEnd != -1 {
                            url := strings.TrimSpace(part[urlStart+4:urlEnd])
                            icons = append(icons, Icon{ClassName: className, URL: url})
                        }
                    }
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading CSS file: %v", err)
    }

    // Write the icons data to the output file
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        log.Fatalf("Error creating output file: %v", err)
    }
    defer outputFile.Close()

    for _, icon := range icons {
        _, err := outputFile.WriteString(fmt.Sprintf("Class: %s\nURL: %s\n\n", icon.ClassName, icon.URL))
        if err != nil {
            log.Fatalf("Error writing to output file: %v", err)
        }
    }

    fmt.Println("Icons data written to", outputFilePath)
}

