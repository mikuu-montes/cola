package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const tamañoVolumen = 10000

// Pre condicion: La cola debe haber sido creada previamente con "CrearColaEnlazada()"
// Post condicion: Hace las pruebas que comprueban que una cola vacia se comporte como tal.
func pruebasColaVacia[T any](t *testing.T, cola TDACola.Cola[T]) {
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

func TestColaVacia(t *testing.T) {
	nuevaCola := TDACola.CrearColaEnlazada[int]()
	pruebasColaVacia(t, nuevaCola)
}

func TestUnSoloElemento(t *testing.T) {
	colaBool := TDACola.CrearColaEnlazada[bool]()

	pruebasColaVacia(t, colaBool)

	colaBool.Encolar(true)

	require.False(t, colaBool.EstaVacia())
	require.Equal(t, true, colaBool.VerPrimero())
	require.Equal(t, true, colaBool.Desencolar())

	pruebasColaVacia(t, colaBool)
}

func TestEsFifo(t *testing.T) {
	colaString := TDACola.CrearColaEnlazada[string]()

	pruebasColaVacia(t, colaString)

	elementosString := []string{"Que", "Tal", "¿Como", "Se", "Encuentra?"}

	for _, elemento := range elementosString {
		colaString.Encolar(elemento)
	}

	require.False(t, colaString.EstaVacia())

	for i := 0; i < len(elementosString); i++ {
		require.Equal(t, elementosString[i], colaString.VerPrimero())
		require.Equal(t, elementosString[i], colaString.Desencolar())
	}

	pruebasColaVacia(t, colaString)
}

// Pre condicion: -
// Post Condicion: Rellena el array de forma ascendente (0, 1, 2, ...)
func rellenarArray(array []int) {
	for i := range array {
		array[i] = i
	}
}

func TestVolumenFifo(t *testing.T) {
	colaInts := TDACola.CrearColaEnlazada[int]()

	pruebasColaVacia(t, colaInts)

	arrayInts := make([]int, tamañoVolumen)
	rellenarArray(arrayInts)

	for _, elemento := range arrayInts {
		colaInts.Encolar(elemento)
	}

	require.False(t, colaInts.EstaVacia())

	for i := 0; i < len(arrayInts); i++ {
		require.Equal(t, arrayInts[i], colaInts.VerPrimero())
		require.Equal(t, arrayInts[i], colaInts.Desencolar())
	}

	pruebasColaVacia(t, colaInts)
}

func TestIntercalado(t *testing.T) {
	colaFloat := TDACola.CrearColaEnlazada[float64]()

	pruebasColaVacia(t, colaFloat)

	colaFloat.Encolar(1.0)
	require.Equal(t, 1.0, colaFloat.VerPrimero())

	colaFloat.Encolar(2.0)
	require.Equal(t, 1.0, colaFloat.VerPrimero())

	require.Equal(t, 1.0, colaFloat.Desencolar())
	require.Equal(t, 2.0, colaFloat.VerPrimero())

	colaFloat.Encolar(3.0)
	require.Equal(t, 2.0, colaFloat.VerPrimero())

	require.Equal(t, 2.0, colaFloat.Desencolar())
	require.Equal(t, 3.0, colaFloat.VerPrimero())
	require.Equal(t, 3.0, colaFloat.Desencolar())

	pruebasColaVacia(t, colaFloat)
}
