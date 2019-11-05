package tempconv

//攝氏轉華氏
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

//華氏轉攝氏
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }