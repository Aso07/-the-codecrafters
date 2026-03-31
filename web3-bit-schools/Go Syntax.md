# What is Syntax 
Syntax can be understood as the rules for writing code correctly, it is the step by step approach used to write code correctly and logically. 
In Go, a Go file consists of the following parts: 
. Package declaration 
. Import packages 
. Functions Statements and expressions 

The example below explains the syntax succinctly: 
package main 
import ("fmt") 
func main() { 
fmt.Println("Hello World!") } 

To explain the example above, in GO, every program is a part of a package. This is defined using the package keyword. Following the example above, the program belongs to the main package. 
The Import "fmt" line allows us to import files ("fmt package") 
More so, the third line in the code is a blank line and Go ignores white space. This makes the code much more readable and clean 
The Func main() {} line is a function. ALL codes inside its braces will be executed 
Furthermore, the fmt.Println() is a function made available from the fmt package. It is used to print outputs in Go