package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}

// TODO: 根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年

// 处理数据
func Process(data *IssuesSearchResult) {
	// 创建一个hash表
	hash := make(map[string][]*Issue, 3) // 映射值是一个指向Issue类型的切片, 初始容量为3
	hash["不到一月"] = make([]*Issue, 0)
	hash["不到一年"] = make([]*Issue, 0)
	hash["超过一年"] = make([]*Issue, 0)

	// 遍历数据
	for _, val := range data.Items {
		// 计算从创建时间到现在的持续时间
		duration := time.Since(val.CreatedAt)
		if duration.Hours() < 24*30 { // 假设每月30天
			hash["不到一月"] = append(hash["不到一月"], val)
		} else if duration.Hours() < 24*365 { // 假设每年365天
			hash["不到一年"] = append(hash["不到一年"], val)
		} else {
			hash["超过一年"] = append(hash["超过一年"], val)
		}
	}

	fmt.Printf("数据总条数，%v\n", data.TotalCount)
	for key, val := range hash {
		fmt.Printf("%s:\n", key)
		for _, v := range val {
			fmt.Printf("#%-5d %9.9s %.55s %9.9v\n",
				v.Number, v.User.Login, v.Title, v.CreatedAt)
		}
	}
}
