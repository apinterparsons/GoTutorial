# Go Data Types and Variables

### Variables 

The requirements for naming variables in Go include starting with a letter or underscore and is not one of the 25 predetermined keywords. These keywords include break, case, chan, etc. and can be found at https://golang.org/ref/spec#Keywords. In addition, variables are case sensitive, so it is vital to stay consistent.

The conventions in Go for naming variables include using camelCase instead of underscores for multiple words. In addition, Go programmers tend to favor shorter names especially when it comes to local variables. If a variable has a large scope, longer and more descriptive names can be used.

### Type of Language
Go is a strongly, statically typed language. 

*strongly*: The compiler will help catch any errors that are present within the program. 

*statically*: You do not necessarily have to declare the type that a variable is when it is created to allow the programmer to declare multiple variables of different types in the same line. When you go to declare a variable, the format var name type = expression is used, though you can omit type or expression (but not both). 

### Possible Errors
Mixing types within a variable will lead to a compiling error. For example:

var x = "5" + 6

fmt.Println(x)

This will return an invalid operation errors stating that mismatched types cannot be combined. If something like this is to be compiled, you would need to make them both the same type. If they are strings ("5" + "6") they will be concatenated and the program will return 56. If they are both ints, integer math will be performed and the program will return 11. 

### Limitations
In Go, you cannot add variables of different data types, as shown above. However, this extends beyond just strings and ints. Go has a variety of different types of ints (ex. int8 and int16) that are based off of the size of the variable. An int8 and an int16 cannot be added together as they are considered different types even though they are both a kind of int.  

Furthermore, arrays in Go must be the same data type. To declare an array, you would use the structure var variableName [SIZE] variableType. 

### Conversions
Go requires an explicit conversion between data types. When converting from a float to an int, Go will be narrowing and get rid of the numbers after the decimal point and will round up or down depending on the numbers following the decimal point. (However, it should be noted that 10.9899999 gets rounded to 10 while 10.99999 gets round to 11.) If you want to convert a string to an int, you can use the strconv package that is within Go's standard library. 

### Example
package main

import "fmt"

func main(){

	//four ways to declare variables (shown through ints):
	var w int 
	w = 6
	var x int = 6
	var y = 6
	z := 6 //this is only allowed within a function 
	fmt.Println("adding ints: ", w+x+y+z) //this will return 24
	
	//strings
	var greeting string = "Hello!"
	fmt.Println("string: ",greeting)
	
	//floating-point numbers
	var fp1 float32 = 10.88
	var fp2 float32 = 11.443
	fmt.Println("adding floats: ",fp1 + fp2)
	
	//boolean
	var tf bool = true
	fmt.Println("boolean: ", tf)
	
	//array
	var num = [5]int{2, 5, 77} //will put a 0 if index is not defined
	fmt.Println("an array with zero values: ", num)
	num[4] = 6 //puts 6 at 4th index; 3rd index still 0
	num[3] = 33
	fmt.Println("an array with no zero values: ", num)
	
	//map
	var stateCapitals map[string]string
	stateCapitals = make(map[string]string)
	
	stateCapitals["Massachusetts"] = "Boston"
	stateCapitals["Alabama"] = "Montgomery"
	stateCapitals["Alaska"] = "Juneau"
	stateCapitals["Arizona"] = "Phoenix"
	stateCapitals["Arkansas"] = "Little Rock"
	stateCapitals["California"] = "Sacramento"
	
	for state := range stateCapitals{
		fmt.Println(stateCapitals[state], " is the capital of ", state)
	}
	
	}
  
  ### Sources 
*The Go Programming Language* by Alan A. A. Donovan & Brian W. Kernighan

https://tour.golang.org/basics/13

https://golang.org/ref/spec

https://www.tutorialspoint.com/go/index.htm 

