package task

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
DAY 3

	🟢 Morning Assignment: Looping Challenges

📌 Task 1: Write a program that prints numbers from 1 to 50 but:

Skips numbers divisible by 5.
Stops if it encounters a number divisible by 9.
*/
func Day3Morn1() {
	for i := 1; i <= 50; i++ {
		if i%5 == 0 {
			continue
		} else if i%9 == 0 {
			break
		} else {
			fmt.Println(i)
		}
	}
}

/*
🏆 Morning Practice Assignments
✅ Task 1: Age Checker
📝 Write a program that asks the user for their age and prints:

"You are a child" if age < 13
"You are a teenager" if 13 ≤ age < 18
"You are an adult" if age ≥ 18
*/
func Day3Morn11() {
	age := 0
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	if age < 13 {
		fmt.Println("You are a child")
	} else if age < 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are an adult")
	}
}

/*
📌 Task 2: Implement a countdown timer from 10 to 1, then print "Blast off!".
*/
func Day3Morn2() {
	for i := 10; i >= 1; i-- {
		println(i)
	}
}

/*
✅ Task 2: Simple Calculator (Using switch)
📝 Ask the user to enter two numbers and an operator (+, -, *, /). Then print the result.
*/
func Day3Morn22() {
	a := 0
	b := 0
	op := ""
	print("Enter a number, an operator and another number_")
	fmt.Scan(&a, &op, &b)
	switch op {
	case "+":
		println(a + b)
	case "-":
		println(a - b)
	case "*":
		println(a * b)
	case "/":
		println(a / b)
	default:
		println("Invalid input")
	}
}

/*
📌 Task 3: Write a program that asks for a positive number from the user.

If the user enters 0 or a negative number, keep asking them until they enter a valid positive number.
*/
func Day3Morn3() {
	var input int
	for {
		fmt.Print("Enter a positive non-zero integer: ")
		fmt.Scanln(&input)
		if input <= 0 {
			println("Please try again!")
			continue
		}
		println("Your number was: ", input)
		break
	}
}

/*
🟢 Evening Assignment: Conditional Challenges
📌 Task 4: Create a simple login system:

Ask the user for a username and password.
If the username is "admin" and the password is "1234", print "Access Granted".
Otherwise, print "Access Denied".
*/
func Day3Eve1() {
	var (
		username string
		password string
	)

	for {
		fmt.Print("Enter Username and Password: ")
		fmt.Scan(&username, &password)
		if username == "admin" && password == "1234" {
			println("Access Granted")
			break
		} else {
			println("Access Denied")
		}
	}
}

/*
✅ Task 3: FizzBuzz Challenge
Write a program that prints numbers from 1 to 20, but:

If the number is divisible by 3, print "Fizz"
If it's divisible by 5, print "Buzz"
If it's divisible by both 3 and 5, print "FizzBuzz"
Otherwise, print the number itself.
*/
func Day3Eve2() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("FIzz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}

/*
✅ Task 4: Print Multiplication Table (1-10)
Write a program that prints a multiplication table for numbers 1 to 10 using loops.
*/
func Day3Eve3() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("\n%d Multiplication Table\n", i)
		for num := 1; num <= 10; num++ {
			fmt.Printf("%d * %d = %v\n", i, num, num*i)
		}
	}
}

// DAY 4
/*
📌 Task 1: Writing a Simple Function

Define a function called sayHello that prints "Hello, Go learner!".
Call this function from main().
📌 Task 2: Functions with Parameters

Modify sayHello to accept a name parameter and print "Hello, [name]!".
*/
func SayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

/*
📌 Task 3: Function Returning a Value

Create a function called addNumbers that takes two integers and returns their sum.
Call this function and print the result.
*/
func AddNumbers(a, b int) (sum int) {
	sum = a + b
	return
}

/*
📌 Task 4: Multiple Return Values

Functions in Go can return multiple values.
Create a function called calculate that takes two numbers and returns both their sum and difference.

📌 Task 5: Named Return Values

Modify calculate to use named return values.
*/
func Calculate(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}

