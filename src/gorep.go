package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
)

const boldColorRed = "\033[1;31m%v\033[0m"

func Gorep(wg *sync.WaitGroup, srcFilePath string, destFile *os.File, pattern string) {
	defer wg.Done()

	// read
	scanner, f := scan(srcFilePath)

	// search
	i := 1
	defer f.Close()
	for scanner.Scan() {
		// fmt.Printf("HELL: %v\n", srcFilePath)
		color := destFile == nil
		matchedString, found := search(pattern, scanner.Text(), color)
		// fmt.Printf("HOO%v\n", matchedString)
		if found {
			if srcFilePath != "" {
				matchedString = fmt.Sprintf("%v:L%v: %v", srcFilePath, i, matchedString)
			}
			// write
			write(matchedString, destFile)
		}
		i++
	}
}

func search(pattern string, text string, color bool) (string, bool) {
	regExp, regErr := regexp.Compile(pattern)
	if regErr != nil {
		log.Fatal(regErr)
	}

	if regExp.MatchString(text) {
		repString := text
		// If matched replace the matched string with a red appened text.
		if color {
			repString = regExp.ReplaceAllString(repString, fmt.Sprintf(boldColorRed, "$1"))
		}
		return repString, true
	}
	return "", false
}

func scan(srcFilePath string) (*bufio.Scanner, *os.File) {
	var scanner *bufio.Scanner
	var file *os.File

	if srcFilePath != "" {
		file, err := os.Open(srcFilePath)
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return scanner, file
}

func write(text string, destFile *os.File) {
	if destFile != nil {
		writeToFile(destFile, text)
	} else {
		fmt.Println(text)
	}
}

// write a string to an existing file.
func writeToFile(destFile *os.File, text string) {

	_, err2 := destFile.WriteString(text + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}
}
