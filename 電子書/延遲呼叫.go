//延遲呼叫
//defer通常用於開啟關閉、連線斷線、上鎖解鎖等成對操作並確保無論流程控制有多複雜，資源一定會釋放
package main

import(
	"net/http"
	"fmt"
	"strings"
	"golang.org/x/net/html"
	"os"
)

func main() {
	//os.Args是[]string(slice)
	//os.Args提供原始命令行參數(包含路徑本身)
	//os.Args[1:]保存所有程序參數(不包含路徑本身)
	for _, url := range os.Args[1:] {
		title(url)
	}
}


func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	//延遲呼叫
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	//不具名函式
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
				fmt.Println(n.FirstChild.Data)
			}
	}
	forEachNode(doc, visitNode, nil)
	return nil
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
