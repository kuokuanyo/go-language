//對worklist中每個項目呼叫f
//f回傳項目被加入worklist中
//每個項目最多呼叫f一次
//兩個參數為slice
package main

import (
    "fmt"
    "log"
    "os"
    "links"
    
    "golang.org/x/net/html"
)

func main() {
    //以廣度優先爬網
    //從命令列參數開始
    breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
    //map
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        //worklist為空
        worklist = nil
        for _,item := range items {
            if !seen[item] {
                seen[item] = true
                //f(item)...使f回傳所有項目到worklist中
                worklist = append(worklist, f(item)...)
            }
        }
    }
}

//crawler中項目式URL
func crawl(url string) []string {
    fmt.Println(url)
    list, err := Extract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}

