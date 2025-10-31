/* To find the type of variable being used, you can use the following format specifiers:
%T format specifier
reflect.TypeOf function from the reflect package
*/

// using %T
package main
import "fmt"
func main() {
  var grades int = 42
  var message string = "hello world"
  var isCheck bool = true
  var amount float32 = 5466.54
  fmt.Printf("variable grades = '%v' is of type %T \n", grades, grades)
  fmt.Printf("variable message = '%v' is of type %T \n", message, message)
  fmt.Printf("variable isCheck = '%v' is of type %T \n", isCheck, isCheck)
  fmt.Printf("variable amount = '%v' is of type %T \n", amount, amount)  
}
/* Outputs:
go run main.go:
variable grades = 42 is of type int
variable message = "hello world" is of type string
variable isCheck = true is of type bool
variable amount = 5466.54 is of type float32
*/
---

//Using reflect.TypeOf
package main
import {
  "fmt"
  "reflect"
}
func main() {
  fmt.Printf("Type %v \n", reflect.TypeOf(1000))
  fmt.Printf("Type %v \n", reflect.TypeOf("priyanka"))
  fmt.Printf("Type %v \n", reflect.TypeOf(46.0))
  fmt.Printf("Type %v \n", reflect.TypeOf(true))

}
/* Output:
Type: int
Type: string
Type: float
Type: bool
*/

---
// Using reflect.TypeOf with variables
package main
import {
  "fmt"
  "reflect"
}
func main() {
  var grades int = 42
  var message string = "hello world"
  fmt.Printf("variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
  fmt.Printf("variable message=%v is of type %v \n", grades, reflect.TypeOf(message))
  
}
/*Output:
variable grades=42 is of type int
variable message='hello world' is of type string
*/


