# HNJobLocations


## Idea
I wanted to get back into writing Go, so I built an application that filters through *Ask HN: Who is Hiring?* threads for job postings based on a supplied location keyword. As of the time of writing, the code finds results posted in the latest month, the previous 12 months, or all time. The full list of postings is not done yet however (list.txt needs some love). This project was built with Go version 1.13 on MacOS.

## Example 
![austin](https://user-images.githubusercontent.com/13093517/83374930-25f08700-a39b-11ea-8aac-288c03b997bd.gif)

## Usage
```
To correctly run this program, please supply two additional arguments.
The first argument is the search term used to find hits of the desired location.
The second argument is a flag used to specify the recency of the results.

E.g.:
	go run hn.go "New York City" --latest
	go run hn.go "Berlin" -a
	go run hn.go "Singapore" --pastYear

The options are as follows:
	-a, --all   		All results from previous Who's Hiring threads.
	-l, --latest		Results from the most recent monthly thread.
	-p, --pastYear		Results from the past 12 months.
```
The program will write the results to an HTML file in your cloned directory. The file is named in snake case after your supplied location search term (e.g. `los_angeles.html` or `barcelona.html`).

## TODO:

- Read API response into the same single struct already in memory rather than make multiple copies
- See how to reduce the number of API calls. If the API doesn't support larger dumps of data, perhaps look into using a web scraper.
- Investigate the necessity of ioutil.ReadAll() on the response body
- Additional command line arguments such as "2020", "2019", etc.
- Convert list.txt into list.yaml for more logical groupings of HN links
- Split hn.go into multiple files for better organization
- Use concurrency when hitting the HN API so that the runtime is reduced 
