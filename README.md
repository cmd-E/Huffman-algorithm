# Huffman-algorithm
Implementation of [Huffman algorithm coding](https://en.wikipedia.org/wiki/Huffman_coding) in golang 
## Usage 
```console
user@pc:~/Huffman algorithm$ go build
user@pc:~/Huffman algorithm$ ./huffman-algorithm -w hello
2021/02/28 10:00:21 Code for h is 110
2021/02/28 10:00:21 Code for e is 111
2021/02/28 10:00:21 Code for l is 0
2021/02/28 10:00:21 Code for o is 10
```
type `./huffman-algorithm -h` to get full list of commands
## Projects structure

`btAndLinkedList` - package where binary tree and linked list and related methods are placed

`occpackage` - package where all structs and methods related to counting occurrences and sorting are placed
