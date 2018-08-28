package name

import "fmt"

// GetName obtiene y retorna el nombre de la persona
func GetName() string {
	var name string
	name = "Sin Nombre"
	fmt.Print("Ingresa tu Nombre: ")
	fmt.Scanf("%s", &name)
	return name
}
