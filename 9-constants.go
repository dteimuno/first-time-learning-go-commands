// Two types of constants: typed and untyped
// Untyped constant: constants are untyped unless they are given a type at declaration; allow for flexibility
/* e.g */ const age = 12 //no data type listed

//typed constants: data type is explciitly mentioned; the flexibility that comes with untyped is lost
const name string = "Harry"
const age int = 12

---
//understanding constants
package main
import "fmt"
func main() {
  const name = "Harry Potter"
  const is_muggle = false
  const age = 12
  fmt.Printf("%v: %T \n", name, name)
  fmt.Printf("%v: %T \n", is_muggle, is_muggle)
  fmt.Printf("%v: %T \n", age, age)


}
/* Outputs:
Harry Potter: string
false: bool
12: int
*/

---
package main
import "fmt"
func main() {
  const name = "Harry Potter"
  name = "Hermione Granger"
  fmt.Printf("%v: %T\n", name, name)
}
// Output: errror: cannot assign to name
---
// you cannot declare a constant and not initialize a value
package main
import "fmt"
func main() {
  const name //will cause error because of no constants need an assignment in golang
  name = "Hermione Granger"
  fmt.Printf("%v: %T\n", name, name)
}
//you will get errors
---
// short hand variable operator does not work for constants in golang
package main
import "fmt"
func main() {
  const name := "Hermione Granger" // variable assginment operator does not work in golang for constants
  fmt.Printf("%v: %T\n", name, name)
}
// expected error: syntax error, unexpected :=, expecting =

---
package main
import "fmt"
const PI float64 = 3.14
func main() {
  var radius float64 = 5.0
  var area float64
  area = PI * radius * radius
  fmt.Println("Area of Circle is: ", area)
}
//output: 78.5
---


