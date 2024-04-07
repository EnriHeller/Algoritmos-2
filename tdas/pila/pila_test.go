package pila_test

import (
	"fmt"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Al ver tope de pila vacia no devuelve un panic")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Al desapilar pila vacia no devuelve un panic")
}

func TestApilar(t *testing.T) {
	{
		pila := TDAPila.CrearPilaDinamica[int]()
		elem1 := 1
		elem2 := 5
		elem3 := 39
		pila.Apilar(elem1)
		pila.Apilar(elem2)
		pila.Apilar(elem3)
		require.Equal(t, elem3, pila.Desapilar())
		require.Equal(t, elem2, pila.Desapilar())
		require.Equal(t, elem1, pila.Desapilar())
	}
	{
		pila := TDAPila.CrearPilaDinamica[string]()
		elem1 := "hola"
		elem2 := "¿como estas?"
		elem3 := "chau"
		pila.Apilar(elem1)
		pila.Apilar(elem2)
		pila.Apilar(elem3)
		require.Equal(t, elem3, pila.Desapilar())
		require.Equal(t, elem2, pila.Desapilar())
		require.Equal(t, elem1, pila.Desapilar())
	}
	{
		pila := TDAPila.CrearPilaDinamica[bool]()
		elem1 := true
		elem2 := false
		pila.Apilar(elem1)
		pila.Apilar(elem2)
		require.Equal(t, elem2, pila.Desapilar())
		require.Equal(t, elem1, pila.Desapilar())
	}
}
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := range 10000 {
		pila.Apilar(i)
	}
	tope := 9999

	for range 1000 {
		require.Equal(t, tope, pila.VerTope(),
			fmt.Sprintf("Tope esperado: %d. Tope recibido %d", tope, pila.VerTope()))

		desapilado := pila.Desapilar()
		require.Equal(t, tope, desapilado,
			fmt.Sprintf("Elemento desapilado esperado: %d. Elemento recibido %d", tope, desapilado))
		tope--
	}
}
func TestDesapilada(t *testing.T) {
	{
		pila := TDAPila.CrearPilaDinamica[string]()
		pila.Apilar("hola")
		pila.Apilar("chau")
		pila.Desapilar()
		pila.Desapilar()

		require.True(t, pila.EstaVacia())
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() },
			"Al ver el tope de pila desapilada por completo no se devuelve un panic")
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() },
			"Al desapilar pila desapilada por completo no se devuelve un panic")
	}
	{
		pila := TDAPila.CrearPilaDinamica[int]()
		pila.Apilar(12)
		pila.Apilar(36)
		pila.Desapilar()
		pila.Desapilar()

		require.True(t, pila.EstaVacia())
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() },
			"Al ver el tope de pila desapilada por completo no se devuelve un panic")
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() },
			"Al desapilar pila desapilada por completo no se devuelve un panic")
	}
	{
		pila := TDAPila.CrearPilaDinamica[*int]()
		var num1 int = 2
		var num2 int = 19

		pila.Apilar(&num1)
		pila.Apilar(&num2)
		pila.Desapilar()
		pila.Desapilar()

		require.True(t, pila.EstaVacia())
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() },
			"Al ver el tope de pila desapilada por completo no se devuelve un panic")
		require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() },
			"Al desapilar pila desapilada por completo no se devuelve un panic")
	}
}
