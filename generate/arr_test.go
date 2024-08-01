package generate

import (
	"slices"
	"testing"
)

func TestArrInts(t *testing.T) {
	arr := ArrInts(30, 10)
	arr2 := ArrInts(30, 10)

	t.Log(arr)
	t.Log(arr2)
	t.Log(slices.Equal(arr, arr2) == false)
}

func TestArrInt32s(t *testing.T) {
	arr := ArrInt32s(30, 10)
	arr2 := ArrInt32s(30, 10)

	t.Log(arr)
	t.Log(arr2)
	t.Log(slices.Equal(arr, arr2) == false)
}

func TestArrInt64s(t *testing.T) {
	arr := ArrInt64s(30, 10)
	arr2 := ArrInt64s(30, 10)

	t.Log(arr)
	t.Log(arr2)
	t.Log(slices.Equal(arr, arr2) == false)
}
