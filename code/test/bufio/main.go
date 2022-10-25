package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio NewReader의 ReadLine 동작 방식
// Enter 입력시 io.EOF 에러가 아니라 "" 빈 값을 반환
// \n까지 복사되지는 않음
func main() {
	reader := bufio.NewReader(os.Stdin)
	var readLine []byte
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Printf("eof\n")
			break
		}
		if string(buf) == "" {
			fmt.Printf("no input\n")
			break
		}

		readLine = append(readLine, buf...)
		readLine = append(readLine, []byte("\n")...)
	}
	fmt.Printf("readLine\n%s\n", string(readLine))
}
