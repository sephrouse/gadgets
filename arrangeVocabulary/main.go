package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var arrangedVocabulary []string

func processLine(line string, lineNo int) {
	//fmt.Println(lineNo)

	if strings.Contains(line, "/") {
		line = strings.Replace(line, "\r\n", "", -1)
		tempArray := strings.Split(line, "/")

		arrangedWord := strings.TrimSpace(tempArray[0]) + " && " + strings.TrimSpace(tempArray[2])

		if len(tempArray) > 3 {
			for i := 3; i < len(tempArray); i++ {
				arrangedWord = arrangedWord + "/" + tempArray[i]
			}
		}

		arrangedVocabulary = append(arrangedVocabulary, arrangedWord)
	} else if strings.Contains(line, " && ") {
		line = strings.Replace(line, "\r\n", "", -1)
		arrangedVocabulary = append(arrangedVocabulary, line)
	} else {
		fmt.Println("there is no pronunciation in line. ", lineNo)
		line = strings.Replace(line, "\r\n", "", -1)
		arrangedVocabulary = append(arrangedVocabulary, line)
	}
}

func outputArrangedVocabulary(newFh *os.File) {
	for _, c := range arrangedVocabulary {
		newFh.WriteString(c + "\n")
	}
}

// 1 获取github上的https://github.com/fanhongtao/IELTS/blob/master/IELTS%20Word%20List.txt
// 2 将该文件中的音标/ [] {}替换为/
// 3 用本程序检查不符合格式的行。检查到后，手动修改。对于短语，手动增加 && 。
// 4 本程序格式化输出整理后的文件后，还会存在中译里存在音标。需要手动清除。
func main() {
	if len(os.Args) != 2 {
		fmt.Println("wrong parameters.")
		return
	}

	fileName := os.Args[1]

	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error. ", fileName)
		return
	}
	defer fh.Close()

	r := bufio.NewReader(fh)
	var lineNo int
	lineNo = 1

	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("read line error. ", err)
			return
		}

		if err == io.EOF {
			fmt.Println("total line is ", lineNo)
			break
		}

		processLine(line, lineNo)

		lineNo++
	}

	newFh, err := os.OpenFile("arrangedVocabulary.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println("open output file error. ", err)
		return
	}
	defer newFh.Close()

	outputArrangedVocabulary(newFh)

}
