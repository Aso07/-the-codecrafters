Array in Go is a collection of elements of the same data type stored at contiguous memory locations. They are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value. The size of an array is fixed. E.g: var arrayName [size]dataType

Array store multiple values of the same type in one variable. e.g. var numbers [3]int = [3]int{10, 20, 30}
fmt.Println(numbers[0])  // Output: 10

## Access Elements of an Array
an element of a specific Array can be accessed by referring to the index number and it can also be changed in the same manner by referring to the index number. Array indexes begings at 0. That means that [0] is the first element, [1] is the second element, etc.