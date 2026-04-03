# What is Go Output Functions

Go Output Functions are functions used to display (print) output to the screen. They are mainly provided by the fmt package.
Go has three functions to generate output text, which include:
    Print()
    Println()
    Printf()

The Print() Function prints its arguments with their default format. e.g: 
package main
import ("fmt")
func main() {
  var i string = "Hello World"

  fmt.Print(i)
}

By implication this prints as it is given in a streight line.

The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end

The Printf() Function:
The Printf() function is a function that first formats its argument based on the given formatting verb and then prints them.

Three formatting verbs will be  used for explaining this:

    %v is used to print the value of the arguments
    %T is used to print the type of the arguments
    %s is used to print the text or strings of an arguement
