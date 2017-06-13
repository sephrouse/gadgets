package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"bufio"
)

func readFile() {
	f, err := os.Open("needdelete.txt")
	if err != nil {
		fmt.Println("ERROR. open file error.", err)
		return
	}

	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			fmt.Println("ERROR. read line error.", err)
			return
		}

		fmt.Println("line is ", string(line))

		//deleteIndex(string(line))

	}
}

func deleteIndex(index string) {
	fullEsUrl := "http://172.22.157.59:9200/" + index

	request, err := http.NewRequest(http.MethodDelete, fullEsUrl, nil)
	if err != nil {
		fmt.Println("delete Index error. ", index, err)
		return
	}

	deleteClient := &http.Client{}
	resp, err := deleteClient.Do(request)
	if err != nil {
		fmt.Println("delete Index error. ", index, err)
		return
	}

	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("INFO. result of deletion is ", string(result))
	resp.Body.Close()
}

func main() {
	readFile()
}