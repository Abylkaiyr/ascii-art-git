package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]

	str := strings.ReplaceAll(arg, "\\n", "\n")

	content, err := ioutil.ReadFile("banner/standard.txt")
	if err != nil {
		fmt.Println("Cannot Read standard.txt file")
	}

	if !checkStdFile(string(content)) {
		fmt.Println("standard.txt file is corrupted")
		return
	}

	if len(str) == 0 {
		return
	}

	for _, l := range str {
		if l < 32 || l > 126 {
			if l == 10 {
				continue
			}
			fmt.Println("Please enter only enter characters from 32 to 126")
			return
		}
	}

	col, let := setFlag()
	colR := setColor(col)
	isWordinText, letT, indexL := setLetter(str, let)
	w1 := str[indexL:(len(letT) + indexL)]
	w2 := str[:indexL]
	w3 := str[(len(letT) + indexL):]

	isThereNewLine, _ := checkNewline(str)
	words1 := strings.Split(str, "\n")
	if isThereNewLine {
		if onlyNewlines(str) {
			for i := 0; i < len(words1[1:]); i++ {
				fmt.Println()
			}
		} else {
			for i := 0; i < len(words1); i++ {
				if words1[i] == "" {
					fmt.Println()
					continue
				} else {
					printWord(string(content), w1, w2, w3, colR, isWordinText)
				}
			}
		}
	} else {
		printWord(string(content), w1, w2, w3, colR, isWordinText)
	}
}

func printWord(content string, w1 string, w2 string, w3 string, col string, isWordinText bool) {
	strArr := [8]string{}
	fontTxt := strings.Split(string(content), "\n")
	if isWordinText {
		for _, l := range w2 {
			pos := int(l)*9 - 287
			if l == 10 {
				continue
			}
			for i := 0; i < 8; i++ {
				strArr[i] += "\u001b[37m" + fontTxt[i+pos]
			}
		}
		for _, l := range w1 {
			pos := int(l)*9 - 287
			if l == 10 {
				continue
			}
			for i := 0; i < 8; i++ {
				strArr[i] += col + fontTxt[i+pos]
			}
		}
		for _, l := range w3 {
			pos := int(l)*9 - 287
			if l == 10 {
				continue
			}
			for i := 0; i < 8; i++ {
				strArr[i] += "\u001b[37m" + fontTxt[i+pos]
			}
		}

	} else {
		fmt.Println("ERR: Letters do not consist in the given text, Please enter in the correct form")
		return
	}

	for i := range strArr {
		fmt.Println(strArr[i])
	}
}

func checkNewline(str string) (bool, int) {
	flag := false
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			flag = true
			count++
		}
	}
	return flag, count
}

func checkStdFile(content string) bool {
	hasher := sha256.New()
	s, err := ioutil.ReadFile("banner/standard.txt")
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}

	l := hasher.Sum(nil)

	hash_std := []byte{195, 236, 117, 132, 251, 126, 207, 189, 115, 158, 107, 63, 111, 99, 253, 31, 229, 87, 210, 174, 62, 36, 248, 112, 115, 13, 156, 248, 178, 85, 158, 148}

	return string(hash_std) == string(l)
}

func onlyNewlines(s string) bool {
	for _, l := range s {
		if l != '\n' {
			return false
		}
	}
	return true
}

func setFlag() (string, string) {
	myWord := os.Args[2]
	fooCmd := flag.NewFlagSet(myWord, flag.ExitOnError)
	color := fooCmd.String("color", "", "a string")
	letter := fooCmd.String("letter", "", "write letter")

	switch os.Args[2] {
	case myWord:
		fooCmd.Parse(os.Args[2:])
	}
	return *color, *letter
}

func setColor(s string) string {
	switch s {
	case "black":
		return "\u001b[30m"
	case "red":
		return "\u001b[31m"
	case "green":
		return "\u001b[32m"
	case "yellow":
		return "\u001b[33m"
	case "blue":
		return "\u001b[34m"
	case "magenta":
		return "\u001b[35m"
	case "teal":
		return "\u001b[36m"
	case "white":
		return "\u001b[37m"
	}
	return ""
}

func setLetter(mWord string, mletter string) (bool, string, int) {
	foundString := ""
	index := 0
	flag := false
	indexL := 0

	for i := 0; i < len(mletter); i++ {
		for j := index; j < len(mWord); j++ {
			if mletter[i] == mWord[j] {
				foundString += string(mletter[i])
				index = j
				break
			}
		}
	}
	if len(foundString) == len(mletter) {
		flag = true
	}

	indexL = strings.Index(mWord, foundString)

	return flag, foundString, indexL
}
