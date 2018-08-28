package numbers

// GetVariables retorna tres numeros enteros
func GetVariables() (int, int32, int64) {
	return 1, 2147000000, 345856454353275634
}

// GetFloat retorna dos numeros de punto flotante
func GetFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

// Sum suma dos numeros enteros y devuelve el resultado
func Sum(a int, b int) int {
	return a + b
}
