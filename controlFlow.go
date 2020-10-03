package main
import "fmt"

func main(){
	//one-condition if/else
	if 3%2 == 0 {
		fmt.Println("3 is even")
	} else {
		fmt.Println("3 is odd")
	}
	
	//multi-condition if/else
	//variable declaration within if statement
	if x := 1; x > 0 && x < 2 {
		fmt.Println("x equals one")
	} else {
		fmt.Println("x is not equal to one")
	}	
	
	//short circuit
	//if: since the first expression is false, the second will not be evaluated 
	//else if: the first expression is true, so the second one will not be evaluated
	if 2+2 == 5 && 3+3 == 6 {
		fmt.Println("should not print")
	} else if 3+3==6 || 2+2==5 {
		fmt.Println("should print")
	} else {
		fmt.Println("shouldn't print")
	}
	
	//for loop
	//init statement: x:=0
	//condition expression: x<5
	//post statement: x++
	for x := 0; x < 5; x++ {
		fmt.Println("this phrase will print five times!")
	}
	
	//alternate for loops
	
	//condition-only for loop
	//acts similarly to a while loop by omitting init statement & post statement
	var y = 0
	for ; y < 5; {
		fmt.Println("this loop will iterate five times")
		y++
	}
	
	//no condition for loop
	//will always be viewed as true
	//have to use an if condition & break statement to stop it
	for z := 0; ; z++ {
		fmt.Println("the loop will go until you make it stop")
		if z > 5{
			break
		}
	}
	
	//break statement
	for x := 0; x < 10; x++ {
		if x == 6{
			fmt.Println("x is equal to 6 and the loop is ending.")
			break
		}
	fmt.Println("the loop is still going!")
	}
	
	//continue statement
	for x := 0; x < 10; x++ {
		if x == 6{
			fmt.Println("x is equal to six, but this won't stop the loop")
			continue
		}
		fmt.Println("this won't print when the continue is executed")
	
	}
	
	//switch
	x := 2+2
	switch x{
	case 3:
		fmt.Println("2 plus 2 is 3")
	case 4:
		fmt.Println("2 plus 2 is 4") //this will print
	case 5:
		fmt.Println("2 plus 2 is 4")
	}


	
	


}