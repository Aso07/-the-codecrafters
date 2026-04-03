# what is a Constant in Go
Go constants are values that cannot be changed once they are declared.
They are defined using the const keyword and are used when you have values that should stay fixed throughout your program.

## Constant Rules

Constant names follow the same naming rules as variables
Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
Constants can be declared both inside and outside of a function

## Constant Types

There are typically two types of constants, which are:
Typed constants: This are values declared with a defined type. E.g., const A int = 1

Untyped constants: while these are declared without a type. e.g. const A = 1

## Multiple Constants Declaration

Multiple constants can be grouped together into a block for readability essense, e.g. const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

Hence, a constant is a variable whose value stays the same and cannot be modified.
