package main
import "fmt"
func main(){
	x:=30
	ptr:=&x
	fmt.Println(ptr);
	fmt.Println(x);
	*ptr=2 //changing the value present in ptr pointer
	fmt.Println(x);
}

