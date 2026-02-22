package main

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@][host:port][/path][?query][#fragment]

	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.Query())
	fmt.Println("Fragment:", parsedURL.Fragment)

	newURL := "https://example.com:8080/path?name=John&age=30#fragment"

	parsedNewURL, err := url.Parse(newURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsedNewURL.Scheme)
	fmt.Println("Host:", parsedNewURL.Host)
	fmt.Println("Path:", parsedNewURL.Path)
	fmt.Println("Query:", parsedNewURL.Query())
	fmt.Println("Fragment:", parsedNewURL.Fragment)

	queryParams := parsedNewURL.Query()
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	//build a new URL
	baseURL := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/path",
		Fragment: "fragment",
		RawQuery: queryParams.Encode(), // encode the query parameters
	}
	newURL = baseURL.String()
	fmt.Println("New URL:", newURL)

	values := url.Values{}
	values.Add("date", "2026-02-17")
	values.Add("time", "10:00:00")
	values.Add("location", "New York")

	newURL = baseURL.String() + "?" + values.Encode()
	fmt.Println("New URL with values:", newURL)
}
