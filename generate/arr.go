package generate

import "math/rand"

func ArrInts(maxValue, length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}

func ArrInt32s(maxValue int32, length int) []int32 {
	arr := make([]int32, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Int31n(maxValue)
	}
	return arr
}

func ArrInt64s(maxValue int64, length int) []int64 {
	arr := make([]int64, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Int63n(maxValue)
	}
	return arr
}
