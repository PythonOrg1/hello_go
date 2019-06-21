package search

import "log"

// Result contains the result of a search.
// 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

//匹配搜索的 接口
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

func Display(result chan *Result)  {
	
}

