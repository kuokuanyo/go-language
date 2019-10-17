//技術標誌使用token
//限制20個並行請求
//避免太過平行
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20) //有緩衝的channel

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //取得token
	list, err := Extract(url)
	<-tokens //釋放

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
	return links, nil
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

func main() {
	worklist := make(chan []string)
	var n int //等待發送給worklist的數量

	//從命令列參數開始
	n++
	go func() { worklist <- os.Args[1:] }()

	//並行爬網路
	seen := make(map[string]bool)
	for ; n > 0; n-- { //n降到為0時代表沒有工作要執行了
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
