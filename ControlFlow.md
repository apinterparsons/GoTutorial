# Go Control Flow

### If/Else Statements
In Go, you are able to use an if/else statement or an if/else if/else statement. There are no conditional statements other than "if" in Go. If you are using multiple conditional statements (ex. an if and an else), the else statement must be on the same line as the previous curly bracket. In addition, no parentheses are required around the expression(s) in an if statement.   

```
 if 3%2 == 0 {
  
	fmt.Println("3 is even")
    
} else {
  
	fmt.Println("3 is odd")
    
}
```
  
The if statements can have one or more conditions. With multi-condition if statements, use the following operators to combine expressions:

and: &&

or: ||

not: !

Variables can also be declared within an if statement, and the variable's scope is within the if/elses.

```
if x := 1; x > 0 && x < 2 {

	fmt.Println("x equals one")
  
} else {

	fmt.Println("x is not equal to one")
  
}	
```

### Short Circuit Evaluation
Go uses short-circuit evaluation with the "and" and "or" (&& and ||) statements. If there is an "or" statement and the first expression is true, the program will not even check the second expression. A similar idea goes with "and" statements; if the first expression is false, the second one will not be checked. 

```
if 2+2 == 5 && 3+3 == 6 {
	
  fmt.Println("should not print")

} else if 3+3==6 || 2+2==5 {
	
  fmt.Println("should print")

} else {
	
  fmt.Println("shouldn't print")

}
```

### Dangling Else
Go deals with the dangling else problem by making sure than any else if or else statement begins on the same line as the closing curly bracket from the previous statement. If an else is on its own line, you will get a syntax error regarding an unexpected else. The same error will occur if an else if statement begins on a new line. 

### Loops
Go only supports a for loop. Every for loop requires an initiating variable that is executed before the first iteration of the loop. Oftentimes, this is a short declaration of a variable like x:=0. It must also have a conditional expression that is tested before each iteration of the loop. The loop will continue if the expression is true, but will stop as soon as the expression becomes false. Finally, the loop must have a post-statement that is executed at the end of each iteration of the loop. Oftentimes this is something as simple as x++. As for syntax, parentheses are not required, but curly brackets are. 

There are also ways to manipulate for loops by omitting one or more parts of the statement to create, for example, a condition-only for loop that represents a while loop in other languages. There is also a way to omit just the condition expression so that the loop always evaluates true until a break statement is used. However, shown below is a simple for loop. 

```
for x := 0; x < 5; x++ {

	fmt.Println("this phrase will print five times!")
    
}
```
  
### Break & Continue Statements
Go supports both break and continue statements. A break statement stops the execution of a loop and is almost always paired with a conditional if statement. This will help stop the loop before it is supposed to finish. However, it will only stop the innermost loop it is in. For example, multiple break statements may be needed for nested loops. 

```
for x := 0; x < 10; x++ {
  
	if x == 6{
    
		fmt.Println("x is equal to 6 and the loop is ending.")
      
		break
      
	}
    
  fmt.Println("the loop is still going!")
  
}
```
  
The continue statement is when you don't want to completely break out of a loop, but you want to halt the current iteration and start a new one. Just like the break statement, this is commonly used with a conditional if statement. 

```
	for x := 0; x < 10; x++ {
  
		if x == 6{
    
			fmt.Println("x is equal to six, but this won't stop the loop")
      
			continue
      
		}
    
		fmt.Println("this won't print when the continue is executed")
	
	}
```
  
### Switch Statements
Go supports switch statements. A break statement is not required to get out of them. A switch statement is contained within a set of curly brackets and each case statement is set off with a colon. You cannot use a continue statement because they are used exclusively in loops. If you try to use one in a switch statement, you will get an error. If a break statement is used, the switch statement will not be executed. However, Go does support a fallthrough statement that makes multiple cases eligible to be executed. This statement must be put as the last line of a case, much like you would use a break or continue statement. 

```
	x := 2+2
  
	switch x{
  
	case 3:
  
		fmt.Println("2 plus 2 is 3")
    
	case 4:
  
		fmt.Println("2 plus 2 is 4") //this will print
    
	case 5:
  
		fmt.Println("2 plus 2 is 4")
    
	}
```
  
### Goto Statements
Go will allow the use of a goto statement which can be declared using a label. This will make the execution of the program jump around to where the label is declared and can mess with the normal flow of program control. To declare a label, the label name followed by a colon. In addition, a label must be declared within a function. However, using a goto statement can easily make your code more complex than it needs to be and the use of them is highly discouraged. 

### Sources
https://gobyexample.com/if-else

https://tour.golang.org/flowcontrol/1

https://gobyexample.com/switch

https://www.callicoder.com/golang-control-flow/

https://golang.org/ref/spec#Logical_operators

https://kuree.gitbooks.io/the-go-programming-language-report/content/22/text.html

https://www.digitalocean.com/community/tutorials/using-break-and-continue-statements-when-working-with-loops-in-go

https://stackoverflow.com/questions/40821855/do-go-switch-cases-fallthrough-or-not

https://go101.org/article/control-flows.html









