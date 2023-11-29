package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Variables
	lifes := 3
	var letra string
	win := false

	// Inicializar el generador de números aleatorios
	rand.Seed(time.Now().UnixNano())
	// Defino un array con 5 palabras
	palabras := [5]string{
		"Lucas",
		"Luciano",
		"Lorenzo",
		"Marti",
		"Matias",
	}

	// Elegir una palabra aleatoria
	indiceAleatorio := rand.Intn(len(palabras))
	palabra := palabras[indiceAleatorio]

	// Inicializar el array de letras adivinadas
	letrasAdivinadas := make([]bool, len(palabra))

	// Juego principal
	for lifes > 0 && !win {
		// Mostrar las letras adivinadas y ocultar las no adivinadas
		for i, char := range palabra {
			if letrasAdivinadas[i] {
				fmt.Printf("%c ", char)
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()

		// Pedir al usuario que ingrese una letra
		fmt.Println("Ingresa una letra:")
		fmt.Scanf("%s", &letra)

		// Verificar si la letra está en la palabra
		letraCorrecta := false
		for i, char := range palabra {
			if string(char) == letra {
				letrasAdivinadas[i] = true
				letraCorrecta = true
			}
		}

		// Descontar vidas solo si la letra no está en la palabra
		if !letraCorrecta {
			lifes--
			fmt.Println("Letra incorrecta. Te quedan", lifes, "vidas.")
		}

		// Verificar si se han adivinado todas las letras
		win = true
		for _, adivinada := range letrasAdivinadas {
			if !adivinada {
				win = false
				break
			}
		}
	}

	// Mostrar el resultado final
	if win {
		fmt.Println("¡Felicidades! Has adivinado la palabra:", palabra)
	} else {
		fmt.Println("Perdiste. La palabra era:", palabra)
	}
}
