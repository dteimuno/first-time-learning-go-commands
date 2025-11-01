/* assignment operators
assign =
add and assign +=
subtract and assign -=
multiply and assing *=
divide and assign /=
*/
---
//add and assign operator 
// assigns left operand with the additional result 
// for example: X +=y is x= x+y
package main
import "fmt"
func main() {
  var x, y int = 10, 20
  x += y
  fmt.Println(x)
}
// result: 30

---
//subtract and assign: assigns left operand with the subtraction result
// for example: X -=y is x= x-y

package main
import "fmt"
func main() {
  var x, y int = 10, 20
  x -= y
  fmt.Println(x)
}
//output: -10

---
//multiple and assign: assigns left operand with the multiplication result
// for example: X *=y is x= x*y

package main
import "fmt"
func main() {
  var x, y int = 10, 20
  x *= y
  fmt.Println(x)
}
//output: 200

---
//divide and assign: assigns left operand with the division result
// for example: X /=y is x= x/y

package main
import "fmt"
func main() {
  var x, y int = 10, 20
  x /= y
  fmt.Println(x)
}
//output: 0.5

---
//divide and assign modulus: assigns left operand with the modulus result
// for example: X %=y is x= x%y

package main
import "fmt"
func main() {
  var x, y int = 210, 20
  x %= y
  fmt.Println(x)
}
//output: 10
