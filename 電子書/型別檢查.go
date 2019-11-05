package language

import "fmt"

//第一種判別檢查方式 if-else
func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
		//介面.(型別) : 判斷是否符合改型別
	} else if _, ok := x.(int); ok {
		//符合int型別的話
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		//符合int型別的話
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		//符合bool
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return sqlQuoteString(s)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

//第二種:switch
//運算元為x.(type)
func sqlQuote(x interface{}) string {
	switch x := x.(type) { //x.(type)回傳type實字
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(s)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
