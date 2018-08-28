package structs

import "fmt"

// GetArray esta funcion imprime dos array
func GetArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 4}
	arr1[0] = "First"
	arr1[1] = "Last"
	fmt.Println(arr1)
	fmt.Println(arr2)
}

// GetSlice Imprime un slice
func GetSlice() {
	var slice1 []string
	slice1 = append(slice1, "Selena Gomez", "Demi Lovato", "Christina Aguilera")
	fmt.Println(slice1)
}
