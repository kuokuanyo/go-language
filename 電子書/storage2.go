//測試但並不想真的發送郵件，因此獨立成一個函式儲存在未匯出的套件層級變數中
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

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000
	percent := 100 * used / quota
	if percent < 90 {
		return //ok
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}