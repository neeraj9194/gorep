package main

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
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

func TestReadDirectory(t *testing.T) {
	directory := "test_dir"
	expected := []string{"test1.txt", "test2.txt"}
	files, _ := ReadDirectory(directory)
	if !reflect.DeepEqual(files, expected) {
		t.Fatalf("Failed, %v %v are not equal.", files, expected)
	}
}



