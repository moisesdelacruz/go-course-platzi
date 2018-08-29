package maps

// GetMap devuelve un map con llave tipo string y valores tipo entero.
func GetMap(name string) int {
	// mapTest := make(map[string]int)
	mapTest := map[string]int{
		"Moises":   21,
		"Cesar":    23,
		"Salvador": 23,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 100
	// delete
	delete(mapTest, "llave1")
	delete(mapTest, "llave2")

	// return
	return mapTest[name]
}
