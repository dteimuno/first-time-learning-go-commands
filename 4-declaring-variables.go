/* how to store variable type
var <variable name> <data type> = <value>
e.g var s string = "Hello World!"
*/

// integer variable
var i int = 100

//boolean variable
var b bool = false //boolean variable can only be true or false

var f float64 = 77.90

// assuming a main.go file, this is how we'll declare variable:
package main
import "fmt"
func main() {
  var greeting string = "Hello World"
  fmt.Println(greeting)
}
