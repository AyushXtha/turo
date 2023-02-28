package main

import (
    "bufio"
    "fmt"
    "net/url"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <filename> <placeholder>")
        os.Exit(1)
    }

    fileName := os.Args[1]
    placeholder := "FUZZ"
    if len(os.Args) > 2 {
        placeholder = os.Args[2]
    }

    file, err := os.Open(fileName)
    if err != nil {
        fmt.Printf("Error opening file '%s': %s\n", fileName, err)
        os.Exit(1)
    }
    defer file.Close()

    urlStrings := make([]string, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        urlStrings = append(urlStrings, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file '%s': %s\n", fileName, err)
        os.Exit(1)
    }

    urlList := make(map[string][]string)
    for _, urlStr := range urlStrings {
        urlStr = strings.TrimSpace(urlStr)
        if urlStr == "" {
            continue
        }
        parsedUrl, err := url.Parse(strings.Trim(urlStr, "\r"))
        if err != nil {
            fmt.Printf("Error parsing URL '%s': %s\n", urlStr, err)
            continue
        }

        if parsedUrl.RawQuery == "" {
            continue
        }

        urlKey := parsedUrl.Hostname() + parsedUrl.Path + "?"
        params, _ := url.ParseQuery(parsedUrl.RawQuery)

        // Convert the keys to a set to deduplicate them
        keys := make(map[string]bool)
        for key := range params {
            keys[key] = true
        }

        // Convert the set of keys back to a list
        paramList := make([]string, 0, len(keys))
        for key := range keys {
            paramList = append(paramList, key)
        }

        urlList[urlKey] = append(urlList[urlKey], paramList...)
    }

    for urlKey, paramKeys := range urlList {
        paramKeys = uniqueStrings(paramKeys)
        params := make([]string, 0, len(paramKeys))
        for _, key := range paramKeys {
            params = append(params, key+"="+placeholder)
        }
        url := "https://" + urlKey + strings.Join(params,"&")
        fmt.Println(url)
    }
}

func uniqueStrings(strings []string) []string {
    uniqueMap := make(map[string]bool)
    for _, str := range strings {
        uniqueMap[str] = true
    }
    uniqueStrings := make([]string, 0, len(uniqueMap))
    for str := range uniqueMap {
        uniqueStrings = append(uniqueStrings, str)
    }
    return uniqueStrings
}

