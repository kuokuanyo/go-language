/*樣板template
帶有一個或多個稱為action的{{...}}包圍部分的字串或檔案
首先輸出相符的紀錄數量
{{.TotalCount}}被替換成TotalCount欄位的值
{{range .Items}}與{{end}}建構出迴圈
 | 記號讓一個操作結果成為另一個操作參數*/
const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

/*daysAge是函式
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
*/