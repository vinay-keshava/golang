package main
import "fmt"
type Student struct{
	name string
	age int
	usn int
}
func main(){
	vinay:=Student{
		name:"Vinay",
		age: 20,
		usn:41,
	}
	//vinay1:=Student{}
	vinay1:=Student{name:"VK 351"}
	vinay1.age=20
	vinay1.usn=2342
	fmt.Println(vinay1)
	fmt.Println(vinay)
}
