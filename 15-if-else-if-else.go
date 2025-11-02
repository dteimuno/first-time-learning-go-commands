// if statement
package main
import "fmt"
func main() {
  var a string = "happy"
  if a == "happy" {
    fmt.Println(a)
  }
}
//output: "happy"

---
// is and else:
package main
import "fmt"
func main() {
  var fruit string = "grapes"
  if fruit == "apples" {
    fmt.Println("Fruit is apple")
  } else {
    fmt.Println("Fruit is not apple")
  }
}

// if, multiple else is, and else in one:
package main
import "fmt"
func main() {
  fruit := "grapes"
  if fruit == "apple" {
    fmt.Println("I love apples")
  } else if fruit == "orange" {
    fmt.Println("Oranges are not apples")
  } else {
    fmt.Println("no appetite")
  }
}
