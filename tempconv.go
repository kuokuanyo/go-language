//製作成套件
//將不同溫度單位放在不同型別當中
package tempconv

import "fmt"

//設定兩個型別(溫度單位)
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g C", c)}
func (f Fahrenheit) String() string { return fmt.Sprintf("%g F", f)} 
