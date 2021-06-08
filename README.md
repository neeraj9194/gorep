# gorep
A grep implementation in golang. A command line program that implements Unix command `grep` like functionality. 
The program has the following features.

- Ability to search for a string in a file. Feel free to assume case-sensitive and exact word matches if required for simplicity’s sake.

```
$ ./gorep "search_string" filename.txt
I found the search_string in the file.
```

- Ability to search for a string from standard input (using pipe)

```
$ ls | ./gorep foo
```
will produce
```
barbazfoo
food
```


- Ability to write output to a file instead of a standard out.

```
$ ./gorep -o out.txt lorem loreipsum.txt 
```

will create an out.txt file with the output from `gorep`. for example,

```
$ cat out.txt
lorem ipsum
a dummy text usually contains lorem ipsum
```

- Ability to search in multiple files.
- Ability to search for a string recursively in any of the files in a given directory.

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
