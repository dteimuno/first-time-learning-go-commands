/* for loop syntax:
for <initialization>; <condition>; <post> {
  //statements
}
*/

/* example:
for i := 1; i <= 3; i++ {
  fmt.Println("Hello World")
}
*/

---
// example in action:
package main
import "fmt"
func main() {
  for i := 1; i <= 5; i++ {
    fmt.Println(i*1)
  }  
}
/* output:
1
4
9
16
25
*/

---
// you can also skip the initialization and post statement
package main
import "fmt"
func main() {
  i := 1
  for i <= 5 {
    fmt.Println(i * i)
    i += 1
  }
}
/* output:
1
4
9
16
25
*/

---
// infinite loop
package main
import "fmt"
func main() {
  sum := 0
  for {
    sum++ //repeated forever
  }
  fmt.Println(sum) //never reached 
}


---
// break on certain iterations in a loop and then continue, use "break" statement. The break statement ends the loop immediately when it is encountered:
package main
import "fmt"
func main() {
  for i := 1; i <= 5; i++ {
    if i == 3
    break
  }
  fmt.Println(i)
}

/* output:
1
2
*/

---
// the "continue" continues after skipping a certain iteration in for-loop:
package main
import "fmt"
func main() {
  for i := 1; i <= 5; i++ {
    for i == 3 {
        continue
    }
    fmt.Println(i)
  }
}
/* output:
1
2
4
5
*/
