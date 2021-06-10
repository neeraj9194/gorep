package src

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// Test writing to file
func TestFileWrite(t *testing.T) {
	writeString := "hello"
	f, _ := os.Create("test.txt")

	writeToFile(f, writeString)

	data, _ := ioutil.ReadFile("file.txt")

	if writeString == string(data) {
		t.Fatal("Failed")
	}

	e := os.Remove("test.txt")
	if e != nil {
		log.Fatal(e)
	}
}
