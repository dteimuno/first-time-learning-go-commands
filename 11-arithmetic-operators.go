// addition +
package main
import "fmt"
func main() {
  var a, b string = "foo", "bar"
  fmt.Println(a + b)

}
//output: foobar

---
// subtraction -
package main
import "fmt"
func main() {
  var a,b string = "foo", "bar"
  fmt.Println(a - b)
}
//output: invalid operator. a- b (operator not defined on string)
---
// another subtraction
package main
import "fmt"
func main() {
  var a, b float64 = 79.02, 75.66
  fmt.Println("%.2f", a -b)
}
---
// increment operator ++
// this a unary operator: acts on only one operand
package main
import "fmt"
func main() {
  var i int = 1
  i++
  fmt.Println(i)
}
// output: 2

---
// decrement operator
package main
import "fmt"
func main() {
  var i int = 1
  i--
  fmt.Println(i)
}
//output: 0
