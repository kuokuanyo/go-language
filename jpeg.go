//從標準輸入讀取PNG圖形並輸出JPEG到標準輸出
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"

	//_只使用套件裡的init函式
	_ "image/png" //登記PNG解碼程序
)

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

/*
解碼程序
package png
func Decode(r io.Reader) (image.Image, error) 
func DecodeConfig(r io.Reader) (image.Config, error)
func init() {
	image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}
*/