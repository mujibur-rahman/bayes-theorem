package main

import (
	"math/rand"
)

//generateComp() should provide the random number between 1-5
func generateComp() []float64 {
	/* 	Bayes theorem
	P(T)	=	(1 / N * Ratio(30%)) + ...nTH
	PE(1)	= (1 / N * Ratio(30%)) / P(T)
	*/
	dr := DefRange{cMin, cMax}
	//randSrc ensures that the number that is generated is random
	randSrc := rand.New(rand.NewSource(now()))
	CF := dr.ConsRandom(randSrc, TotalProfileRead)
	pr := make([]float64, TotalProfileRead)
	//fmt.Printf("Worker's profile(Random)				= %v\n", CF)
	tcm := totalValueMatrix(CF)
	for i, f := range CF {
		p := float64(float64(f) / float64(tcm)) //5 replaced by all factors
		pr[i] = p * w1
	}
	return pr
}

func totalValueMatrix(tpm []int) int {
	var pt int
	for _, m := range tpm {
		pt += m
	}
	return pt
}
