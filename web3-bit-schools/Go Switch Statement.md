# What is Go Switch Statement
In Go, the switch statement is a clean way to handle multiple conditions, often replacing long if-else chains. It’s flexible and has some unique features compared to other languages.
A switch evaluates an expression and compares it against different case values.

# Types of Switch Statement
* Single Case (One value per case) e.g,
day := 3

switch day {
case 1:
    fmt.Println("Sunday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown day")
}
This is the simplest form each case checks one value.

* Multiple Values in One Case (Multi-case)
Go allows multiple values in a single case, separated by commas. e.g:
day := 6

switch day {
case 1, 2, 3, 4, 5:
    fmt.Println("Weekday")
case 6, 7:
    fmt.Println("Weekend")
}

## Important Features of Go Switch
1. No break needed
In Go, switch automatically stops after a match.

## When to Use Multi vs Single Case
Single case: when each value has different logic
Multi case: when several values share the same logic