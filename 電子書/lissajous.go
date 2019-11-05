//動態GIF圖
//const宣告、struct型別、組合實字
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//color.顏色 屬於image/color套件
//[].color.Color{}是組合實字
var palette = []color.Color{color.White, color.Black} 

const (
	whiteIndex = 0 //調色盤的第一個顏色
	blackIndex = 1 //調色盤的下一個顏色
)

func main() {
	lissajous(os.Stdout)
}

//創建lissajous
func lissajous(out io.Writer) {
	//const宣告賦予常數名稱，編譯固定的值
	//const宣告可以出現在套件階級
	const (
		cycles = 5 //x震盪旋轉數
		res = 0.001 //角解析度
		size = 100 //畫布大小[-size..+size]
		nframes = 64 //幀數
		delay = 8 //以10ms為單位的幀間隔
	)
	freq := rand.Float64() * 3.0 //y震盪相對頻率
	//gif.GIF{}屬於image/gif套件，也屬於組合實字
	//git.GIT是struct型別，是一群稱為欄位的值
	//實字建構出LoopCount欄位設為nframes的struct值
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //相位差
	//外層迴圈執行64輪(i = 0~63)，每一輪產生一幀，建構出201*201大小黑白兩色圖形
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		//內圈跑二為震盪
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//SetColorIndex設定相對應像素(x, y)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		//anim變數是gif.GIF型別的struct
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //out的型別是io.Writer
}
















