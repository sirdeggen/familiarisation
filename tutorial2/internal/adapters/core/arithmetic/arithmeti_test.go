package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("excpected: %v, got: %v", 2, err)
	}
	require.Equal(t, answer, int32(2))
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Subtraction(123, 23)
	if err != nil {
		t.Fatalf("excpected: %v, got: %v", 100, err)
	}
	require.Equal(t, answer, int32(100))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(4, 2)
	if err != nil {
		t.Fatalf("excpected: %v, got: %v", 8, err)
	}
	require.Equal(t, answer, int32(8))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(16, 4)
	if err != nil {
		t.Fatalf("excpected: %v, got: %v", 4, err)
	}
	require.Equal(t, answer, int32(4))
}
