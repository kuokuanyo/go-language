/*
將visit改成不具名函式
直接加入links並以forEachNode處理
只需要pre函式，nil作為post參數
*/
package links

import(
	"fmt"
	"net/http"
	"golang.org/x/net/html"
)

//回傳slice, error
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

