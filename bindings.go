package main
import ("fmt")

var y int = 7

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

	scopeTest()
	fmt.Println(x) //prints 6

	x = returnTest()
	fmt.Println(x) //prints 20

	//showing that Go is pass by value
	a := [3]string{"c","a","t"}
	b := [3]string{"d","o","g"}
	a = b
	b[1] = "u"
	fmt.Println(a) //prints dog
	fmt.Println(b) //prints dug
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
