# Functions and Parameters
##### a more in depth look at the code can be found in functions.go

### Creating a Function
A function in Go can take zero or more arguments. If there is more than one argument, the parameters are separated by commas. The type of variable is declared after the variable name. If a function returns a value, the type is listed after the parameters, but before the opening curly brace. To call a function, type the name(argument), much like any other language.

In terms of the naming conventions for a Go function, the rules applied to variables carry over. A function must begin with a letter and cannot have any spaces. If it starts with an uppercase letter, the function will be exported to other packages. In the case of a multi-word function, camelCase should be used. 

There are also no rules regarding where in a code file a function has to be placed in order to run. The program will automatically start by running the main function no matter where in the file it is located. 
```
var product int = multiply(3, 4)

func multiply(x int, y int) int {
	return x*y
}

```

### Recursion
It is possible to create a recursive function in Go. That means that a function declares itself until it reaches an exit condition that requires it to stop. For example, the code below shows how a recursive function can determine a factorial. 

```
func factorial(i int) int {
	if(i <= 1){
		return 1
	}
	return i * factorial(i-1)
}
```

### Accepting Multiple Parameters
Functions can accept multiple parameters, so long as they are separated by commas. If multiple parameters are the same type you can omit the type declaration until the final parameter. If the parameters are not the same type, you must declare the type following the name of the variable. 

```
//parameters of like type
//can exclude int declaration until final type
func add(x, y, z int, b string) int {
	return x+y+z
}


//parameters of different types
func multiType(x int, y string){
	fmt.Println("the int is ", x)
	fmt.Println("the string is", y)
}
```

### Returning Multiple Values
Go supports the returning of multiple values from a function. When you are writing the return type in the function declaration, you must list the types of values you are expecting. For example, if the function is going to return three strings, you would list (string, string, string) after the parameters. 
Everything that is being returned is in the same return statement separated by commas. 
If the values are to be saved, the variable declarations are listed in one line separated by commas. They are saved in the order that they are listed. (The first returned value is saved in the first declared variable.)

```
var squareOne, squareTwo int = squares(3, 4)

//multiple returns
func squares(x int, y int) (int, int){
	return x*x, y*y
}
```
However, you can also declare a variable name when stating what types the function is to be returning. In that case, the return statement would be empty and the variables would be assigned values within the function. When the function is called, however, the variables must still be declared in one line, and the values are returned in the order that they are assigned values within the function. 
```
var add, mult int = returns(5, 8)

//multiple returns
//return types listed as variables
func returns(x int, y int) (addition int, multiplication int){
	addition = x+y
	multiplication = x*y
	return
}
```

### Passing Parameters
Go is not pass by reference and is instead a language that passes parameters by value. This means that every time a variable is passed as a parameter, a new copy of the variable is created at a new memory location and then passed. This also means that a swap function will not be able to change the values of variables that are used as parameters. 

### Passing Functions
In Go, you can pass a function into another function. It is also possible to have a function be declared as a return type. In such cases, the function declaration may look complicated, but all the components are the same as a basic function call and declaration. 

### Sources
https://tour.golang.org/basics/4

https://blog.golang.org/declaration-syntax	

http://golangprograms.com/naming-conventions-for-golang-functions.html

https://gobyexample.com/functions

https://www.geeksforgeeks.org/function-returning-multiple-values-in-go-language/#:~:text=In%20Go%20language%2C%20you%20are,defined%20in%20the%20parameter%20list.

https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/#:~:text=Strictly%20speaking%2C%20there%20is%20only,to%20called%20function%20or%20method.&text=In%20case%20a%20variable%20is,same%20memory%20address%20is%20created.

https://gobyexample.com/recursion

https://www.tutorialspoint.com/go/go_recursion.htm	

https://www.geeksforgeeks.org/function-arguments-in-golang/




