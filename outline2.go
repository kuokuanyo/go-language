//對n的每個x節點呼叫pre(x)、post(x)
//pre再造訪子節點前呼叫
//post在之後呼叫
package main

import(
	"golang.org/x/net/html"
	"fmt"
)

func main() {
	forEachNode(doc, startElement, endElement)
}

//第二、三個參數皆為函式
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//startElement在造訪子節點前呼叫
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		//%*s中的*輸出字串在前面加上設定數量的空白
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
		depth++
	}
}

//endElement在造訪子節點後呼叫
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		//%*s中的*輸出字串在前面加上設定數量的空白
		fmt.Printf("%*s<%s>\n", depth * 2, "", n.Data)
	}
}