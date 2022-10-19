package calculadora

import(
	"testing"
	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestRestar(t *testing.T) {
	// Arrange
	num1, num2 := 20, 14
	esperado := 6

	// Act
	resultado := Restar(num1, num2)

	// Assert
	assert.Equal(t, esperado, resultado, "Deben ser iguales")
}