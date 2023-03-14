package main

import "fmt"

func main(){
	var i = 21
	var j bool = true
	var x float64 = 123.456;

	fmt.Printf("%v\n", i) //21
	fmt.Printf("%T\n", i) //int
	fmt.Printf("%%\n") //%
	fmt.Printf("%t\n\n", j) //true

	fmt.Printf("%b \n", i) //10101
	fmt.Printf("%c\n", '\u042F') // menampilkan unicode russia : Я
	fmt.Printf("?\n") //?
	fmt.Printf("%d\n", i) //21
	fmt.Printf("%o\n", i) //25
	fmt.Printf("%x\n", 15) //f
	fmt.Printf("%X\n", 15) //F
	
	fmt.Printf("%U\n\n", 'Я')  //U+042F
	fmt.Printf("%f \n", x) //123.456000
	fmt.Printf("%E \n", x) //1.234560E+02
}