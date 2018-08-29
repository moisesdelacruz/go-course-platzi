package main

import (
	"fmt"
	"strings"
	// "github.com/moisesdelacruz/gocourse/flow"
	// "github.com/moisesdelacruz/gocourse/name"
	// "github.com/moisesdelacruz/gocourse/numbers"
	// "github.com/moisesdelacruz/gocourse/structs"
	"github.com/moisesdelacruz/gocourse/maps"
)

const helloWorld string = "Hola %s, bienvenido\n"

func main() {
	// firstName := name.GetName()
	// fmt.Printf(helloWorld, firstName)
	//
	// // get ints
	// a, b, c := numbers.GetVariables()
	// fmt.Println(a, b, c)
	//
	// // sum ints numbers
	// number := numbers.Sum(45, 55)
	// fmt.Println(number)
	//
	// // get floats
	// f32, f64 := numbers.GetFloat()
	// fmt.Println(f32, f64)
	//
	// // string utf-8
	// stringUTF8 := name.GetUnicode()
	// fmt.Println("Cadena UTF-8: ", stringUTF8)
	//
	// // string index ASCII
	// fmt.Println("hello"[0])
	// // string index string
	// fmt.Println(string("hello"[0]))
	// // len
	// fmt.Println(len("Hello"))
	//
	// // Array
	// structs.GetArray()
	// // Slice
	// structs.GetSlice()
	//
	// // if test
	// flow.IfTest()
	//
	// // for Test
	// flow.ForTest()
	//
	// // strings2
	// strings2()
	//
	// // switch
	// flow.SwitchTest()

	// maps
	var name string
	fmt.Print("Name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Age: ", maps.GetMap(name))
}

func strings2() {
	text := "Hello world, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, ", "))
}
