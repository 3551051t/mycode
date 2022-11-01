<<<<<<< HEAD
/* Alta3 Research | RZFeeser
   Maps - Using empty interfaces as a "wildcard" to create
   composite types */

package main

import "fmt"


func main() {

    // instead of declaring a type, we have supplied the "empty interface"
    // interface{} be in the place of the key, value, or both when making the map
    futurama := make(map[string]interface{})

    // now we can create a map of mixed types
    futurama["name"] = "Turanga Leela"  // string: string
    futurama["age"] = 30                // string: int
    futurama["height"] = 182.5          // string: float

    // display the "mixed" results
    fmt.Printf("%+v\n", futurama)
}

=======
package main

/*
CHALLENGE 01 - Map several programming languages to their file extension, for example, the key Python should map to the value .py. The key Golang should map to .go, the key Java should map to .java, the key Ansible should map to .yml, and the key C++ should map to .cpp. Once the map is constructed, display it on the screen. Then remove the language C++, and add Julia and the key .jl. Display the new map after the change.
*/

import (
	"fmt"
)

func main() {
	programmingLangExt := map[string]string{
		"Golang":  ".go",
		"Java":    ".java",
		"Ansible": ".yml",
		"C++":     ".cpp",
	}
	fmt.Println(programmingLangExt)
	fmt.Println("Delete C++")
	delete(programmingLangExt, "C++")
	programmingLangExt["Julia"] = ".jl"
	fmt.Println(programmingLangExt)
}
>>>>>>> e7d7012dda3cb2d93800c9bad517eab9c07dd35d
