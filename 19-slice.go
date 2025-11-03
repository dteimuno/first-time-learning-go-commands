// slices are continuous segment of an underlying array
/*  a slice has three components: a pointer, length, and capacity
length of a slice can be obtained using len()
and the capacity of a slice can be obtained using cap()
*/

// declaring and initializing a slice:
<slice_name> := []<data_type>{values}
---
grades := []int{10, 20, 30}
---

package main
import "fmt"
func main() {
  slice := []int{10, 20, 30}
  fmt.Println(slice)
}
// output: [10 20 30]

---
// declaring and initializing a slice:
array[start_index : end_index]

---
package main
import "fmt"
func main() {
  arr := [10]int{10, 20, 30, 40, 50, 60 ,70, 80 ,90, 100}
  slice := arr[1:8]
  sub_slice := slice[0:3]
  fmt.Println(sub_slice)
}

---
// declaring and initializing a slice; another method using "make"
slice := make([]<data_type>, length, capacity)

---
package main
import "fmt"

func main() {
  slice := make([]int, 5, 8)
  fmt.Println(slice)
  fmt.Println(len(slice))
  fmt.Println(cap(slice))
}

---
package main
import "fmt"
func main() {
  arr := [10]int{10, 20, 30, 40, 50, 60 ,70, 80 ,90, 100}
  slice := arr[1:8]
  fmt.Println(cap(arr))
  fmt.Println(cap(slice))
}

---
//appending to a slice:
func append(s []T, vs ...T) []T
slice = append(slice, element-1, element-2)
---
package main
import "fmt"
func main() {
  arr := [4]int{10, 20, 30, 40}
  slice := arr[1:3]
  fmt.Println(slice)
  fmt.Println(len(slice))
  fmt.Println(cap(slice))
  
  
}

---
// appending to a slice
slice = append(slice, anotherSlice...)
---
//appending to a slice:
package main
import "fmt"
func main() {
  arr := [5]int{10, 20, 30, 40, 50}
  slice := arr[:2]
  arr_2 := [5]int{5, 15, 25, 35, 45}
}

---
// deleting from a slice
package main
import "fmt"
func main() {
  arr := [5]int{10, 20, 30, 40, 50}
  i := 2
  fmt.Println(arr)
  slice_1 := arr[:i]
  slice_2 := arr[i+1:]
  new_slice := append(slice_1, slice_2...)
  fmt.Println(new_slice)
}

//copy from a slice:
func copy(dst, src []Type) int
num := copy(dest_slice, src_slice)
