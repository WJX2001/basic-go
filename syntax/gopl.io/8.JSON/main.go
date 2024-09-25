package main

import (
	"fmt"
	"log"
	"os"
)
import githubAPI "github.com/WJX2001/basic-go/syntax/gopl.io/8.JSON/github"

func main() {
	// MarshalingDemo()
	// githubSearchDemoTest()
	// queryCount()
	// queryComic()
	Xkcd1()
}

func githubSearchDemoTest() {
	result, err := githubAPI.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func queryCount() {
	query := os.Args[1:]
	data, err := githubAPI.SearchIssues(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	githubAPI.Process(data)
}
