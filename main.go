package main

import "fmt"

const helloWorld string = "Hola %s, bienvenido\n"

func main() {
	// name := getName()
	// fmt.Printf(helloWorld, name)

	// get ints
	a, b, c := getVariables()
	fmt.Println(a, b, c)

	// sum ints numbers
	number := sum(45, 55)
	fmt.Println(number)

	// get floats
	f32, f64 := getFloat()
	fmt.Println(f32, f64)

	// string utf-8
	stringUTF8 := getUnicode()
	fmt.Println("Cadena UTF-8: ", stringUTF8)

	// string index ASCII
	fmt.Println("hello"[0])
	// string index string
	fmt.Println(string("hello"[0]))
	// len
	fmt.Println(len("Hello"))

	// Array
	getArray()
	// Slice
	getSlice()

	// if test
	ifTest()
}

func getName() string {
	var name string
	name = "Sin Nombre"
	fmt.Print("Ingresa tu Nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, int32, int64) {
	return 1, 2147000000, 345856454353275634
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func sum(a int, b int) int {
	return a + b
}

func getUnicode() string {
	return "ハローワールド"
}

func getArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 4}
	arr1[0] = "First"
	arr1[1] = "Last"
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "Selena Gomez", "Demi Lovato", "Christina Aguilera")
	fmt.Println(slice1)
}

func ifTest() {
	var number int
	number = 0
	fmt.Print("Ingresa un numero del 1 al 10: ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Printf("El numero %d es par\n", number)
	} else {
		fmt.Printf("El numero %d es impar\n", number)
	}

	if number := 2; number == 2 {
		fmt.Println("Entro...")
	} else {
		fmt.Println("No entro...")
	}
}
