package random

import "math/rand"

func RandomNumberSelectOne(seed int64, randomNum int) int {
	randSeed := rand.New(rand.NewSource(seed))
	res := randSeed.Intn(randomNum)
	return res
}
