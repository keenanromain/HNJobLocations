# HNJobLocations


## Idea
I wanted to get back into writing Go, so I built an application that filters through *Ask HN: Who is Hiring?* threads for job postings based on a supplied location keyword. As of the time of writing, the code finds results posted in the latest month, the previous 12 months, or all time. The full list of postings is not done yet however (list.txt needs some love). This project was built with Go version 1.13 on MacOS.

## Example 
![austin](https://user-images.githubusercontent.com/13093517/83374930-25f08700-a39b-11ea-8aac-288c03b997bd.gif)

## Usage
```
To correctly run this program, please supply two additional arguments within quotations.
These additional arguments are used to specify the desired location and the recency of the results.
	
	E.g.:
	- go run hn.go "New York City" "latest"
	- go run hn.go "Berlin" "year"
	- go run hn.go "Singapore" "all"
```
The program will write the results to an HTML file in your cloned directory. The file is named in snake case after your choosen location keyword (such as `los_angeles.html` or `barcelona.html`).

## TODO:

- Read API response into the same single struct already in memory rather than make multiple copies
- See how to reduce the number of API calls. If the API doesn't support larger dumps of data, perhaps look into using a web scraper.
- Investigate the necessity of ioutil.ReadAll() on the response body
- Additional command line arguments such as "2020", "2019", etc.
- Convert list.txt into list.yaml for more logical groupings of HN links
- Split hn.go into multiple files for better organization
- Use concurrency when hitting the HN API so that the runtime is reduced 
- convert second command line argument into a flag so that "all" becomes -a or --all, "latest" becomes -l or --latest, etc.
