// sorta like if, else if, else but here the expression is tested against multiple cases until the right one is picked. If none match the expression it will pick a default.
package main
import "fmt"
func main() {
  var i int = 100
  switch i {
    case 10: 
        fmt.Println("i is 10")
    case 100, 200:
        fmt.Println("i is either 100 or 200")
    default:
        fmt.Println("i is neither 0, 100 or 200")
  }
}
//output: i is either 100 or 200

---
package main
import "fmt"
func main() {
  var i int = 800
    switch i {
        case 10:
            fmt.Println("i is 10")
        case 100, 200:
            fmt.Println("i is either 100 or 200")
        default:
            fmt.Println("i is neither 0, 100 or 200")
        
  }
}
//Output: i is neither 0, 100 or 200

---

/* fallthrough:
The fallthrough keyword is used to switch-case to force the execution flow to fall through the successive case block
*/
package main
import "fmt"
func main() {
  var i int = 10
  switch i {
      case -5: 
          fmt.Println("-5")
      case 10:
          fmt.Println("10")
          fallthrough
      case 20:
          fmt.Println("20")
          fallthrough 
      default:
          fmt.Println("default")
  }
   
}
 /* output:
  10
  20
  default
  */
 
---
// switch with conditions
package main
import "fmt"
func main() {
  var a, b int = 10, 20
  switch {
      case a+b == 30:
          fmt.Println("equal to 30")
      case a=b <= 30:
          fmt.Println("less than or equal to 30")
      default:
          fmt.Println("greater than 30")
  }
}
//output: equal to 30

---



