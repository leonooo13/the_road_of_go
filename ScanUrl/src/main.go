package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func scanURL(url string, paths []string) {
	for _, path := range paths {
		fullURL := fmt.Sprintf("%s/%s", url, path)
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", fullURL, err)
			continue
		}
		defer resp.Body.Close()
		fmt.Printf("Checking %s/%s/  [%d]\n", url, path, resp.StatusCode)
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Found: %s\n", fullURL)
		}
	}
}

func loadPathsFromFile(filename string) ([]string, error) {
	var paths []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return paths, nil
}

func main() {
	url := "https://www.example.com" // 替换为你要扫描的网站 URL
	filename := "../top3000.txt"     // 替换为包含路径的文本文件

	paths, err := loadPathsFromFile(filename)
	if err != nil {
		fmt.Printf("Error loading paths from file: %v\n", err)
		return
	}

	scanURL(url, paths)
}
