/* arrays: fixed length
- should be of same data type

array declaration syntax:
var <array name> [size of array] <data type>
*/
var grades [5] int
var fruits [3] string 

---
// array declaration example:
package main
import "fmt"
func main() {
  var grades [5] int 
  fmt.Println(grades)
}
//output: [0 0 0 0 0]
---
// array declaration example with empty string and integer
package main
import "fmt"
func main() {
  var grades [5] int 
  var fruits [3] string
  fmt.Println(grades)
  fmt.Println(fruits)
}
/*output: 
[0 0 0 0 0]
[ ]
*/

---
// array initialization methods
var grades [3]int = [3]int{10, 20, 30}  //method1. [3] is number of elements in the array
grades := [3]int{10, 20, 30} //method2: [3] in number of elements in array, which is 3
grades := [...]int{10, 20, 30} //method: when you use the 3 dots we need not specify the number of elements in array

---
// array intialization
package main
import "fmt"
func main() {
  var fruits [2]string = [2]string{"apples", "oranges"}
  fmt.Println(fruits)

  marks := [3]int{10, 20, 30}
  fmt.Println(marks)

  names := [...]string{"Rachel", "Phoebe", "Monica"}
  fmt.Println(names)
}
/* Output:
[apples oranges]
[10 20 30]
[Rachel Phoebe Monica]
*/

---
// length of array
package main
import "fmt"
func main() {
  var fruits [2]string = [2]string{"apples", "orange"}
  fmt.Println(len(fruits))
}
// Output: 2

---
// array indexes: first element is zero and then proceeding

/*
grades: [90 86 76 42 85]
index: 0 1 2 3 4
*/
 grades[1] 

---
package main
import "fmt"
func main() {
  var fruits [5]int = [5]string{"apples", "oranges", "grapes", "mango", "papaya"}
  fmt.Println(fruits[2])
}
// output: grapes

---
package main
import "fmt"
func main() {
  var grades [5]int = [5]int{90, 80, 70, 80, 97}
  fmt.Println(grades)
  grades[1] = 100
  fmt.Println(grades)
  
}
/* Output:
[90 80 70 80 97]
[90 100 70 80 97]
*/

---
// looping through an array
for i := 0; i < len(grades); i++ {
  fmt.Println(grades[i])
}

---
// you can also loop through an array with the "range" keyword
for index, element := range grades {
  fmt.Println(index, "=>", element)
}

///
package main
import "fmt"
func main() {
  var grades [5]int = [5]int{90, 80, 70, 80, 97}

  for index, element := range grades {
    fmt.Println(index, "=>", element)
  }
}
/* output:
0 => 90
1 => 80
2 => 70
3 => 80
4 => 97
*/

---
// multidimesional arrays
package main
import "fmt"
func main() {
  arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}} // three elements and two elements within an element
  fmt.Println(arr[2][1])
}
// output: 64


