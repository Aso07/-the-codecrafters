# SENTINEL — Command & Control Console

This is a Command Line Interface (CLI) system built in Go as part of the CodeCrafters Hackathon.

This tool acts as a "Command & Control (C&C) console", allowing users to execute multiple operations at once, operations such as calculations, base conversions, and string transformations, all from one interface.


## Features

* Multi-command CLI system
* Calculator with advanced operations
* Base converter (binary, hex, decimal)
* String transformer
* Command piping (`|`) support
* Command history tracking
* Session control (clear, exit)


## Available Commands

### Calculator

User type:
calc add 5 3
calc sub 10 4
calc mul 2 7
calc div 10 2
calc mod 10 3
calc pow 2 3


Extras:

User type:
calc last       → shows last result
calc history    → shows calculation history



### Base Converter

User type:
base dec 255
base hex 1E
base bin 1010


* Decimal → Binary & Hex
* Binary/Hex → Decimal


### String Transformer

User type:
str upper hello world
str lower HELLO
str cap hello world
str title the lord of the rings
str snake Hello World
str reverse hello world


### System Commands

User type:
help        -> show all commands
history     -> show command history
clear       -> clear session
exit        -> exit program



## Command Piping

You can chain commands using `|`:

User type:
calc add 5 5 | base dec

This takes the result of the first command and uses it in the next.


## Example Usage

User type:
C&C> calc add 10 5
   Result: 15

C&C> base dec 15
   Binary : 1111
   Hex    : F

C&C> str upper hello
 HELLO



## Concepts Demonstrated

* CLI design and input handling
* Command parsing and execution
* Error handling and validation
* String manipulation
* Number conversion
* State management (history, last result)


## How to Run

1. Navigate to the folder:

```bash
cd thecodecrafters
```
2. Enter into the file:
```bash
cd command-and-control

2. Run the program:

```bash
go run main.go
```



## Group: GOROUTINE

** Members:

* Agene Okoh
* Janai Egeonu
* Chibuzo Maxwell
* Emmanuel Unogwu
* Edwin Ejembi
* Blessing Anebi
* Faith Ejembi
* Ruth Agi
* Ummulkusum Musa


## Summary

SENTINEL is a unified CLI system that combines multiple tools into one interface, demonstrating efficient command handling and modular design in Go.
