/*
comparison operators in golang are:
== equal to
!= not equals
< less than
<= less than or equal to
> greater than
>= greater than or equal
*/
---
// equal to (==)
package main
import "fmt"
func main() {
  var city string = "Kolkata"
  var city_2 string = "Calcutta"
  fmt.Println(city == city_2)
}
//output: false

---
// not equal to (!=)
package main
import "fmt"
func main() {
  var city string = "Kolkata"
  var city_2 string = "Calcutta"
  fmt.Println(city != city_2)
}
//output: true

---
// less than
package main
import "fmt"
func main() {
  var a,b int = 5, 10
  fmt.Println(a < b)
}
//output: True

---



