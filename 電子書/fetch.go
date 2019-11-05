/*將HTTP回應寫到檔案而非標準輸出
以path.Base從URL的最後一個元件產生檔名
下載URL並回傳檔案名稱與長度
*/
package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	//問題
	if err != nil {
		return "", 0, err
	}
	//延遲函式
	defer resp.Body.Close()
	//從URL取得物件最後一個元素
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	//開啟檔案供寫入
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	//關閉檔案
	if closeErr := f.close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
