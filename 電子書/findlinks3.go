//對worklist中每個項目呼叫f
//f回傳項目被加入worklist中
//每個項目最多呼叫f一次
//兩個參數為slice
package main

import (
	"fmt"
	"log"
	"os"
	//"links"
	"net/http"
	"golang.org/x/net/html"
)

func main() {
	//以廣度優先爬網
	//從命令列參數開始
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	//map
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		//worklist為空
		worklist = nil
		for _,item := range items {
			if !seen[item] {
				seen[item] = true
				//f(item)...使f回傳所有項目到worklist中
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//crawler中項目式URL
func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func Extract(url string) ([]string, error) {
	//根據網址取的該網頁內容
	resp, err := http.Get(url)
	//出現錯誤
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	//將html結果解析並回傳一個根節點
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	//slice
	var links []string
	//不具名函式
	//函式內不會使用visitNode，可以直接宣告不具名
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				//解析成相對文件位址resp.Request.URL的URL
				//產生的link為絕對路徑，可用來呼叫http.Get
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links,nil
}

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