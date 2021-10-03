package main



import ("fmt")




func main(){

	var foo bool = true
	var i int = 5
	var u uint = 56
	f:=2.4e12

	var c complex64 = 1 + 2i // float 32 bit
	var comp complex64 = complex(4,5) // (real,imaginary)

	str :="hello I am a string" // int 8 ascii value

	// str :='hello I am a string' // int32 

	// coverting string to byte array

	s := []byte(s) // utf8 


	fmt.Println("%v %T\n",string(str[0]),s)
	fmt.Println("%v %T\n",real(c),real(c))
	fmt.Println("%v %T\n",imag(c),imag(c))


	a :=45 
	b :=46 

	fmt.Println("Hello world")
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)

	// binary operator

	fmt.Println(a&b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
}