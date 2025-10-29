{ //outer block
  {
    // inner block
  }
}
// in golang inner blocks can access variables in outer blocks but outer blocks cannot access variables in inner blocks

// A WRONG example:
package main
import "fmt"
func main() {
  city := "London"
  {
    country := "UK"
    fmt.Println(country)
    fmt.Println(city)
  }
  fmt.Println(country) //this will cause error because outer blocks cannot access variables in inner blocks
  fmt.Println(city)
}

// THE RIGHT ERROR PROOF WAY:
package main
import "fmt"
func main() {
  city := "London"
  {
    country := "UK"
    fmt.Println(country)
    fmt.Println(city)
  }
  fmt.Println(city)
}

// Local vs Global Variables:
//local variable example:
package main
import "fmt"
func main() {
  name := "Lisa"
  fmt.Println(name)
}

//Global variables
package main
import "fmt"
var name string = "Lisa"

func main() {
  fmt.Println(name)
}
