package main
import . "fmt"
func main(){
	var a int
	a=2343
	var b *int
	b=&a
	Println(b)
	Println(*b)

	//Declaration and initialization of pointers can be done in single line
	var bb *int=&a
	Println("Single line pointer declaration and initialization",*bb)
}



/*default value of a variable is always nil

nil means zero value of a variable

*/
