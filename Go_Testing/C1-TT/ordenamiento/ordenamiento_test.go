package ordenamiento

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

func TestOrdenamiento(t *testing.T) {
	// Arrange
	slice := []int{3,5,1,2,8}
	esperado := []int{1,2,3,5,8}

	// Act
	resultado := Ordenamiento(slice)

	// Assert
	assert.Equal(t, esperado, resultado, "Deben ser iguales los slices")
}