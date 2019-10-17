package main

import (
	"log"
	"os"
	"sync"
	"thumbnail"
)

/*利用func ImageFile(infile string) (string, error)
從infile讀取圖片並在同一個目錄輸出縮圖，回傳的檔名類似"foo.thumb.jpg"
*/

//製作指定檔案縮圖
func makeThumbnails(filenames []string) {
	for _, f := range filenames { //迴圈成單個檔名
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//平行製作指定檔案的縮圖
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{}) //channel
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	//等待goroutine完成
	for range filenames {
		<-ch
	}
}

//平行製作指定檔案的實體，若有步驟錯誤則回傳錯誤
//有個bug，當它遇到第一個非空的錯誤時會回傳錯誤給呼叫方，使得errors這個channel沒有goroutine來抽乾
//會發生goroutine洩漏，可能會使程式停止或耗盡記憶體
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err //goroutine洩漏
		}
	}

	return nil
}

//平行製作指定檔案的實體
//以任意順序回傳檔名，若步驟錯誤則回傳錯誤
//使用有緩衝的channel回傳縮圖檔案名稱與錯誤
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames)) //有緩衝channel
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil { //有問題
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

//最終版
//回傳新檔案所占總位元組數
//在啟動每一個goroutine前遞增一個計數器，在每一個goroutine完成時遞減它
//對從channel收到的每個檔案製作縮圖
func thumbnail6(filenames <-chan string) int64 { //filenames只接收
	sizes := make(chan int64)  //channel
	var wg sync.WaitGroup      //goroutine數(計數器)
	for f := range filenames { //接收
		wg.Add(1) //必須在goroutine啟動前呼叫
		//運算
		go func(f string) { //goroutine
			defer wg.Done() //延遲，確保計數器在錯誤情況下也會遞減
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) //忽略錯誤
			sizes <- info.Size()
		}(f)
	}

	//終結
	go func() { //等待與關閉操作必須與sizes迴圈並行
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
