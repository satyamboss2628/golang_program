package main



import ("fmt")


// const (
// 	user string = "Admin"
// 	yoyoy string = "Adminfdfd"
// )

const (
	i = iota
	j = iota
	k = iota
)

const (
	a = iota+1
	b
	_
	d
)

func main(){
	// const i int = 7 // the complier does not give error if we do not use const value in file 
	// const f float32 = 67

	// fmt.Println("Hello world")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Println("\n\n")

	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c)
	fmt.Println(d)



}