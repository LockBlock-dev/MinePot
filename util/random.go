package util

import "math/rand"

func RandRange(min int, max int) int {
	return rand.Intn(max+1-min) + min
}
