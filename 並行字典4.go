package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

//檢查
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//建構標準輸入讀取的goroutine
func main() {

	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//偵測到輸入時取消遍歷
	go func() {
		os.Stdin.Read(make([]byte, 1)) //讀取一個位元組
		close(done)
	}()

	//平行遍歷每個root
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:

	for {
		select {
		case <-done:
			//抽乾fileSizes已允許現有goroutine完成
			for range fileSizes {
				//什麼都沒做
			}
			return
		case size, ok := <-fileSizes:

			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //取得token
	case <-done:
		return nil //取消
	}
	defer func() { <-sema }() //釋放token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
