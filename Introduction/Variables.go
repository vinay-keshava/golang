package main
import "fmt"
func main(){
/*
var name type = expression 
either type / =expression can be skipped.
*/	
var s int =20;
fmt.Println(s);
var i,j,k int;
fmt.Println(i,j,k);
i=20
fmt.Println(i,"Value after Updating ");
var o=true //skipping the type (var name type= expression)
fmt.Println(o);
//Multiple variable assignment with different type together
var aa,bb,cc= true,2323.32,'h'
fmt.Println(aa,bb,cc)

//Short Variable Declration .
/*
Short variable Declaration usually reserved for local variable
:= declaration
= assignment

*/
aaa:=true;
ii,jj:=0,1;
fmt.Println("\n--------Short Variable Declaration:------- ",aaa,ii,jj);
fmt.Println("Before Swapping ",i,j);
j,i = i,j;
fmt.Println("After Swapping ",i,j);



// Creating new variable using new 
fmt.Println("--------Creating new Variable using new keyword");
abcc:=new(int);
fmt.Println(abcc);



//multiple variable declartion 
var(
	name string="vinay"
	age int=20
	addr string="Bengaluru"
)
fmt.Println(name,age,addr)


}
