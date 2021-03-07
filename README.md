# Huffman-algorithm
![GitHub](https://img.shields.io/github/license/cmd-e/Huffman-algorithm?style=plastic)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cmd-e/Huffman-algorithm?style=plastic)

Implementation of [Huffman algorithm coding](https://en.wikipedia.org/wiki/Huffman_coding) in golang 
## Usage
Provide word only:
```console
user@pc:~/Huffman algorithm$ go build
user@pc:~/Huffman algorithm$ ./huffman-algorithm -w hello
Code for h is 110
Code for e is 111
Code for l is 0  
Code for o is 10 
```
Provide file with word:
```console
user@pc:~/Huffman algorithm$ go build
user@pc:~/Huffman algorithm$ ./huffman-algorithm -f very-long-word.txt
Code for V is 100100001
Code for i is 000
Code for t is 1000
Code for a is 1100
Code for e is 1110
...
```
Provide file with predefined occurrences for each symbols:
```console
user@pc:~/Huffman algorithm$ go build
user@pc:~/Huffman algorithm$ ./huffman-algorithm -p occurrences.txt
Code for h is 110
Code for e is 111
Code for l is 0  
Code for o is 10 
```
occurrences.txt contains:
```
h-1
# Comment
e-1
l-2
o-1
```
Lines which has `#` as first symbol are ignored

Provide file with predefined occurrences for each symbols but treat numbers as probabilities:
```console
user@pc:~/Huffman algorithm$ go build
user@pc:~/Huffman algorithm$ ./huffman-algorithm -p occurrences.txt -prob
Code for h is 110
Code for e is 111
Code for l is 0  
Code for o is 10 
```
occurrences.txt contains:
```
h-0.5
# Comment
e-0.5
l-1
o-0.5
```
Lines which has `#` as first symbol are ignored

Provide file with predefined occurrences for each symbols and set own separator:
```console
user@pc:~/Huffman algorithm$ go build
user@pc:~/Huffman algorithm$ ./huffman-algorithm -p occurrences.txt -s "*"
Code for h is 110
Code for e is 111
Code for l is 0  
Code for o is 10 
```
occurrences.txt contains:
```
h*1
# Comment
e*1
l*2
o*1
```
Lines which has `#` as first symbol are ignored

Type `./huffman-algorithm -h` to get full list of commands:
```console
user@pc:~/Huffman algorithm$ go build
user@pc:~/Huffman algorithm$ ./huffman-algorithm -h
-w - input to encode
-f - file with input to encode
-w = -f, but from file. If -w and -f flags are defined, -f will be executed
-p - path to file where custom occurrences for all symbols in word are defined
-prob - available if -p is defined. Occurrences for symbols are treated as possibilities
```
## Projects structure

`btAndLinkedList` - package where binary tree and linked list and related methods are placed

`occpackage` - package where all structs and methods related to counting occurrences and sorting are placed

`tests` - package with tests

`userinput` - package to accept input and flags from user

`help.txt` - file with help to display if user reqested it