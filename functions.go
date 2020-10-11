package main
import (
	"fmt"
	"strings"
	)

func main(){
	
	//basic function
	var product int = multiply(3, 4)
	fmt.Println(product)
	
	//recursion
	var fac int = factorial(5)
	fmt.Println(fac) //returns 120, or 5!
	
	//sends in arguments of different types
	multiType(2, "hi")
	
	//multiple return values 
	var squareOne, squareTwo int = squares(3, 4)
	fmt.Println(squareOne) //returns 9 (3 squared)
	fmt.Println(squareTwo) //returns 16 (4 squared)
	
	//multiple return values
	var add, mult int = returns(5, 8)
	fmt.Println(add) //returns 13 
	fmt.Println(mult) //returns 40
	
	//pass by value
	var x int = 10
	var y int = 15
	fmt.Println("The values before the function is called:")
	fmt.Println("x = ", x, "\ny = ", y)
	swap(x,y)
	fmt.Println("\nThe values after the function is called:")
	fmt.Println("x = ", x, "\ny = ", y)
	
	//splitting strings
	var word1, word2 string = splitting("hi hello")
	fmt.Println("word one is: ",word1)
	fmt.Println("word two is: ",word2)	
}

//standard function
//multiplies two numbers & returns the value
func multiply(x int, y int) int {
	return x*y
}

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


//multiple returns
func squares(x int, y int) (int, int){
	return x*x, y*y
}

//multiple returns
//return types listed as variables
func returns(x int, y int) (addition int, multiplication int){
	addition = x+y
	multiplication = x*y
	return
}

//recursion
//from tutorialspoint
func factorial(i int) int {
	if(i <= 1){
		return 1
	}
	return i * factorial(i-1)
}

//pass by value
//from geeksforgeeks
func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

//splitting strings
func splitting(x string) (string, string) {
	var s []string = strings.Split(x, " ")
	return s[0], s[1]
}








