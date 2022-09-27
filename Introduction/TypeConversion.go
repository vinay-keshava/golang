package main
import . "fmt"
func main(){
	var a int=22
	var f float64=float64(a)
	var integer uint=uint(f)
	Println(a,f,integer)


	//Another way of  type conversion is by shorthand notation
	i:=42
	j:=float64(i)
	k:=int(j)

	Println(i,j,k)
}
