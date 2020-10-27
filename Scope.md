# Naming, Scope, & Binding

### Same Name Within Main
If you declare a variable in the main function and then another of the same name within a for loop, you will end up having two variables of the same name. Go will allow you to perform such an action without conflict, but it could very easily get confusing for a programmer to determine what is going on. The value of x within the for loop is only true as long as you are still within the loop. As soon as you exit out, the value of x goes back to what it had been declared as in main.
 
 ```
func main(){

	var x int = 5
	x = x + 1
	fmt.Println(x) //prints 6

	for i:=0; i<2; i++ {
		var x int = 10
		x = x + 1
		fmt.Println(x) //prints 11
	}

	fmt.Println(x) //prints 6
  }
```
 
 ### Same Name in a Function
 If you declare a variable within main and then create a function that declares a variable with the same name, but a different value, the results will mirror what happened in the previous example and you will end up with two variables of the same name. Within the function, the variable will be equal to a particular value, but as soon as you return to main, the variable has its original value. However, if the returned value of a function is set equal to that variable, the value of the variable will change. 
 
 ```
 func main(){
 	fmt.Println(x) //prints 6

	scopeTest()
	fmt.Println(x) //prints 6

	x = returnTest()
	fmt.Println(x) //prints 20
  
}

func scopeTest(){
	var x int = 20
	x = x + 1
	fmt.Println(x) //prints 21
}

func returnTest()(int){
	var x int = 20
	return x
}

```

### Global Variables
Go does allow the use of global variables. They must be declared outside of all functions, including main. If a global variable and a local variable have the same name, the local variable will take priority and the program will use that value. 

### Pass by Value
Go does not have any pass by reference and instead relies on pass by value. There is no distinction between different kinds of variables in this regard. 

The code below shows two arrays that are filled with different values. They are then set equal to each other, a value in one of the arrays is altered, and both are printed. When they are printed, only one of the arrays reflects the change that had been made. This means that Go uses pass by value in order to deal with variables. 

```
	a := [3]string{"c","a","t"}
	b := [3]string{"d","o","g"}
	a = b
	b[1] = "u"
	fmt.Println(a) //prints dog
	fmt.Println(b) //prints dug
```

### Sources
https://www.tutorialspoint.com/go/go_scope_rules.htm

https://kuree.gitbooks.io/the-go-programming-language-report/content/26/text.html

