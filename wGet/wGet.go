package wget

import (
    "bufio"
    "io"
    "net/http"
    "os"
    "strings"
)

const (
    bufSize   = 1024 * 8
)

// Write results of a GET request to file. If a fileName is given an empty string then the 
// last chunk of the input url is used as a filename. Eg: http://foo/baz.jar => baz.jar
func Wget(url, fileName string) {
    resp, err := getResponse(url)
    if err != nil {
        return // Return early if there is an error in getting the response
    }
    if fileName == "" {
        urlSplit := strings.Split(url, "/")
        fileName = urlSplit[len(urlSplit)-1]
    }
    writeToFile(fileName, resp)
}

// Make the GET request to a url, return the response and any error encountered
func getResponse(url string) (*http.Response, error) {
    tr := new(http.Transport)
    client := &http.Client{Transport: tr}
    resp, err := client.Get(url)
    if err != nil {
        return nil, err // Return nil and the error if an error occurs
    }
    if resp.StatusCode == http.StatusNotFound {
        return nil, nil // Return nil if the status code is 404
    }
    return resp, nil
}

// Write the response of the GET request to file
func writeToFile(fileName string, resp *http.Response) {
    if resp == nil {
        return // Return early if the response is nil
    }
    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
    if err != nil {
        return // Return early if an error occurs opening the file
    }
    defer file.Close()
    bufferedWriter := bufio.NewWriterSize(file, bufSize)
    _, err = io.Copy(bufferedWriter, resp.Body)
    if err != nil {
        return // Return early if an error occurs during file writing
    }
}
