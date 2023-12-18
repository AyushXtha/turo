package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func dedu(urlStr string, fuzz string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err,urlStr)
        fmt.Println(urlStr)
		return ""
	}
	
	queryDict, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		fmt.Println("Error parsing query parameters:", err,urlStr)
        fmt.Println(urlStr)
		return ""
	}

	if fuzz != "" {
		for key := range queryDict {
			queryDict.Set(key, fuzz)
		}
	} else {
		for key, values := range queryDict {
			queryDict.Set(key, values[0])
		}
	}

	parsedURL.RawQuery = queryDict.Encode()
	return parsedURL.String()
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <filename> [fuzz]")
		return
	}

	filename := args[1]
	var fuzz string
	if len(args) > 2 {
		fuzz = args[2]
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    buf := make([]byte, 0, 64*1024)
    scanner.Buffer(buf, 10240*10240)
	urls := []string{}
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	urlDict := map[string][]string{}

	for _, urlStr := range urls {
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
            fmt.Println("Error parsing URL:", err,": ",urlStr)
            fmt.Println(urlStr)
			continue
		}
		
		if parsedURL.RawQuery != "" {
			urlKey := fmt.Sprintf("%s://%s%s?", parsedURL.Scheme, parsedURL.Host, parsedURL.Path)
			queryDict, err := url.ParseQuery(parsedURL.RawQuery)
			if err != nil {
                fmt.Println("Error parsing query parameters:", err,": ",urlStr)
                fmt.Println(urlStr)
				continue
			}
			
			urlValue := []string{}
			seenParams := map[string]bool{}
			for key, values := range queryDict {
				for _, value := range values {
					if !seenParams[key] {
						urlValue = append(urlValue, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
						seenParams[key] = true
					}
				}
			}
			urlDict[urlKey] = append(urlDict[urlKey], urlValue...)
		}
	}

	for urlKey, urlValues := range urlDict {
		joinedValues := strings.Join(urlValues, "&")
		uralo := fmt.Sprintf("%s%s", urlKey, joinedValues)
		newURL := dedu(uralo, fuzz)
		fmt.Println(newURL)
	}
}

