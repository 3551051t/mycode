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
