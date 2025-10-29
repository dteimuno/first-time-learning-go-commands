//method 1
package main
import "fmt"
func main() {
  fmt.Print("Hello World")
}

//method 2- printing a variable
package main
import "fmt"
func main() {
  var city string = "Kolkata"
  fmt.Print(city)
}

///method 3- print variable and string together
package main
import "fmt"
func main() {
  var name string = "Kodekloud"
  var user string = "Harry"
  fmt.Print("Welcome to ", name, ", ", user)
}

//method4: printing on newline
package main
import "fmt"
func main() {
  var user string = "Kodekloud"
  var user string = "Harry"
  fmt.Print(name, "\n")
  fmt.Print(user)
}

//method5: an even better way to do method4 of printing a new line:
package main
import "fmt"
func main() {
  var user string = "Kodekloud"
  var user string = "Harry"
  fmt.Println(name)
  fmt.Println(user)
}

//method 6: Printf- string formatting
//Printf- format specifier
/*
%v - formats the value in a default format
so, example:
*/
var name string = "Kodekloud"
fmt.Printf("Nice to see you here, at %v", name) //the name variable replaces %v
/*Output*/: Nice to see you here, at Kodekloud

//Printf-format specifier %d
// %d - formats decimal integers
var grades int = 42
fmt.Printf("Marks: %d", grades)
//output>> Marks: 42


//add format specifiers together
package main
import "fmt"
func main() {
  var name string = "Joe"
  var i int = 78
  fmt.Printf("Hey, %v! You have scored %d/100 in Physics", name, i)
}
