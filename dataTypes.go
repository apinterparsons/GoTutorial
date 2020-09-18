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