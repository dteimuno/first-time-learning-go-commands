//type casting: changing variable from one type to another

//integer to float
package main
import "fmt"

func main() {
  var i int = 64
  var f float64 = float64(i)
  fmt.Printf(".2f\n", f)
}

---
// convert float to integer
package main
import "fmt"
func main() {
  var f float64 = 45.89
  var i int = int(f)
  fmt.Printf("%v\n", i)
}
//output: 45

---
//strconv package that can be used to convert string to integer and vice versa
/*
Itoa() - converts integer to string ; returns one value - string formed with given integer
*/

package main
import {
  "fmt"
  "strconv"
}
func main() {
  var i int = 32
  var s string = strconv.Itoa(i) //convert integer to string
  fmt.Printf("%q", s) 
}
//output: "42"

---
/* method 2
Atoi() - converts string to integer - returns two values: the corresponding integer and error, if any
*/
package main
import {
  "fmt"
  "strconv"
}
func main() {
  var s string = "200"
  i, err := strconv.Atoi(s)
  fmt.Printf("%v, %T \n", i, i)
  fmt.Printf("%v, %T \n", err, err)
  
}
/* Output:
200, int
<nil>, <nil>
*/

---
// string to integer
package main
import {
  "fmt"
  "strconv"
}

func main() {
  var s string = "200abc"
  i, err := strconv.Atoi(s)
  fmt.Printf("%v, %T \n", i, i)
  fmt.Printf("%v, %T \n", err, err)
  
}
/* Output:
0, int
strconv.Atoi: parsing "200a": invalid syntax, *strconv.NumError
*/
