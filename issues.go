//輸出符合搜尋條件的github紀錄表格
package main

import (
	"fmt"
	"log"
	"os"

	"github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	//出現問題
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}