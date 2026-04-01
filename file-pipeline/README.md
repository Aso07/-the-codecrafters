# File Pipeline

## Overview

This program is a File-pipeline which reads an input file, processes each line using a set of transformation rules, and writes the result to an output file.


## How to Run

go run . input.txt output.txt



## Transformation Rules

1. Trim leading and trailing spaces
2. Replace `TODO:` with ` ACTION:`
3. Replace `CLASSIFIED:` with `[REDACTED]:`
4. Remove empty lines and lines with only dashes
5. Add line numbers (001., 002., ...)

---

## Output Format

* The output file starts with:


SENTINEL FIELD REPORT — PROCESSED


* Each line is numbered:


001. Example line
002. Another line




## Example

### Input


   TODO: fix bug   
CLASSIFIED: secret
-----
normal line


### Output


SENTINEL FIELD REPORT — PROCESSED
001.  ACTION: fix bug
002. [REDACTED]: secret
003. normal line



## Notes

* If the input file does not exist, the program shows an error
* Empty or dash-only lines are removed
* The program prints a summary in the terminal after running
