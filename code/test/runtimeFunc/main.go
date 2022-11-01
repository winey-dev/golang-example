package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	fmt.Println("vim-go")

	fmt.Println(Func())
	TestA()
}

func Func() string {
	//pc, file, line, _ := runtime.Caller(1)
	//temp := runtime.FuncForPC(pc).Name()
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}

func TestA() {
	fmt.Println(Func())
}
