package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// MarshalingDemo()
	// githubSearchDemoTest()
	// queryCount()
	queryComic()
}

func githubSearchDemoTest() {
	result, err := SearchIssues(os.Args[1:])
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
	data, err := SearchIssues(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	Process(data)
}

// 对漫画服务进行查询
func queryComic() {
	// 获取离线列表
	// GetSaveInfo()
	key := os.Args[1:]
	ResLink(strings.Join(key, " "))
}
