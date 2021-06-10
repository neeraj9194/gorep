package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	gorep "github.com/neeraj9194/gorep/src"
)

var (
	outputFile = flag.String("o", "", "Output file path")
	inputDir   = flag.String("d", "", "Search all files in a directory.")
	pattern    string
	fileNames  []string
)

func main() {
	flag.Parse()

	// Pattern
	pattern = fmt.Sprintf("(%v)", flag.Arg(0))
	if len(flag.Args()) < 1 {
		fmt.Println("Wrong Usage. \nTry './gorep --help' for more information.")
		os.Exit(0)
	}

	// Output
	var destFile *os.File
	if *outputFile != "" {
		destFile = createNewFile(*outputFile)
		defer destFile.Close()
		defer fmt.Printf("Results are recorded in %v.\n", *outputFile)
	} else {
		destFile = nil
	}

	// Input
	if *inputDir != "" {
		fileNames = ReadDirectory(*inputDir)
	} else if len(flag.Args()[1:]) > 0 {
		fileNames = flag.Args()[1:]
	}
	gorepHandler(fileNames, destFile, pattern)
}

// Create a new file truncate if already exists.
func createNewFile(filePath string) *os.File {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return f
}

func gorepHandler(fileList []string, destFile *os.File, pattern string) {
	var wg sync.WaitGroup
	if len(fileNames) > 0 {
		fmt.Printf("Searching %v in file(s) %v\n\n", flag.Arg(0), fileNames)
		for _, i := range fileList {
			wg.Add(1)
			go gorep.Gorep(&wg, i, destFile, pattern)
		}
	} else {
		fmt.Printf("Searching: %v ...\n\n", flag.Arg(0))
		wg.Add(1)
		go gorep.Gorep(&wg, "", destFile, pattern)
	}
	wg.Wait()
}

func ReadDirectory(root string) []string {
	var files []string
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return files
}
