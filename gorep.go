package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	outputFile = flag.String("out", "", "Output file path")
	inputDir   = flag.String("dir", "", "Search all files in a directory.")
	pattern    string
	fileNames  []string
)

const boldColorRed = "\033[1;31m%v\033[0m"

func main() {
	flag.Parse()
	pattern = fmt.Sprintf("(%v)", flag.Arg(0))
	if len(flag.Args()) < 1 {
		fmt.Println("Wrong Usage. \nTry './gorep --help' for more information.")
		os.Exit(0)
	}
	if *inputDir != "" {
		var err error
		fileNames, err = ReadDirectory(*inputDir)
		if err != nil {
			log.Fatal("Could not read directory.")
			os.Exit(0)
		}
	} else {
		fileNames = flag.Args()[1:]
	}

	if len(fileNames) > 0 {
		fmt.Printf("Searching %v in file(s) %v\n\n", flag.Arg(0), fileNames)
		readFromMultiFiles(fileNames)
	} else {
		fmt.Printf("Searching: %v ...\n\n", flag.Arg(0))
		readFromStdIn()
	}
}

// Initilize flags
func init() {
	flag.StringVar(outputFile, "o", "", "Output file path")
	flag.StringVar(inputDir, "d", "", "Search all files in a directory.")
}

// Create a new file truncate if already exists.
func createNewFile(filePath string) *os.File {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return f
}

// write a string to an existing file.
func writeToFile(destFile *os.File, text string) {

	_, err2 := destFile.WriteString(text + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}
}

func readFromMultiFiles(fileList []string) {
	var destFile *os.File
	if *outputFile != "" {
		destFile = createNewFile(*outputFile)
		defer destFile.Close()
	} else {
		destFile = nil
	}
	for _, i := range fileList {
		scanAndWrite(i, destFile)
	}
}

func readFromStdIn() {
	var destFile *os.File
	if *outputFile != "" {
		destFile = createNewFile(*outputFile)
		defer destFile.Close()
	} else {
		destFile = nil
	}
	scanAndWrite("", destFile)

}

// scan stdin or a file if given, match the pattern and write to file or stdout.
func scanAndWrite(srcFilePath string, destFile *os.File) {
	var scanner *bufio.Scanner
	if srcFilePath != "" {
		file, err := os.Open(srcFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	// regular expression match
	regExp, regErr := regexp.Compile(pattern)

	if regErr != nil {
		log.Fatal(regErr)
	}

	for scanner.Scan() {
		if regExp.MatchString(scanner.Text()) {
			// If matched replace the matched string with a red appened text.
			repString := regExp.ReplaceAllString(scanner.Text(), fmt.Sprintf(boldColorRed, "$1"))
			if srcFilePath != "" {
				repString = fmt.Sprintf("%v: %v", srcFilePath, repString)
			}
			if destFile != nil {
				writeToFile(destFile, repString)
			} else {
				fmt.Println(repString)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func ReadDirectory(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, filepath.Join(*inputDir, file.Name()))
	}
	return files, nil
}
