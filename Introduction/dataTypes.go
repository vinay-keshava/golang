/*
		Datatypes present in Golang

1.Basic Types  					2.Composite Types
   -> Integers					   -> Collection / Aggregation / Non-reference Types
   	- Signed					-Arrays
	   int						-Structs
	   int8
	   int16				   -> Reference Types
	   int32					-Slices
	   int64					-Maps
							-Channels
	- Unsigned					-Pointers
	   uint						-Functions/Methods
	   uint8				   -> Interface
	   uint16					-special case of empty interface
	   uint32
	   uint64

     -> Floats
        -float32
	-float64
     -> Complex Number
        -complex64
	-complex128
     -> byte
     -> Rune
     -> String
     -> Boolean

*/
package main
import (
	. "fmt"
	. "math/bits"
)

func main(){
Println(bits.UnintSize)
}
