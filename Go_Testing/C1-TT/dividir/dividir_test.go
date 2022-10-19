package dividir

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestDividir(t *testing.T) {
	// Arrange
	num, den := 20, 10
	esperado := 2

	// Act
	resultado, err := Dividir(num, den)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado, "Deben ser iguales")
}