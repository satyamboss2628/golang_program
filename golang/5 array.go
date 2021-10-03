package main


import("fmt")

func main(){

	var amount[3]int = [3]int{4345,545,345}
	arr:=[...]int{34,34324,32}

	// copy := arr
	// othercopy :=arr[:]

	// pointing to the array using pointer
	// pointer :=&arr // if we change the pointer value the arr value will also change


	var mutli [3][3]int = [3][3]int{
		[3]int{1,0,0},
		[3]int{0,1,0},
		[3]int{0,0,1},
	}




	fmt.Printf("amount %v \n",amount)
	fmt.Printf("arr %v \n",arr)

	fmt.Printf("length of arr %v \n",len(arr))

	// arr[0] = 564



}




















