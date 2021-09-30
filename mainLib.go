package main

import (
	"external-libraries/numberLib"
	"fmt"
)

func main() {
	fmt.Println("##### Practicing module import")
	n1 := numberLib.Number{12}
	fmt.Println(n1)
	n2 := numberLib.Number{8}
	fmt.Println(n2)
	fmt.Println(n1.Add(n2))
	fmt.Println(n1.Remove(n2))
	fmt.Println(n1.WeirdResult(n2))
	fmt.Println(n1.Odd())
	fmt.Println(n1.Even())
}
