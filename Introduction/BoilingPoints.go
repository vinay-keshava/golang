package main
import "fmt"
var abc=20
func hell()int {
fmt.Print("Hello Shit this is from hell function (method call)");
return 1;
}
func main(){
	const a=10
	var b=20
	fmt.Print(a);
	b=20;
	fmt.Print(b);
	aa:=hell()
	fmt.Println(aa);
}
