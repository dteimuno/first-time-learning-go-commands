// single input
package main
import "fmt"
func main() {
  var name string //set variable that input will be equal to
  fmt.Print("Enter your name: ") // set input query
  fmt.Scanf("%s", &name) //ampersand sign needs to be used with Scanf function
  fmt.Println("Hey there, ", name)
}
---
//multiple inputs
package main
import "fmt"
func main() {
  var name string //define variable 1
  var is_muggle bool //define variable 2
  fmt.Print("Enter your name and are you a muggle: ") //input ask message
  fmt.Scanf("%s %f", &name, &is_muggle) //format specifiers for string and bool and 
}
---
package main
import "fmt"
func main() {
  var a string
  var b int
  fmt.Print("Enter a string and a number: ")
  count, err := fmt.Scanf("%s %d", &a, &b)
  fmt.Println("count : ", count)
  fmt.Println("error : ", err)
  fmt.Println("a : ", a)
  fmt.Println("b : ", b)
  
  
}
