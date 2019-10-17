//回傳第一個非空值的title元素
//若不只一個則回報錯誤
package main

import(
	"golang.org/x/net/html"
	"fmt"
)

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	//defer、不具名函式
	defer func() {
		//recover恢復流程
		//recover放在defer函式中才有效果
		//在其他時間執行recover會回傳nil(沒有效益)
		switch p := recover(); p {
		case nil:
			//no panic
		case bailout{}:
			//預期中的panic
			err = fmt.Errorf("multiple title elements")
		default:
			//未預期的panic
			panic(p)
		}
	}() //延遲函式後面要有()

	//如果發現一個以上非空白標題則脫離遞迴
	for EachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
		n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) //多個title元素
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}