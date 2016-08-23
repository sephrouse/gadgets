package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//var recordNum int = 0

func getAuthorAndChinese(content string) (englishAuthor, chineseQuote, chineseAuthor string) {
	englishAuthor = ""
	chineseQuote = ""
	chineseAuthor = ""
	// if strings.ContainsRune(content, rune('「')) && strings.ContainsRune(content, rune('」')) {
	// 	fmt.Println("there is no 「 or 」 in content. ", content)
	// 	return englishAuthor, chineseQuote, chineseAuthor
	// }
	//recordNum++

	tempArray := strings.Split(content, "「")
	if len(tempArray) != 2 {
		// fmt.Println("!!!!!! ", recordNum)
		// fmt.Println("the length of array is less than 2. ", content)
		return englishAuthor, chineseQuote, chineseAuthor
	}

	englishAuthor = tempArray[0]

	chineseContent := "「" + tempArray[1]

	tempArray = strings.Split(chineseContent, "」")
	if len(tempArray) != 2 {
		// fmt.Println("!!!!!! !! ", recordNum)
		// fmt.Println("the length of array is less than 2. ", content)
		return englishAuthor, chineseQuote, chineseAuthor
	}

	chineseQuote = tempArray[0] + "」"
	chineseAuthor = tempArray[1]

	englishAuthor = strings.TrimSpace(englishAuthor)
	chineseQuote = strings.TrimSpace(chineseQuote)
	chineseAuthor = strings.TrimSpace(chineseAuthor)

	return englishAuthor, chineseQuote, chineseAuthor
}

func getAllQuotesFromSinglePage(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("getAllQuotesFromSinglePage: goquery new document error. ", err)
		return err
	}

	doc.Find(".entry").Each(func(i int, s *goquery.Selection) {
		englishQuote := s.Find("p").First().Text()
		lastContent := s.Find("div").Next().Next().Next().Text()
		//chinese := s.Find("p").Next().Text()
		lastContent = strings.TrimSpace(lastContent)

		englishAuthor, chineseQuote, chineseAuthor := getAuthorAndChinese(lastContent)
		if englishAuthor == "" && chineseQuote == "" && chineseAuthor == "" {
			return
		}

		// fmt.Println(englishQuote)
		// fmt.Println(englishAuthor)
		// fmt.Println(chineseQuote)
		// fmt.Println(chineseAuthor)

		//use " && " to split the following string and then get each information.
		fmt.Printf("%s && %s && %s && %s \n", englishQuote, englishAuthor, chineseQuote, chineseAuthor)
	})

	return nil
}

func getAllQuotes(pageNum int) error {
	var i int
	var url string

	url = "http://www.dailyenglishquote.com/"

	// get the first page, that is the home page.
	err := getAllQuotesFromSinglePage(url)
	if err != nil {
		fmt.Println("getAllQuotes: getAllQuotesFromSinglePage the first page couldn't be retrieved successfully.")
		return err
	}

	for i = 2; i <= pageNum; i++ {
		url = fmt.Sprintf("http://www.dailyenglishquote.com/page/%d/", i)

		err := getAllQuotesFromSinglePage(url)
		if err != nil {
			fmt.Println("getAllQuotes: getAllQuotesFromSinglePage the ", i, " page couldn't be retrieved successfully.")
			return err
		}
	}

	return nil
}

func check(fileName string) {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println("check: open file failed.")
		return
	}
	defer fi.Close()

	r := bufio.NewReader(fi)

	var i int = 1
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("check: readline error. ", err)
			return
		}
		if err == io.EOF {
			break
		}

		lineArray := strings.Split(line, " && ")
		if len(lineArray) != 4 {
			fmt.Println(i, " line array does not equal to 4.")
		}

		if strings.TrimSpace(lineArray[0]) == "" {
			fmt.Println(i, " there is no content in the english quote.")
		}

		if []byte(lineArray[0])[0] == ' ' {
			fmt.Println(i, " there is a space in the front of the quote.")
		}

		if strings.Contains(lineArray[1], "“") {
			fmt.Println(i, " there is also a quote in author field.")
		}

		i++
	}

	fmt.Println("check: end. total records is ", i)
}

func revise(fileName string) {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Println("check: open file failed.")
		return
	}
	defer fi.Close()

	r := bufio.NewReader(fi)

	var i int = 1
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("check: readline error. ", err)
			return
		}
		if err == io.EOF {
			break
		}

		lineArray := strings.Split(line, " && ")

		if strings.Contains(lineArray[1], "“") && strings.Contains(lineArray[1], "”") {
			tempArray := strings.Split(lineArray[1], "—")
			lineArray[1] = "—" + tempArray[1]
		}

		lineArray[3] = strings.Replace(lineArray[3], "\n", "", 1)

		fmt.Println(i, " && "+lineArray[0]+" && "+lineArray[1]+" && "+lineArray[2]+" && "+lineArray[3])
		i++
	}
}

// main ...
func main() {
	// the URL is http://www.dailyenglishquote.com/ , total pages are 268
	if len(os.Args) == 3 {
		//check output file.
		// usage: xxx.exe check output.txt
		if os.Args[1] == "check" && os.Args[2] != "" {
			check(os.Args[2])
		}

		// usage: xxx.exe revise output.txt > result.txt
		// revise无法修正全部问题。output.txt还需要手动删除一些非法的格式。
		if os.Args[1] == "revise" && os.Args[2] != "" {
			revise(os.Args[2])
		}
		return
	} else {
		err := getAllQuotes(268)
		if err != nil {
			fmt.Println("main: getAllQuotes is failed.")
			return
		}
	}

}
