package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type XkcdResp struct {
	Img   string
	Title string
	Link  string
}

func Xkcd1() {
	xkcdUrl := "https://xkcd.com/"
	xkcdSuffix := "/info.0.json"
	f, err := os.OpenFile("storage.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	var max int
	// fmt.Println("------------------- how match you want -----------------------")
	// fmt.Scanln(&max)
	for {
		fmt.Println("------------------- how many do you want -----------------------")
		fmt.Scanln(&max)

		if max > 1 {
			break // 输入大于 1 时，退出循环
		}

		fmt.Println("输入无效，必须大于1，请重新输入。") // 提示用户重新输入
	}

	for i := 1; i < max; i++ {
		url := xkcdUrl + fmt.Sprint(i) + xkcdSuffix
		fmt.Printf("query url: %v\n", url)
		body, err := queryUrl(url)
		if err != nil {
			fmt.Printf("query url error: %v\n", err)
			break
		}
		var parseBody XkcdResp
		if err := json.Unmarshal(body, &parseBody); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
			continue
		}
		row := parseBody.Img + " " + parseBody.Title + parseBody.Link + "\n"
		if _, err := f.Write([]byte(row)); err != nil {
			fmt.Printf("write info error: %v\n", err)
		}
	}
	f.Close()

	rf, rErr := os.OpenFile("storage.txt", os.O_RDONLY, 0)
	if rErr != nil {
		log.Fatal(err)
	}

	var order int
	fmt.Println("------------------- want order -----------------------")
	fmt.Scanln(&order)

	reader := bufio.NewReader(rf)
	flag := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF: %#v\n", line)
				break
			}
		}
		if flag == order-1 {
			fmt.Printf("%v", line)
		}
		flag++
	}

	defer f.Close()
}

func parseBody(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func queryUrl(url string) ([]byte, error) {
	// 1. query url
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("query url error: %v", err)
		return nil, err
	}
	// 2. parse response body
	body, err := parseBody(resp)
	if err != nil {
		fmt.Printf("parse response error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}
