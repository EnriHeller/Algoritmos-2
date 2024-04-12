package cola_test

import (
	"fmt"
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Al ver tope de cola vacia no devuelve un panic")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Al desencolar cola vacia no devuelve un panic")
}

func TestEncolar(t *testing.T) {
	{
		cola := TDACola.CrearColaEnlazada[int]()
		elem1 := 1
		elem2 := 5
		elem3 := 39
		cola.Encolar(elem1)
		cola.Encolar(elem2)
		cola.Encolar(elem3)
		require.Equal(t, elem1, cola.Desencolar())
		require.Equal(t, elem2, cola.Desencolar())
		require.Equal(t, elem3, cola.Desencolar())
	}
	{
		cola := TDACola.CrearColaEnlazada[string]()
		elem1 := "hola"
		elem2 := "¿como estas?"
		elem3 := "chau"
		cola.Encolar(elem1)
		cola.Encolar(elem2)
		cola.Encolar(elem3)
		require.Equal(t, elem1, cola.Desencolar())
		require.Equal(t, elem2, cola.Desencolar())
		require.Equal(t, elem3, cola.Desencolar())
	}
	{
		cola := TDACola.CrearColaEnlazada[bool]()
		elem1 := true
		elem2 := false
		cola.Encolar(elem1)
		cola.Encolar(elem2)
		require.Equal(t, elem1, cola.Desencolar())
		require.Equal(t, elem2, cola.Desencolar())
	}
}
func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	const volumen int = 10000


	for i := 0; i <= volumen; i++ {
		cola.Encolar(i)
	}

	for primero := 0; primero > volumen; primero++ {
		require.Equal(t, primero, cola.VerPrimero(),
			fmt.Sprintf("Tope esperado: %d. Tope recibido %d", primero, cola.VerPrimero()))

		desencolado := cola.Desencolar()
		require.Equal(t, primero, desencolado,
			fmt.Sprintf("Elemento desencolado esperado: %d. Elemento recibido %d", primero, desencolado))
	}
}
func TestDesencolada(t *testing.T) {
	{
		cola := TDACola.CrearColaEnlazada[string]()
		cola.Encolar("hola")
		cola.Encolar("chau")
		cola.Desencolar()
		cola.Desencolar()

		require.True(t, cola.EstaVacia())
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() },
			"Al ver el tope de cola desencolada por completo no se devuelve un panic")
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() },
			"Al desencolar cola desencolada por completo no se devuelve un panic")
	}
	{
		cola := TDACola.CrearColaEnlazada[int]()
		var num1 int = 124
		var num2 int = 200

		cola.Encolar(num1)
		cola.Encolar(num2)
		cola.Desencolar()
		cola.Desencolar()

		require.True(t, cola.EstaVacia())
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() },
			"Al ver el tope de cola desencolada por completo no se devuelve un panic")
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() },
			"Al desencolar cola desencolada por completo no se devuelve un panic")
	}
	{
		cola := TDACola.CrearColaEnlazada[*int]()
		var num1 int = 2
		var num2 int = 19

		cola.Encolar(&num1)
		cola.Encolar(&num2)
		cola.Desencolar()
		cola.Desencolar()

		require.True(t, cola.EstaVacia())
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() },
			"Al ver el tope de cola desencolada por completo no se devuelve un panic")
		require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() },
			"Al desencolar cola desencolada por completo no se devuelve un panic")
	}
}
