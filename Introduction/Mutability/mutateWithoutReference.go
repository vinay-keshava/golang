package main
import . "fmt"
type Student struct{
	name,add string
	age int
	allowances float64
	pocketMoneyLeft float64
}
func leftOutAfterPartying(s Student ) float64{

	s.pocketMoneyLeft=s.allowances-20
	return s.pocketMoneyLeft
}
func main(){
	s:=Student{name:"vinay", add:"beng", age:20, allowances:23.23, pocketMoneyLeft:123.21}	
	Println(s.pocketMoneyLeft)
	Println(leftOutAfterPartying(s))
	Println(s.pocketMoneyLeft)
	



}	
