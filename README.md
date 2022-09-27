## Declarations 
 
var name type = expression;

either type or the =expression can omitted both cannot be ommitted;


## Import (ways of importing a package)


   import (

     	"fmt"

    	"math/rand"

  )


non standard lib packages are namespaced using a web url.
some like 

`import (

	github.com/vinay-keshava/go-code-package

)
`
the above go code tells the compiler to import the package from the specified repository


# Introduction to Golang

package main - specifies that the program is a standalone executable program.

# golang-introduction
Getting started with GO programming language.
Golang is platform independent programming languages that means that the binary can run on any operating system 
# www.golangbootcamp.com/book/

# Execution of Golang 
go run filename.go

go is Strongly Typed language which means that variable can have only single type string cannot be converted to int;

The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library. 

# Basic Data types available
bool string int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64

byte(alias for uint8) rune(alias for int32) float32 float64 complex64 complex128


# Declaration 
 var name type= expression

either the type or the expression is igmored 

# Short Variable Declaration
name:=expression  -> mainly used for local variables;
Here the type of the variable is automatically detected 

Examples are:
hello:=20;     Integer type

helloString:="String variable";

helloFloat:=23.43;

: is the declaration where = is the assignment

	var dd,ee=hello();

	fmt.Println(dd,ee);

func hello()(int,int){

	a,b:=10,20;

	a=a+b;

	return a, b;

}

# Mutability of Structures in Golang

Without Reference
`
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
`



`
With reference any changes in the object made must be reflected back to object this is done using pointers

func newRelease(a *Artist) int {

	a.Songs++

	return a.Songs

}

func main() {
	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}

	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))

	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)

}

`
