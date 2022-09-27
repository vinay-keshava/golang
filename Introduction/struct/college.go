package main
import . "fmt"
type College struct{
	name string
	colCode int
	capacity int
}
func main(){
	aiet:=College{
		name:"Alvas Institute of Engineering and Technology",
		colCode: 20,
		capacity: 300,
	}
	Println(aiet.colCode)
	Println(aiet)
}

