package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Load allowed domains from .env file
	allowedDomains := loadAllowedDomains()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		fmt.Println("**** domain origin", origin)

		filePath := r.URL.Query().Get("file")

		if filePath == "" {
			http.NotFound(w, r)
			return
		}

		filePath = "./src/" + filePath

		fmt.Print(filePath)

		if isAllowed(origin, allowedDomains) {
			http.ServeFile(w, r, filePath) // Replace with your CSS file path
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":777", nil))
}

func loadAllowedDomains() []string {
	file, err := os.Open("./domains.txt")
	if err != nil {
		log.Fatal("Error opening domains file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var domains []string

	for scanner.Scan() {
		domain := scanner.Text()
		domains = append(domains, domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading domain file: ", err)
	}

	return domains
}

func isAllowed(origin string, domains []string) bool {
	for _, domain := range domains {
		if strings.Contains(origin, domain) {
			return true
		}
	}
	return false
}
