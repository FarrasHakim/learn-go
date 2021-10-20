package main

import(
	"fmt"
)

func main(){
	fmt.Println("Hello world")

	var x int
	x = 1

	y := 1

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)

	var arr [5]int
	fmt.Println("arr", arr)
	
	arr[0] = 1
	arr[1] = 2
	
	fmt.Println("arr", arr)
	
	arr = [5]int{5,4,3,2,1}
	fmt.Println("arr", arr)
	
	slice := []int{1,2,3,4,5}
	fmt.Println("slice", slice)
	
	slice = append(slice, 6)
	fmt.Println("slice", slice)
	
	testmap := make(map[string]string)
	fmt.Println("map", testmap)
	testmap["key1"] = "value1"
	testmap["key2"] = "value2"
	fmt.Println("map", testmap)
	
	delete(testmap, "key2")
	fmt.Println("map", testmap)

	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	for index, value := range arr {
		fmt.Print("index: ", index)
		fmt.Println("value: ", value)
	}
}