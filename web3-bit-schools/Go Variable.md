# what is a variable

Variables in Go are like containers that holds data value.

## Types of Variable 

In Go there are different types of variables, e.g:

* strings: which holds text and there are always wraped in double quotes. e.g. "Hello word"

* int: This holds whole numbers e.g 12345

* Bool: This holds one of two possible values, either True or False.

* Float32: which holds decimal numbers, this are numbers which are not whole but has remainders. e.g. 12.2, -32.54

## How to creat Variable

In Go there is typically two ways to declare a variable:

with the Var keyword:

with the := sign:

## Multiple variable declaration

This means you can define several variables in one line instead of writing separate lines for each variable. e.g. var a, b, c, d int = 1, 3, 5, 7

If you use the type keyword, it is only possible to declare one type of variable per line.

If the type keyword is not specified, you can declare different types of variables on the same lin

## Variable declaration in a block

This means declaring multiple variables together inside a pair of parentheses using the var keyword. e.g. var (
     a int
     b int = 1
     c string = "hello"
)

## variable naming rules

This are the guildlines that determines how someone can correctly name variables in Go.

Variable name containletters, numbers,and undesrscores
It cannot start with a number
it cannot use Go keywords(e.g func, var, package)
Names are case-sensitive (age and Age are different)
These rules ensure your variable names are valid and understood by the Go compiler.

## Multi word variable

This are variables made up of more than one word, written in a readable format. In Go, they are usually written using camelCase, where the first word is lowercase and each new word starts with a capital letter. e.g.
UserName := "Aso"

The essense is that it makes code easier to read and understand.