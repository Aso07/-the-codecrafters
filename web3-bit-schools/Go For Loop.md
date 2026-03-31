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