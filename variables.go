//Go variables are explecitly declared and used by compiler


package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	b, c := 1, 2
	fmt.Println(b, c)

	//Variable declaration short form
	d := true
	fmt.Println(d)

	//Variable declared without any value initialization
	// will return 0
	var e int
	fmt.Println(e)
}
