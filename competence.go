package main

import (
	"fmt"
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
	fmt.Printf("Worker's profile(Random)				= %v\n", CF)
	for i, f := range CF {
		p := float64(float64(f) / 5.0)
		pr[i] = p
	}
	fmt.Printf("Competence Ratio(gen(single worker)/5.0)		= %.2f\n", pr)
	pt := totalProfileMatrix(pr)
	fmt.Printf("Total competence ratio(Sum)				= %.2f\n", pt)
	//pcm => single person choose by matrix and weight factor
	pcm := make([]float64, TotalProfileRead)
	for i, m := range pr {
		N := float64(len(pr))
		profileRatio := float64(1 / N)
		prMatrix := float64(profileRatio * m)
		spp := float64(prMatrix / pt)
		pcm[i] = spp * w1
	}
	return pcm
}

func totalProfileMatrix(tpm []float64) float64 {
	pt := 0.0
	//pt is the summation of (i)th person's matrix
	//pt += [1/total * person's matrix %]
	N := float64(len(tpm))
	NR := float64(1 / N)
	for _, m := range tpm {
		spt := float64(NR * m)
		pt += spt
	}

	return pt
}
