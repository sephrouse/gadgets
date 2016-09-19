package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Examination ...
type Examination struct {
	Questions [20]Question
}

// Question ...
type Question struct {
	Subject string
	Options [4]string
	Answer  int
}

// cOptionPrefixes ...
var cOptionPrefixes = []string{"A. ", "B. ", "C. ", "D. "}

// OneWord ...
type OneWord struct {
	enWord string
	cnWord string
}

var words []OneWord
var totalWords int

const cQuestionNum = 20

func storeAllWords() error {
	fh, err := os.Open("arrangedVocabulary.txt")
	if err != nil {
		fmt.Println("storeAllWords error. ", err)
		return err
	}

	defer fh.Close()

	r := bufio.NewReader(fh)

	i := 1

	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("storeAllWords read string error. ", i, err)
			return err
		}

		if err == io.EOF {
			break
		}

		i++

		singleWordArray := strings.Split(line, " && ")
		var singleWord OneWord
		singleWord.enWord = singleWordArray[0]
		singleWord.cnWord = strings.TrimRight(singleWordArray[1], "\n")

		words = append(words, singleWord)
	}

	totalWords = len(words)
	fmt.Println("done storeAllWords. ", i, totalWords)

	return nil
}

func generateQuestions() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 1; i <= cQuestionNum; i++ {
		no := r.Intn(totalWords)
		fmt.Println(i, words[no].enWord, words[no].cnWord, no)
	}

	return
}

// implement generateCnToEnExamination as follows.
func generateEnToCnExamination() Examination {
	var e Examination

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < cQuestionNum; i++ {
		no := r.Intn(totalWords)
		e.Questions[i].Subject = words[no].enWord

		//e.Questions[i].Options[0] = words[no].cnWord

		answerPos := r.Intn(4)
		e.Questions[i].Answer = answerPos
		e.Questions[i].Options[answerPos] = words[no].cnWord

		for j := 0; j < 4; j++ {
			if j == answerPos {
				e.Questions[i].Options[j] = cOptionPrefixes[j] + e.Questions[i].Options[j]
				continue
			}
			other := r.Intn(totalWords)
			e.Questions[i].Options[j] = cOptionPrefixes[j] + words[other].cnWord
		}

		//sort.Sort(e.Questions[i].Options)

		fmt.Println(i, e.Questions[i].Subject, e.Questions[i].Options[0], e.Questions[i].Options[1], e.Questions[i].Options[2], e.Questions[i].Options[3], e.Questions[i].Answer)
	}

	return e
}

func generateCnToEnExamination() Examination {
	var e Examination

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < cQuestionNum; i++ {
		no := r.Intn(totalWords)
		e.Questions[i].Subject = words[no].cnWord

		//e.Questions[i].Options[0] = words[no].cnWord

		answerPos := r.Intn(4)
		e.Questions[i].Answer = answerPos
		e.Questions[i].Options[answerPos] = words[no].enWord

		for j := 0; j < 4; j++ {
			if j == answerPos {
				e.Questions[i].Options[j] = cOptionPrefixes[j] + e.Questions[i].Options[j]
				continue
			}
			other := r.Intn(totalWords)
			e.Questions[i].Options[j] = cOptionPrefixes[j] + words[other].enWord
		}

		//sort.Sort(e.Questions[i].Options)

		fmt.Println(i, e.Questions[i].Subject, e.Questions[i].Options[0], e.Questions[i].Options[1], e.Questions[i].Options[2], e.Questions[i].Options[3], e.Questions[i].Answer)
	}

	return e
}

func generatePaper() error {
	r := bufio.NewReader(os.Stdin)
	flag := true

	for flag {
		fmt.Println("please input the instruction. generate, examine or quit.")
		line, _, _ := r.ReadLine()
		// if err != nil {
		// 	fmt.Println("generatePaper error. ", err)
		// 	return err
		// }

		cmd := string(line)

		switch cmd {
		case "q":
			fmt.Println("quit the program. bye!")
			flag = false
		case "g":
			fmt.Println("generate a new paper as follows.")
			// generate questions.
			generateQuestions()
		case "e":
			fmt.Println("generate a new en to cn examination as follows.")
			generateEnToCnExamination()
		case "f":
			fmt.Println("generate a new cn to en examination as follows.")
			generateCnToEnExamination()
		default:
			fmt.Println("your input is wrong. please enter q, e or g to go on.")
		}
	}

	return nil
}

func main() {
	// read all words from vocabulary. and store them in a list.
	if nil != storeAllWords() {
		fmt.Println("main error.")
		return
	}

	// generate examination paper which contains 20 different questions.
	if nil != generatePaper() {
		fmt.Println("main calls generatePaper error.")
		return
	}
}
