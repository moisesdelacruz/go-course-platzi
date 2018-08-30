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

// PlatziCourse estructura de curso
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

// Subscribe subscribe to course
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La Persona %s se a registrado a el curso de %s\n", name, p.Name)
}

// PlatziCareer estructura de career que toma los metodos de PlatziCourse
type PlatziCareer struct {
	PlatziCourse
}

// Subscribe subscribe to course
func (p PlatziCareer) Subscribe(name string) {
	fmt.Printf("La Persona %s se a registrado a la carrera de %s\n", name, p.Name)
}
