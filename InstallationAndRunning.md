# Installing Go and Running a Hello, World

### Installation
Go can be installed on any operating system at golang.org. In order to install the language, go to the previously mentioned website and click on the Download Go button on the left side of the screen. That will bring you to a Downloads page where you can select the version of Go that you would like to download based on your operating system. After that, the website recommends that you open a Command Window and check to make sure that the proper version of Go was successfully installed. 

### Running a Program
In order to write code for Go, you can use a simple text editor. While any can work, the Go website recommends using VSCode, Vim, or GoLand as they seem to be the most popular editors for Go users. I'm choosing to use Notepad++ because that is what is already installed on my device, but I may switch to one of the recommended editors that provides more support for Go. 
In order to run a program, you can use either a PowerShell or Command Window. Navigate to the directory in which you have saved the file you want to run. Then, type in **go run [filename].go** and hit enter. 


### Getting Started
To start a program, there is a little bit of boiler-plate code that you can include. For a simple Hello, World, you should add the following to the beginning of the program:
package main  
import "fmt"   
This is declaring a main package which is a way that Go groups functions. It also imports the fmt package that is a part of Go's standard library and helps format text and print to the console. When creating a Hello, World you must also have a main function. 
In order to write a comment, you use two forward slashes //. You can also use /* and */ to create block comments. 


### Hello World
Overall, your code should look something like this:

**package main**

**import "fmt"**


**func main(){**

	//this will print hello, world to the console
  
	fmt.Println("Hello, World")
  
**}**

### Sources
https://golang.org/dl/

https://golang.org/doc/tutorial/getting-started

https://www.digitalocean.com/community/tutorials/how-to-write-comments-in-go
