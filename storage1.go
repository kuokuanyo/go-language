//顯示提供儲存體網路服務的容量檢查邏輯
//當使用者用量超過90%時系統會發出警告郵件
package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 { return usage[username] }

//送件者設定
//不要把密碼放在原始碼中
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"
const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000
	percent := 100 * used / quota
	if percent < 90 {
		return //ok
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}
