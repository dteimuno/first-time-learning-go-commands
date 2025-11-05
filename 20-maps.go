// declaring and initalizing a map
var <map_name> map[key_data_type]<value_data_type>
var my_map map[string]int
---
//declaring and initializing a map:
<map_name> := map[key_data_type]<value_data_type>{key-value-pair}
// example of declaring and initializing a map:
codes := map[string]string{"en": "English", "fr": "French"}
---
// declaring and initializing a map with a full golang script:
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French"}
  fmt.Println(codes)
}
//output: map[en:English fr:French]

---
// declaring and initializing a function a map with the make() function
<map_name> :=
make(map[key_data_type]<value_data_type>,<initial_capacity>) //initial_capacity is an optional argument
---

// declaring and initializing a function a map with the make() function in full golang script:
package main
import "fmt"
func main() {
  codes := make(map[string]int)
  fmt.Println(codes)
}
//output: map[]

---
//length of a map
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  fmt.Println(len(codes))
}
//output: 3

---
//accessing items in a map:
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  fmt.Println(codes["en"])
  fmt.Println(codes["fr"])
  fmt.Println(codes["hi"])
  
}
/* Output:
English
French
Hindi
*/

---
// getting the value associate with a map key
package main
import "fmt"
func main() {
  codes := map[string]int{"en": 1, "fr": 2, "hi": 3}
  value, found := codes["en"]
  fmt.Println(found, value)
  value, found := codes["hh"]
  fmt.Println(found, value)

  
}
/*Output:
true 1
false 0
*/

---
//adding key value pair to map
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  codes["it"] = "Italian"
  fmt.Println(codes)

}

---
//update key-value pair value:
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  codes["en"] = "English Language"
  fmt.Println(codes)

}
//output: map[en:English Language fr:French hi:Hindi]
---
//delete key-value pair:
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  delete(codes, "en") //delete map key pair with key that is "en"
  fmt.Println(codes)
  

}
//output: map[fr:French hi:Hindi] subtraction of first key-pair
---
//iterate over a map:
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  for key, value := range codes {
    fmt.Println(key, "=>", value)
  }  

}
/* Output:
en => English
fr => French
hi => Hindi
*/

---
// Truncate a map. Method1: iterating over map and deleting keys one by one
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  for key, value := range codes {
    delete(codes, key)
  }  

}
// Output: map[]

---
// Truncate a map. Method2: re-initialize map with emptiness
package main
import "fmt"
func main() {
  codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
  codes = make(map[string]string) //re-initialized map with empty map
  }  

}
// Output: map[]




