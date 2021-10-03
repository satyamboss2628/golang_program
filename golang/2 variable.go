package main


import ("fmt"

	"strconv")



// if the variable is in small letter Variable is  allowed within the package itself
// if the variable is in capital letter  exported to the variable of the other package also

// var I = 675
var i = 675

func main(){
	var i int = 12
	a :=34
	// var b = 45

	// var str string = "Hello wold"
	// var str string = string(i) // this is not valid
	var str string = strconv.Itoa(i) 
	var bar bool = true

	// fmt.Println("hello")
	fmt.Println(i)
	fmt.Printf("%v and %T\n",i,i)
	fmt.Println(a)
	fmt.Printf("%v and %T\n",str,str)

}