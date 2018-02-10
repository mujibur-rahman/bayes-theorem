package main

import (
	"fmt"
	"math/rand"
)

/*	<Rules>
*	w[3,4,5,7,8,1,10]
*	W(T)	=	(1 / N * (3/10)%) + (1 / N * (4/10)%) +...nTH
					(1 / N * (3/10)%)
*	WE(1)	= 1 -	------------------------
						W(T)
*	</Rules>
*/
func generateWaitingTime() []float64 {
	dr := DefRange{wtMin, wtMax}
	//randSrc ensures that the number that is generated is random
	randSrc := rand.New(rand.NewSource(now()))
	WTF := dr.ConsRandom(randSrc, wtMax)
	wrs := make([]float64, wtMax)
	fmt.Printf("Waiting Factor(Random) 				= %v\n", WTF)
	for i, f := range WTF {
		sw := float64(float64(f) / 10.0) //10 is the maximum number of jobs
		wrs[i] = sw
	}
	fmt.Printf("Waiting Factor Ratio(gen(single WT)/10.0)	= %.2f\n", wrs)
	twt := totalWeightTime(wrs)
	fmt.Printf("Total weight ratio(Sum)				= %.2f\n", twt)
	//wcm => single weight choose by matrix and weight factor(w3)
	wcm := make([]float64, wtMax)
	N := float64(len(wrs))
	weightRatio := float64(1 / N)
	for i, m := range wrs {
		wMatrix := float64(weightRatio * m)
		swr := float64(wMatrix / twt)
		wcm[i] = (1 - swr) * w3
	}
	return wcm
}

func totalWeightTime(wrs []float64) float64 {
	wt := 0.0
	//pt is the summation of (i)th weight's matrix
	//pt += [1/total * weight's matrix %]
	N := float64(len(wrs))
	NW := float64(1 / N)
	for _, m := range wrs {
		swt := float64(NW * m)
		wt += swt
	}
	return wt
}
