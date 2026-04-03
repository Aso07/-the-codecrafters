A comment is a text and it is ignored upon execution. It is used to explain code and make it more understandable, readable. go supports single or multi line comments. 

## Go single line comments

In Go a single line comments start with two forward slashes(//). all the text between // and the end of the line is ignored by the compiler and will not be executed. 
For example: 

// This is a comment package main 
import ("fmt") 
func main() { 
// This is a comment
fmt.Println("Hello World!") } 

## Go Multi line comments 

In this regard the comment start with /* and ends with */. All the text between it will be ignored and not executed by the compiler. 

For example: 
package main 
import ("fmt") 
func main() { 

/* The code below will print Hello World to the screen, and it is amazing */

fmt.Println("Hello World!") } 
Hence, it is up to the person to choose which form of comments he wants to use.