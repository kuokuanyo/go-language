//製作一個套件
//提供github的api
//https://developer.github.com/v3/search/#search-issues
package github

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const IssueURL = "https://api.github.com/search/issues"

//struct欄位名稱都要大寫開頭(匯出)
//unmaeshaling過程中是有區分大小寫的
type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string //Markdown格式
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

//查詢github紀錄
//SearchIssues發出http請求並以json解碼
func SearchIssues(terms []string) (*IssueSearchResult, error) {
	//url.QueryEscape確保會以字面處理
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	//出現問題
	if err != nil {
		return nil, err
	}

	//必須在所有執行路徑上關閉resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	//json.Decoder整個串流解碼程序
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}