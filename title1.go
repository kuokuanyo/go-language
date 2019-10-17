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

//檢查伺服器回應的Content-Type標頭並在文件不是HTML時回報錯誤
func title(url string) error {
	resp, err := http.Get(url)
	//出現問題
	if err != nil {
		return err
	}

	//檢查Content-Type 
	ct := resp.Header.Get("Content-Type")
	//strings.HasPrefix檢查字串是否以"..."開頭(回傳bool)
	//出現問題:不是text/html
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	//重複resp.Body.Close()呼叫，可以確保title對失敗在內所有執行路徑關閉網路連線
	resp.Body.Close()
	//出現問題
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
