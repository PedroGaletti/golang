package main

/*
1. How do you check a variable type at runtime?
A: The type switch is the best way to check a variable's type at runtime. The type switch evaluate variable's by type rather than value.
Each switch contains at least one case, which acts as a conditional statement, and a default case, which executes if none of the cases are true.

2. How do you concatenates strings?
A: The easiest way to concatenate Strings is to use the concatenation operator (+), which allow you to add strings as you would numerical values

3. Explain GO interfaces. What are they and how do they work?
A: Interfaces are a special type in Go that define a set of method signatures but do not provide implementations. Values of interface type can hold any value that implements those methods.
Interfaces essentially act as placeholders for methods that will have multiple implementations based on what object is using it.

5. Format a string without printing it?
A: The easiest way to format without printing is to use the fmt.Sprintf(), which returns a string without printing it.

6. Explain the difference between concurrent and parallelism in Golang?
A: Concurrency is when your program can handle multiple tasks at once while. Parallelism is when you program can execute multiple tasks at once using multiple processors.
Concurrency is a property of a program that allows you to have multiple tasks in progress at the same time, but not necessarily executing at the same time.
Parallalism is a runtime property where two or more tasks are executed at the same time.
Parallelism can therefore be a means to achieve the property of concurrency, but it's just one of many means available to you
The key tools for concurrency in GOlang are goroutines and channels. Goroutines are concurrent lightweight threads while channels allow goroutines to communicate with each other during execution.

9. What are the differences between unbuffered and buffered channels?
A: For unbuffered channel, the sender will block on the channel until the receiver receives the data from the channel,
while the receiver will also block on the channel until sender sends data into the channel
Compared with unbuffered counterpart, the send of buffered channel will block when there is no empty slot of the channel, while the receiver will block on the channel when it is empty.

10. What is type assertion in Golang? What does it do?
A: A type assertion takes an interface value and retrieves from it a value of the specified explicit type. Type conversion is used to convert dissimilar types in Golang.

11. Deriveds type in Golang?
A: Interface, map, channels

12. About the loops statement in Go:
A: If condition is available, then for loop executes as long as condition is true and if range is available, then for loop executes for each item in the range.

13. Do we have a while loop in golang?
A: No

14. What Cap method does?
A: The cap function returns the capacity of slice as how many elements it can be accomodate.

15. About the interfaces:
A: Go programming provides another data type called interfaces which represents a set of methods signatures.
Struct data type implements these interfaces to have method definitions for the method signature of the interfaces.

16. About the structs:
A: To access any member of a structure, we use the member access operator ().
You would use struct keyword to define variables of structure type
You can pass a structure as a function argument in very similar way as you pass any other variable or pointer.

17. How can I get the type of object?
A: Using the reflect lib inside Golang using the valueOf and Kink functions or TypeOf

18. Garbage Collection?
A: Is a mechanism to find memory space that is allocated recently but is no longer needed, hence the need to deallocate them to create a clean slate so that further allocation can be done on the same space or so that memory can be reused

19. how map works?
A: Golang Maps is a compilation of unordered key-value pairs. It is commonly used because it offers simple searches and values with the aid of keys that can be retrieved

20. Does Go support overloading?
A: No, Golang doesn't no support overloading

21. Does Go support pointer arithmetics?
A: No, Golang doesn't no support pointer arithmetics

22. Does Go support genreric programming?
A: No, Golang doesn't no support genreric programming

23. What is the purpose of goto statement?
A: goto transfers control to the labeled statement and a label is a string and is created by setting a semicolon after the label.
*/

func main() {

}
