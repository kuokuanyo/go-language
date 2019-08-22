//並行抓取URL內容
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	//使用make建構channel字串
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		//啟動goroutine，是一個並行函式的執行，非同步的呼叫fetch以使用http.Get抓取URL內容
		go fetch(url, ch) 
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //從channel ch 接收
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//創建fetch
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch<- fmt.Sprint(err) //發送到channel ch
		return
	}

	//io.Copy()讀取回應內容並寫道ioutil.discard輸出串流將他拋棄
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() //目的是不洩漏資源
	if err != nil {
		ch<- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}



