/*
📌 Task 6: Variadic Functions

Go allows functions to accept a variable number of arguments using ....
Call sumAll(1, 2, 3, 4, 5) and print the result.
*/
func SumAll(numbers ...int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

/*
📌 Task 7: Variadic Functions

	Logging system

Variadic functions are often used in logging systems to allow flexible message formatting:
*/
func LogMessage(level string, messages ...string) {
	fmt.Println(level + ":")
	for _, msg := range messages {
		fmt.Println("-", msg)
	}
}

//DAY 5
/*
✅ Task 1: Creating and Using Pointers
Create a simple program that:

Declares an integer variable
Creates a pointer to the variable
Prints the memory address and the value using the pointer
📌 Key Takeaways:

& (Address-of operator) is used to get the memory address of a variable.
* (Dereference operator) is used to access the value stored at the memory address.
*/
func D5M1() {

	number := 42
	ptr := &number

	fmt.Println("Value of number:", number)
	fmt.Println("Memory address of number:", ptr)
	fmt.Println("Value accessed through pointer:", *ptr)
}

/*
✅ Task 2: Modifying a Value through a Pointer
Write a program that:

Declares an integer variable
Creates a pointer to that variable
Modifies the variable using the pointer
Prints the updated value
*/
func D5M2() {

	number := 42
	ptr := &number
	*ptr = 35

	fmt.Println("Value of number:", number)
	fmt.Println("Value of pointer:", *ptr)
	fmt.Println("New value of number:", number)
}

/*
✅ Task 3: Pointer Functions
Modify the function below to accept a pointer instead of a value and change the original variable.
*/
func D5M3() {
	value := 30
	updateValue(&value)
	fmt.Println("Updated Value:", value)
}

func updateValue(num *int) {
	*num = *num + 23
}

type Person struct {
	name string
	age  int
}

/*
🟢 Evening Session: Pointers and Structs
✅ Task 4: Using Pointers with Structs
Modify the following program so that it updates a struct's value using pointers.
*/
func D5M4() {

	john := Person{name: "John", age: 25}
	fmt.Println("Before birthday func:", john)
	birthday(&john)
	fmt.Println("After birthday func:", john)

}

func birthday(p *Person) {
	p.age++
}

/*
✅ Task 5: Pointer liblingsEssens
Write a program that:
Creates an liblingsEssen of integers
Uses a pointer to modify one of the elements
*/
func D5M5() {

	numliblingsEssen := [5]int{3, 4, 2, 6, 1}
	fmt.Println("liblingsEssen before:", numliblingsEssen)
	ptr := &numliblingsEssen[3]
	*ptr = 5
	fmt.Println("liblingsEssen pointer after:", numliblingsEssen)
}

/*
🎯 Final Challenge: Swapping Values using Pointers
Write a function that swaps two numbers using pointers.
*/
func D5M6(a, b *int) {
	c, d := *a, *b
	*a, *b = d, c
}

// DAY 6
/*
Task 1: Working with liblingsEssens
Create an liblingsEssen of 5 integers and assign values.
Print the first and last elements.
Change the third element and print the updated liblingsEssen.
*/
func D6E1() {
	liblingsEssen := [5]int{2, 4, 6, 8, 10}
	fmt.Println("Original liblingsEssen: ", liblingsEssen)
	fmt.Printf("First element: %v\nLast element: %v\n", liblingsEssen[0], liblingsEssen[4])
	liblingsEssen[2] = 16
	fmt.Println("Updated liblingsEssen: ", liblingsEssen)
}

/*
Task 2: Slices in Action
Create a slice of 3 favorite foods.
Append two more items to the slice.
Print the length and capacity of the slice.
*/
func D6E2() {

	liblingsEssen := []string{"Pizza", "Burger", "Beans"}
	fmt.Printf("Original Slice: %s\nOriginal Slice Length:%v\nOriginal Slice Capacity: %v\n", liblingsEssen, len(liblingsEssen), cap(liblingsEssen))
	liblingsEssen = append(liblingsEssen, "Rice", "Akpu")
	fmt.Printf("Updated Slice: %s\nUpdated Slice Length:%v\nUpdated Slice Capacity: %v\n", liblingsEssen, len(liblingsEssen), cap(liblingsEssen))
}

/*
Task 3: Iterating Over a Slice
Create a slice of 5 random numbers.
Use a loop to print each number with its index.
*/
func D6E3() {
	numbers := []int{3, 4, 2, 5, 1}
	for index, num := range numbers {
		fmt.Printf("Index: %v -> Value: %v\n", index, num)
	}
}

/*
Task 4: Extracting a Sub-Slice
Create a slice of 6 colors.
Extract the middle 3 colors into a new slice.
Print both the original and the sub-slice.
*/
func D6E4() {
	colors := []string{"Blau", "Grau", "Grün", "Rot", "Weiß", "Schwarz"}
	subColors := colors[2:5]
	fmt.Printf("Original colors: %s\n Sub-slice: %s\n", colors, subColors)
}

/*
Task 5: Understanding Slice References
Create a slice of 4 names.
Create a sub-slice of the last 2 names.
Change one name in the sub-slice.
Print the original slice—what happened? 🤔
*/
func D6E5() {
	names := []string{"Jinks", "Michael", "Ashia", "Abuashia"}
	fmt.Println(names)
	somaNames := names[2:]
	somaNames[0] = "Wanloq"
	fmt.Println(names)
}

/*
🎯 Final Challenge: Sorting a Slice
Write a program that:

Creates a slice of 6 numbers (unsorted).
Sorts the slice in ascending order.
Prints the sorted slice.
*/
func D6E6() {
	names := []string{"Jinks", "Michael", "Ashia", "Abuashia", "Wanloq", "Lord"}
	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)
}

