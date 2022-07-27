//Go supports constants of character, string, boolean and numeric values

package main

import (
	"fmt"
	"math"
)

//Declares a constant value
const s string = "constant"

func main() {
	fmt.Println(s)

	//Constant can appear anywhere a var statement can
	const n = 5000000000000

	//Constant expressions perform arithmetic with
	// arbitrary percision
	const d = 3e20 / n
	fmt.Println(d)

	//Numeric constant has no type until it's given one,
	//such as by an explicit conversion
	fmt.Println(int64(d))

	//Number can be given a type by using it in a
	//context that requires one, such as a variable assignment or function call.
	//Here math.Sin expects a float64
	fmt.Println(math.Sin(n))
}
