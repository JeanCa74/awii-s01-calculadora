package main

import "fmt"

func main() {

	fmt.Println("\n====CALCULADORA CIENTÍFICA V1.0====")
	for {
		var a, b float64
		var operacion string

		fmt.Print("Ingrese el primer valor :")
		fmt.Scan(&a)
		fmt.Print("Ingrese el segundo valor (Pon un 0 si vas a usar Factorial):")
		fmt.Scan(&b)
		fmt.Print("Que operacion desea realizar (+, -, *, /, ^, !) o 's' para salir :")
		fmt.Scan(&operacion)

		if operacion == "s" {
			fmt.Println("Saliendo de la calculadora ¡Hasta pronto!")
			break
		}

		switch operacion {
		case "+":

			suma := a + b

			fmt.Printf("El resultado es : %.2f", suma)

		case "-":

			resta := a - b

			fmt.Printf("El resultado es : %.2f", resta)

		case "*":

			multiplicacion := a * b

			fmt.Printf("El resultado es : %.2f", multiplicacion)

		case "/":

			if b == 0 {
				fmt.Println("NO SE PUEDE DIVIDIR PARA 0")
			} else {
				division := float64(a) / float64(b)
				fmt.Printf("El resultado es : %.2f", division)
			}

		case "^":
			resultado := 1.0
			exponente := int(b)

			for i := 0; i < exponente; i++ {
				resultado = resultado * a
			}
			fmt.Printf("El resultado es: %2.f\n", resultado)

		case "!":
			resultado := 1
			numero := int(a)
			i := 1
			for i <= numero {
				resultado = resultado * i
				i++
			}
			fmt.Printf("Resultado: %d! = %d\n", numero, resultado)

		default:
			fmt.Println("!Error¡ OPERACION NO VALIDA")
		}

	}
}
