package main
import "fmt"
type Student struct{
	name string
	age int
	usn int
}
func super(s Student ){
	//s=Student{name:"Vinay",age:3444,usn:2343}
	s.age=12312
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
	vin:=Student{"hello",2342,43}
	super(vin)
	fmt.Println(vin.age)
}
