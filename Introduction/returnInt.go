package main
import "fmt"
func hello()int{
	return 1;
}
func helloString()string{
return "vinay";
}
func helloFloat() float64{

	return 454.54}
func main(){
	var integer1 int=hello();
	fmt.Print(integer1);
	var string11 string=helloString();
	fmt.Print(string11);
	flt64:=helloFloat();
	fmt.Println("Float 64 variable",flt64);
	fmt.Println(helloFloat());
}
