package language

import "io"

func writeHeader(w io.Writer, contentType string) error {
	//write參數要求為slice，需要[]byte(...)做型別轉換
	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
		return err
	}
	if _, err := w.Write([]byte("contentType")); err != nil {
		return err
	}
	// ...
}

//----------------------------------------------------------------------------

//將s寫入w
//如果w具有WriteString方法則呼叫它而不是w.Write
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	//判別w是否有WriteString方法(符合介面)
	//介面.(型別) : 行為判別
	if sw, ok := w.(stringWriter); ok {
		//如果符合此介面，呼叫
		return sw.WriteString(s)
	}
	//如果不符合此介面
	return w.Write([]byte(s))
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, "contentType"); err != nil {
		return err
	}
}
