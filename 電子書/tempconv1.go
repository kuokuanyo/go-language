//*celsiusFlag符合flag.Value介面
//flag.Value包含String()、Set(string)
package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type celsiusFlag struct{ Celsius Celsius }

//設定String()方法
func (c *celsiusFlag) String() string { return fmt.Sprintf("%g°C", c) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//設定Set(string)方法
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit) //無須檢查錯誤

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invaild temperature %q", s)
}

//定義攝氏旗標名稱
//預設值與使用方法，回傳旗標變數位址
//旗標參數必須有數量與單位
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	var f celsiusFlag = celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
