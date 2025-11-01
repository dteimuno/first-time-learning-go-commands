// in golang the logical operators are:
/*
logical AND: (&&)
logical OR: (||)
logical NOT: (!)
*/

//logical AND
package main
import "fmt"
func main() {
  var x int = 10
  fmt.Println((x < 10) && (x < 200))
  fmt.Println((x < 300) && (x < 0))

  
}
//output: TRUE
//output2: FALSE
---
// OR logical operator:
// returns true if one of the variables is true; returns false if none is true
package main
import "fmt"
func main() {
  var x int = 10
  fmt.Println((x < 0) || (x < 200))
  fmt.Println((x < 0) || (x > 200))

}
/* output:
true
false
*/


---
/* NOT logical operator: unary operator
reverses the result, returns false if expression evaluates to true and vice versa
*/
package main
import "fmt"
func main() {
  var x, y int = 10, 20
  fmt.Println(!(x > y))
  fmt.Println(!(true))
  fmt.Println(!(false))

}
/* output:
true
false
true
*/