// DAY 7
/*
📌 Create a map called studentGrades that stores
student names (keys) and their grades (values).

Add three students and their grades.
Print all students and their grades.
*/
func D7E0() {
	studentGrades := map[string]int{"Alice": 76, "Bob": 82}
	studentGrades["June"] = 88
	for student, grade := range studentGrades {
		fmt.Printf("Student Name: %v --> Grade: %v\n", student, grade)
	}
}

/*
Task 1: Create and Modify a Map
Create a map that stores country names as keys and their capital cities as values.
Add at least five countries and their capitals.
Update one of the capital cities.
Delete one country from the map.
Print the final map.
*/
func D7E1() {
	countries := map[string]string{
		"Nigeria": "Abuja",
		"USA":     "Washington",
		"Germany": "Munich",
		"China":   "Honkong",
		"France":  "Paris",
	}
	for country, capital := range countries {
		fmt.Printf("%v --> %v\n", country, capital)
	}
	countries["Germany"] = "Berlin"
	delete(countries, "France")
	fmt.Println("Countries and their capitals")
	for country, capital := range countries {
		fmt.Printf("%v --> %v\n", country, capital)
	}
}

/*
Task 2: Check for Key Existence
Write a program that asks the user to enter a country name.
Check if the country exists in the map from Task 1.
If it exists, print the capital city; otherwise, display a message saying it is not found.
*/
func D7E2() {
	countries := map[string]string{
		"Nigeria": "Abuja",
		"Usa":     "Washington",
		"Germany": "Munich",
		"China":   "Honkong",
		"France":  "Paris",
	}
	countries["Germany"] = "Berlin"
	var userInput string
	fmt.Print("Please enter a country: ")
	fmt.Scan(&userInput)
	userInput = strings.TrimSpace(userInput)
	userInput = strings.Title(strings.ToLower(userInput))
	if capital, exists := countries[userInput]; exists {
		fmt.Printf("Capital is : %v\n", capital)
	} else {
		fmt.Printf("%v Not Found\n", userInput)
	}

}

/*
Task 3: Iterating Over a Map
Create a map that stores a list of students and their grades.
Use a loop to print each student's name and grade.
*/
func D7E3() {
	studentGrades := map[string]int{
		"Alice": 76,
		"Bob":   82,
		"June":  78,
		"Clive": 83,
		"Phill": 71,
	}
	fmt.Println("Students grades table")
	for student, grade := range studentGrades {
		fmt.Printf("%v --> %v\n", student, grade)
	}

}

/*
Task 4: Word Frequency Counter
Take a sentence as input from the user.
Count how many times each word appears in the sentence.
Store the word counts in a map and print the result.
*/
func D7E4() {
	wordCount := map[string]int{}
	reader := bufio.NewReader(os.Stdin)
	counter := 0
	fmt.Print("Enter text here then press enter: ")
	sentenceString, _ := reader.ReadString('\n')
	sentenceString = strings.Title(strings.ToLower(strings.TrimSpace(sentenceString)))
	words := strings.Fields(sentenceString)
	for _, word := range words {
		if _, exists := wordCount[word]; exists {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}
	for word, count := range wordCount {
		fmt.Printf("%v\t -> %v\n", word, count)
		counter += count
	}
	fmt.Printf("Number of distinct words: %v\nTotal number of words: %v\n", len(wordCount), counter)
}

/*
Task 5: Using Maps in Functions
Write a function that takes a map of employee salaries (employee name → salary).
The function should increase each employee's salary by 10% and return the updated map.
Print the new salaries.
*/
func CurrentTask() {
	wordCount := map[string]int{}
	reader := bufio.NewReader(os.Stdin)
	counter := 0
	fmt.Print("Enter text here then press enter: ")
	sentenceString, _ := reader.ReadString('\n')
	sentenceString = strings.Title(strings.ToLower(strings.TrimSpace(sentenceString)))
	words := strings.Fields(sentenceString)
	for _, word := range words {
		if _, exists := wordCount[word]; exists {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}
	for word, count := range wordCount {
		fmt.Printf("%v\t -> %v\n", word, count)
		counter += count
	}
	fmt.Printf("Number of distinct words: %v\nTotal number of words: %v\n", len(wordCount), counter)
}
