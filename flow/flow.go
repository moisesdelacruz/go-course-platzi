package flow

import "fmt"

// IfTest Verifica si un numero es por o impar
func IfTest() {
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

// SwitchTest solicita un numero del 1 al 10 y lo evalua
func SwitchTest() {
	var number int
	number = 0
	fmt.Print("[SWITCH] Ingresa un numero del 1 al 10: ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("El numero es 1")
	default:
		fmt.Println("El numero no es 1")
	}

	switch {
	case number%2 == 0:
		fmt.Printf("EL numero %d es par\n", number)
	default:
		fmt.Printf("EL numero %d es impar\n", number)
	}
}

// ForTest ...
func ForTest() {
	// traditional for
	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}

	// only conditional
	c := 100
	for c > 0 {
		c -= 10
		fmt.Printf("[FOR] solo con una condicion de %d > 0 \n", c)
	}

	// for braak
	num := 1000
	for {
		num--
		if num == 0 {
			fmt.Println("Termina el for 'infinito' con 'num': ", num)
			break
		}
	}
}
