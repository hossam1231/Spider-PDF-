package main

import (
	"fmt"

	"github.com/hahnicity/go-wget"
	"github.com/hossam1231/Spider-PDF/m/v2/urlMatch"
)



func downloadFile(url string) () {
 // Search for URLs in the specified directory
  	wget.Wget(url, "")
    return 
}



func main() {
    // Specify the directory to search
    directory := "./output/www_theavocagroup_com.txt"

    // Specify the file extensions to search for in URLs
    fileExtensions := []string{".pdf", ".docx", ".xlsx"} // Add more file extensions as needed

    // Search for URLs in the specified directory
    urls, err := urlMatch.SearchURLsInDirectory(directory, fileExtensions)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the found URLs
    fmt.Println("Found URLs:")
    for _, url := range urls {
        fmt.Println(url)
		downloadFile(url)
		fmt.Println("Downloading", url)
    }


}
