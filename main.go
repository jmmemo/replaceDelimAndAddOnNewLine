package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	sourceFile      = "aaa.txt"
	destFile        = "dest.txt"
	readStopBy      = 'n'
	oldStrToReplace = "\\n" //windows下是双斜
)

func main() {
	open, err := os.Open(sourceFile)
	if err != nil {
		panic(err)
	}

	dest, err := os.OpenFile(destFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(open)

	for {
		ss, err := reader.ReadString(byte(readStopBy))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			fmt.Println("replace done.")
			break
		}

		//fmt.Println(ss)
		replace := strings.Replace(ss, oldStrToReplace, "", -1)
		_, err = dest.WriteString(replace + "\n")
		if err != nil {
			panic(err)
		}
	}

}
