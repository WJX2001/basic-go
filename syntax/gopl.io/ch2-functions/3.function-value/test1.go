package main

import (
	"fmt"
	"golang.org/x/net/html"
)

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

// 完善上述两个函数，使其成为通用的HTML输出器
/**
要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。
使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）
*/
