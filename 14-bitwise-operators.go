//bitwise operator work at bit level

/* Golang bitwise operators:
& - bitwise AND
| - bitwise OR
^ - bitwise XOR
>> - right shift
<< - left shift
*/

---
// bitwise AND(&)
/* takes two numbers as operands and does AND on every bit of two numbers
12 - 00001100
*/
---
//bitwise AND
package main
import "fmt"
func main() {
  var x, y int = 12, 25
  z := x & y
  fmt.Println(z)
}
// output: 8

---
//bitwise OR (||)
package main
import "fmt"
func main() {
  var x, y int = 12, 25
  z := x | y
  fmt.Println(z)
}
// output: 29

---
//bitwise XOR (^)
// takes two numbers as operands and does XOR on every bit of two numbers. Result of XOR is 1 if the two bits are opposite; and zero when both bits are same
package main
import "fmt"
func main() {
  var x, y int = 12, 25
  z := x ^ y
  fmt.Println(z)
}
//output: 21

---
// leftshift operator (<<)
// shifts all bits left by a certain number of specified bits. 
package main
import "fmt"
func main() {
  var x int = 212
  z := x << 1
  fmt.Println(z)
}
// output: 424
---

// rightshift operator (>>)
// shifts all bits right by a certain number of specified bits. 
package main
import "fmt"
func main() {
  var x int = 212
  z := x >> 2
  fmt.Println(z)
}

