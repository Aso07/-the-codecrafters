# What is a “for loop” in Go?

A for loop is just a way to tell your program: “Repeat this task multiple times.”

Instead of writing the same code again and again, you use a loop to do it automatically.

E.G. Instead of writing "Hello" 5 times, with a loop : for i := 0; i < 5; i++ {
    fmt.Println("Hello")
}

The loop handles the repetition and writes "Hello" 5 times.


Hence, A for loop in Go means: Do this task repeatedly until I tell you to stop.
It can:
Repeat something a fixed number of times
Keep going while a condition is true
Loop through a list of items
Even run forever

## The continue Statement
The continue statement in Go is used inside loops to skip the current iteration and move to the next one.
When continue is reached:
The rest of the code in the loop is skipped
The loop immediately goes to the next iteration

## The break Statement
While the break statement in Go is used inside loops or switch statements to stop the loop or switch immediately.
When break is reached:
The loop or switch ends right away,
And the program continues with the code after the loop or switch.

## Nested Loops
A nested loop runs one full loop inside each iteration of an outer loop.
It is Useful for tasks like printing patterns or iterating over multi-dimensional data.

## The Range Keyword
The range keyword in Go is used to iterate over elements in a collection like arrays, slices, maps, or strings.
