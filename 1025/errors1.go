package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		//log套件會有時間
		log.Fatalln(err) //等同於fmt.println()
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("not true")
	}
	return 42, nil
}

/*
type error interface {
	Error() string
}
*/

/*
errors.New有Error方法，可以回傳type error
func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

*/
