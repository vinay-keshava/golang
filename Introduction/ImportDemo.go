package main
import (
	"fmt"
	"os")
func main(){
	if len(os.Args)!=2 {
		os.Exit(1)}
	fmt.Println("Its over, blah blah,",os.Args[1])
}
