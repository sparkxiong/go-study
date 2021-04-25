package main

import (
	"errors"
	"fmt"
)

type errorString string

//自定义error实现error接口
type MyError struct {
	Op   string
	Path string
	Err  error
}

//panic recover
func positive2(i int) bool {
	if i == 0 {
		panic("UnDefined")
	}
	return i > -1
}

func check2(i int) {
	defer func() {
		if recover() != nil {
			fmt.Println(i, "is neither")
		}
	}()
	if positive2(i) {
		fmt.Println(i, "is positive")
	} else {
		fmt.Println(i, "is negative")
	}
}
func positive(i int) (bool, error) {
	if i == 0 {
		return false, errors.New("Undefined")
	}
	return i > -1, nil

}

func check(i int) {
	Pos, err := positive(i)
	if err != nil {
		fmt.Println(i, err)
		return
	}
	if Pos {
		fmt.Println(i, "is positive")
	} else {
		fmt.Println(i, "is negative")
	}

}

func (e *MyError) Error() string {
	return e.Op + e.Path + e.Err.Error()
}

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}
	check(0)
	check2(0)

}
