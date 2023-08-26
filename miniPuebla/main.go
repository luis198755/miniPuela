// Script en go para Raspberry Pi 4 que enciende y apaga 8 leds al mismo tiempo con un retardo de un segundo
package main

// Se importa el paquete rpi para controlar los pines GPIO
import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

// Se importa el paquete time para manejar el tiempo

func main() {
	// Se abre la conexión con los pines GPIO
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	// Se cierra la conexión al terminar el programa
	defer rpio.Close()

	// Se define un arreglo con los números de los pines donde se conectan los leds
	pinLeds := [8]uint8{2, 3, 4, 5, 6, 7, 8, 9}

	// Se define una variable para el tiempo de retardo en milisegundos
	tiempo := time.Duration(1000) * time.Millisecond

	// Se configuran los pines como salidas
	for _, pin := range pinLeds {
		rpio.Pin(pin).Output()
	}

	// Se inicia un bucle infinito
	for {
		// Se encienden todos los leds
		for _, pin := range pinLeds {
			rpio.Pin(pin).High()
		}

		println("Hola")

		// Se espera un segundo
		time.Sleep(tiempo)

		// Se apagan todos los leds
		for _, pin := range pinLeds {
			rpio.Pin(pin).Low()
		}

		println("Mundo")

		// Se espera otro segundo
		time.Sleep(tiempo)
	}
}
