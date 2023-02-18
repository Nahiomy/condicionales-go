package main

import "fmt"

func main() {
	//Condicional if - else 

	//Determinar si es mayor de edad
	fmt.Println("Bienvenido, por favor ingresa tu edad")
	var age int
	fmt.Scanf("%d", &age)

	if age >= 18 {
		fmt.Println("Felicidades, eres mayor de edad!!!")
	} else {
		fmt.Println("Lo sentimos, aún no eres mayor de edad 😣")
	}


	//Determinar si un numero es par o impar
	fmt.Println("Por favor, ingrese un número entero")
	var numero_ingresado int
	fmt.Scanf("%d", &numero_ingresado)

	if (numero_ingresado % 2) == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Ess impar")
	}

	//Deeterminar si un número es mayor o menor a 5
	fmt.Println("Por favor, ingrese un número entero")
	fmt.Scanf("%d", &numero_ingresado)

	if numero_ingresado > 5 {
		fmt.Println(numero_ingresado, "es mayor que 5")
	} else if numero_ingresado ==5 { 
		fmt.Println(numero_ingresado, "es igual que 5")
	} else {
		fmt.Println(numero_ingresado, "es menor que 5")
	}
}