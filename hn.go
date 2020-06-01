package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type post struct {
	By    string `json:"by"`
	ID    int64  `json:"id"`
	Kids  []int  `json:"kids"`
	Text  string `json:"text"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

func loadCSS(cityName string) string {
	return `
	<!DOCTYPE html>
	<html>
		<head>
		<meta charset="utf-8" name="viewport">
		<link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
		<title>` + cityName + `</title>
		<style>

		body{
			background: #de302f;
			font-family: 'Roboto', sans-serif;
			font-size: 14px;
			line-height: 1.4;
			color: #ffffff;
			font-weight: 100;
		  }
	  
		  .container{
			position: relative;
			max-width: 600px;
			height: auto;
			border: 2px solid #ffffff;
			margin: 100px auto;
			padding: 30px;
			box-sizing: border-box;
	  
		  }
	  
		  .container:after{
			position: absolute;
			width: 50px;
			height: 50px;
			border-top: 0px solid #ffffff;
			border-right: 2px solid #ffffff;
			border-bottom: 2px solid #ffffff;
			border-left: 0px solid #ffffff;
			top:100%;
			left: 50%;
			margin-left: -25px;
			content: '';
			transform: rotate(45deg);
			margin-top: -25px;
			background: #de302f;
		  }

		  h1 {
			text-align: center;
			font-size: 32px;
			margin-bottom: -35px;
		  }

		  h2 {
			  text-align: center;
			  color: #ffce00;
		  }

		  a {
			  color: #ffce00;
		  }

		  a:hover{
			position: relative;
			top: -1px;
		  }

		</style> 
		</head>

		<body>
		<h1>` + cityName + "</h1>"
}

func validateURL(site string) string {
	//check if url
	_, err := url.ParseRequestURI(site)
	if err != nil {
		log.Fatal(err)
	}

	//check if HackerNews
	prefix := "https://news.ycombinator.com/item?id="
	if !strings.HasPrefix(site, prefix) {
		log.Fatal("Please supply a valid HackerNews URL")
	}

	//parse post ID
	postID := site[strings.IndexByte(site, '=')+1:]

	return postID
}

func printUsage() {
	usage := `
	To correctly run this program, please supply two additional arguments within quotations.
	These additional arguments are used to specify the desired location and the recency of the results.
	
	E.g.:
	- go run hn.go "New York City" "latest"
	- go run hn.go "Berlin" "year"
	- go run hn.go "Singapore" "all"
	`
	fmt.Println(usage)
	os.Exit(0)
}

func queryAPIForJSON(postID string) []byte {
	//query HN api for JSON response
	request := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%s.json", postID)
	resp, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}

	//read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	return body
}

func unmarshalIntoStruct(data []byte) post {
	var post post

	if err := json.Unmarshal(data, &post); err != nil {
		log.Fatal(err)
	}
	return post
}

func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func findCityInComments(post post, cityName string) []string {

	filtered := []string{}
	// for each comment, search if city name is in text
	for _, kid := range post.Kids {
		data := queryAPIForJSON(strconv.Itoa(kid))
		post := unmarshalIntoStruct(data)
		foundCity := caseInsensitiveContains(post.Text, cityName)

		if foundCity {
			fmt.Printf("Found submission by user %s\n", post.By)
			filtered = append(filtered, post.Text)
		}
	}

	return filtered

}

func createFile(results [][]string, cityName string) {
	// make the HTML file
	fileName := strings.ReplaceAll(strings.ToLower(cityName), " ", "_") + ".html"
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	html := loadCSS(cityName)
	entryCount := 0

	// loop over the contents of the array and write to file
	for _, result := range results {
		for _, entry := range result {
			entryCount++
			div := "<div class=\"container\">" + fmt.Sprintf("<h2>%d</h2>", entryCount) + entry + "</div>"
			html += div

		}
	}
	html += "</body></html>"
	f.WriteString(html)

	fmt.Println("\nFinished search!")
	fmt.Printf("%d total result(s) now written to \"%s\"\n", entryCount, fileName)
}

func startSearch(postID, cityName string) []string {
	//grab prerequisite information
	data := queryAPIForJSON(postID)
	post := unmarshalIntoStruct(data)

	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	title := reg.ReplaceAllString(post.Title[strings.IndexByte(post.Title, '?')+1:], "")
	fmt.Printf("\nSearching references from %s...\n", strings.TrimSpace(title))

	//begin search
	result := findCityInComments(post, cityName)
	if len(result) == 0 {
		fmt.Printf("No submissions found for \"%s\"\n", cityName)
	}
	return result
}

func readFile() []string {
	file, err := os.Open("./list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := []string{}
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}

func validateArgs(args []string) bool {
	if len(args) == 3 {
		var arguments = []string{"latest", "year", "all"}

		for _, item := range arguments {
			if item == args[2] {
				return true
			}
		}
	}
	return false
}

func parseArgs(args []string) (string, int) {
	cityName := strings.Title(string(args[1]))
	limitations := map[string]int{
		"all":    -1,
		"latest": 1,
		"year":   12,
	}
	return cityName, limitations[strings.ToLower(string(args[2]))]
}

func main() {
	args := os.Args
	if !validateArgs(args) {
		printUsage()
	}
	cityName, limit := parseArgs(args)
	list := readFile()
	var results [][]string

	for i, post := range list {

		if i < limit || limit == -1 {
			postID := validateURL(post)
			result := startSearch(postID, cityName)
			if len(result) != 0 {
				results = append(results, result)
			}
		} else {
			break
		}

	}
	if len(results) != 0 {
		createFile(results, cityName)
	}
}
