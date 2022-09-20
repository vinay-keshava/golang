package main
import(
	"fmt"
)
func log(message string)string{
	return "Hello from Log:)"+message
}

func returnArrayFunc() []int {
	a:=[]int{1,2,3}
	
	return a
}
func main(){
	log("Hello ")
	var b[3] int 
	//b[]:=returnArrayFunc()
	fmt.Println(returnArrayFunc())



	
	//one line array declaration and initialization
//		varArr:=[5]int{1,2,3}

//	b[0]=10
	fmt.Println(b[0])
}

