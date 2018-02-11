package main

import (
	"math/rand"
)

func generateWaitingTime() []float64 {
	dr := DefRange{wtMin, wtMax}
	//randSrc ensures that the number that is generated is random
	randSrc := rand.New(rand.NewSource(now()))
	WTF := dr.ConsRandom(randSrc, TotalProfileRead)
	wrs := make([]float64, TotalProfileRead)
	//fmt.Printf("Waiting Factor(Random) 					= %v\n", WTF)
	twm := totalValueMatrix(WTF)
	for i, f := range WTF {
		sw := float64(float64(f) / float64(twm)) //10 is the maximum number of jobs
		w := 1 - sw
		wrs[i] = w * w3
	}
	return wrs
}
