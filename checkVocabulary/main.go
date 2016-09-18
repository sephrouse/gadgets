package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fh, err := os.Open("arrangedVocabulary.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fh.Close()

	r := bufio.NewReader(fh)

	var i int
	i = 1
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(i, err)
			return
		}

		if err == io.EOF {
			break
		}

		if strings.Contains(line, " && ") == false {
			fmt.Println(i, "does not contain  && .")
		}

		i++
	}

	fmt.Println("total check ", i, " lines.")

	return
}
