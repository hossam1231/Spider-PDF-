package urlMatch

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Function to check if a string ends with any of the specified file extensions
func endsWithFileExtension(s string, extensions []string) bool {
    for _, ext := range extensions {
        if strings.HasSuffix(s, ext) {
            return true
        }
    }
    return false
}

// Function to find URLs ending with specified file extensions in a text file
func findURLsInFile(filePath string, extensions []string) ([]string, error) {
    var urls []string

    // Read the content of the file
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    // Regular expression pattern to match URLs
    urlPattern := regexp.MustCompile(`https?://[^\s]+`)

    // Find all URLs in the file content
    matches := urlPattern.FindAllString(string(content), -1)

    // Check if each URL ends with any of the specified file extensions
    for _, url := range matches {
        if endsWithFileExtension(url, extensions) {
            urls = append(urls, url)
        }
    }

    return urls, nil
}

// Function to search for URLs ending with specified file extensions in a directory and its subdirectories
func SearchURLsInDirectory(dirPath string, extensions []string) ([]string, error) {
    var urls []string

    // Walk through the directory and its subdirectories
    err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && filepath.Ext(path) == ".txt" {
            // Process only text files
            fileURLs, err := findURLsInFile(path, extensions)
            if err != nil {
                return err
            }
            urls = append(urls, fileURLs...)
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    return urls, nil
}

