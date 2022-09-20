package main
import "fmt"
func main(){
	var name string
	fmt.Scan(&name)
	//fmt.Scan takes characters until space is occured 
	fmt.Print(name)

	var age int
	fmt.Scan(&age)
	fmt.Println(age)
}
