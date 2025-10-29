//method1: declaring multiple variables in one line
package main
import "fmt"
func main() {
  var s, t string = "foo", "bar" //assigning multiple variables in a line
  fmt.Println(s)
  fmt.Println(t)
}

//method 2: variables of different data types; another way to declare
package main
import "fmt"
func main() {
  var (
  s string = "foo"
  i int = 5)
  fmt.Println(s)
  fmt.Println(i)
}

//Short variable declaration
s := "Hello World"
///
package main
import "fmt"
func main() {
  name := "Lisa"
  name = "Peter"
  fmt.Println(name)
}
//output: Peter
