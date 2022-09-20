package main
import "fmt"
func log(message string){
	fmt.Print("Hello from log() function")
}
func add(a int,b int) int{
	return a+b
}
func addd(a int, b int) (int,bool){
	return a+b,true
}
func main(){
	log("Hello ")
	add(1,3)
	fmt.Println(addd(4,5))
}

