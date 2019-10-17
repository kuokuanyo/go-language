//處理錯誤
//嘗試加接觸伺服器
//假設所有嘗試均式敗則回報錯誤
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		//成功
		if err == nil {
			return nil 
		}
		log.Printf("server not responding (%s); retrying...", err) 
		//指數後退
		time.Sleep(time.Second << uint(tries)) 
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
