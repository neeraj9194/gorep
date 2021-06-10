# gorep
A grep implementation in golang. A command line program that implements Unix command `grep` like functionality. 
The program has the following features.

## Install

To install you can use makefile or build using commands

```
make build
OR
go build main.go
```

To run game,

```
make run
OR
./gorep   // after building
```

To run tests
```
make test
OR
go test ./... -v
```

## Usage

Usage: ./gorep [OPTION]... [PATTERN] [FILES]... <br>
Search for PATTERN in each FILE. <br>
Example: ./gorep 'hello world' menu.txt main.txt

## Examples

- Search for a string in a file

```
$ ./gorep "search_string" filename.txt
filename.txt:L2:I found the search_string in the file.
```

- Search for a string from standard input (using pipe)

```
$ ls | ./gorep foo
```
will produce
```
barbazfoo
food
```


- Write output to a file instead of a standard out for example,

```
$ ./gorep -o out.txt lorem loreipsum.txt // It will create an out.txt file with the output from `gorep`.
$ cat out.txt
lorem ipsum
a dummy text usually contains lorem ipsum
```

- Search in multiple files.
- And alos search for a string recursively in any of the files in a given directory.

```
$ ./gorep -d tests test 
tests/test1.txt:this is a test file
tests/test1.txt:one can test a program by running test cases
tests/inner/test2.txt:this file contains a test line

$ ./gorep "test" test1.txt test2.txt
tests/test1.txt:this is a test file
tests/test1.txt:one can test a program by running test cases
tests/inner/test2.txt:this file contains a test line
```
