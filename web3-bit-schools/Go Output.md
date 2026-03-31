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
By implication this prints as it vis giving in a streight line.

The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end

The Printf() Function:
The Printf() function first formats its argument based on the given formatting verb and then prints them.

Two formatting verbs are used for this:

    %v is used to print the value of the arguments
    %T is used to print the type of the arguments
