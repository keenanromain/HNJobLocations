# hn_job_locations


## Idea
I wanted to get back into writing Go, so I built an application that filters through the HackerNews API for historic postings based on a location keyword. This project was built with Go version 1.13 on MacOS.


## Usage
```
To correctly run this program, please supply two additional arguments within quotations.
These additional arguments are used to specify the desired location and the recency of the results.
	
	E.g.:
	- go run hn.go "New York City" "latest"
	- go run hn.go "Berlin" "year"
	- go run hn.go "Singapore" "all"
```

## Improvements (Notes to self)


- Read into the same single struct already in memory rather than make multiple copies
- See how to reduce the number of api calls. If the api doesn't support larger dumps of data, perhaps look into using a web scraper.
- Investigate the necessity of ioutil.ReadAll() on the response body
- Additional command line arguments such as "2020", "2019", and so on
- Convert list.txt to list.yaml for logical grouping of HN links
- Split hn.go into multiple files for better organization
