package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin) //func NewDecoder(r io.Reader) *Decoder
	var stack []string              //slice，元素名稱堆疊
	for {
		tok, err := dec.Token() //func (d *Decoder) Token() (Token, error)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) { //介面.(type)回傳型別實字
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

//依序回報x是否帶有y元素
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

/*會運用到encoding/xml套件
type Name struct {
	Space, Local string
}
type Attr struct {
	Name  Name
	Value string
}
type Token interface{}
type StartElement struct {
	Name Name
	Attr []Attr
}
type EndElement struct {Name Name}
type CharData []byte
type Comment []byte
type Decoder struct {...}

func NewDecoder(r io.Reader) *Decoder
func (d *Decoder) Token() (Token, error)
*/
