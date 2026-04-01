# String Transformer

A string Transformer is a command-line tool that cleans and reformats corrupted text.

## How to Run
bash
go run main.go


## Commands

| Command | What it does |
|---------|-------------|
| `upper <text>` | Converts text to UPPERCASE |
| `lower <text>` | Converts text to lowercase |
| `cap <text>` | Capitalises the first letter of every word |
| `title <text>` | Like cap, but small connector words stay lowercase |
| `snake <text>` | Converts text to snake_case |
| `reverse <text>` | Reverses each word individually |
| `exit` | Shuts down the program |

## Example Session

SENTINEL STRING TRANSFORMER — ONLINE
──────────────────────────────────────
> upper sentinel is online
> SENTINEL IS ONLINE

> lower ALERT LEVEL FIVE
> alert level five

> cap director adaeze okonkwo
> Director Adaeze Okonkwo

> title the fall of the western power grid
> The Fall of the Western Power Grid

> snake Operation Gopher Protocol!
> operation_gopher_protocol

> reverse Go is fun
> oG si nuf

> exit
Shutting down String Transformer. Goodbye.


## Author

- Name: [Agene Okoh]  
- Squad: [Goroutine]