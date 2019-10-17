package main

import(
	"os"
	"fmt"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _,link := range links {
			fmt.Println(link)
		}
	}
}

//對url執行HTTP的GET請求，解析HTML
//回應，取出與回傳連結
func findLinks(url string) ([]string, error) {
	//如果http.Get失敗會回傳HTTP錯誤給呼叫方而沒其他動作
	resp, err := http.Get(url)
	//出現問題
	if err != nil {
		//直接回傳
		return nil, error
	}
	//第二個與第三個錯誤從fmt.Errorf取得額外資訊
	if resp.StatusCode != http.StatusOK {
		//必須確保resp.Body被關閉使發生錯誤時也能釋放網路資源
		resp.Body.Close()
		//fmt.Errorf函式將錯誤訊息使用fmt.Sprintf格式化並回傳新的error值
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	//html.Parse失敗不會回傳HTML解析程序的錯誤
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/*
示範僅回傳的函式(bare return)
對某個URL的HTML發出HTTP GET 請求並回傳文件的字詞與圖片數量
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return  回傳0, 0, err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		retrun 回傳0, 0, err
	}
	words, images = countWordsAndImages(doc)
	return 回傳words, images, nil
}
*/