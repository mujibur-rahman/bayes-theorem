package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	w1   = .45
	w2   = .35
	w3   = .20
	cMin = 1
	cMax = 5
	//TotalProfileRead represents the number of user profile to be selected
	TotalProfileRead = 10
	dMin             = 1
	dMax             = 50
	wtMin            = 1
	wtMax            = 10
)

// DefRange specification, note that min <= max
type DefRange struct {
	min, max int
}

func main() {
	WoC := generateComp()
	fmt.Printf("Workers(CR * Weight(w1)) 				= %.2f\n", WoC)
	fmt.Println("---------------------------------- Competence factor ---------------------------------------")
	WoT := generateWaitingTime()
	fmt.Printf("WTF(WR * Weight(w3)) 				= %.2f\n", WoT)
	fmt.Println("---------------------------------Waiting time factor-------------------------------------------")
	genDis := generateDistance()
	WoD := distFactorByCluster(genDis)
	fmt.Println("------------------------------------ Distance factor ------------------------------------------")
	workersByClusterArea(WoC, WoT, WoD)
	fmt.Println("-----------------------------Workers by clustered area------------------------------------")
}

func workersByClusterArea(WoC, WoT []float64, WoD map[string][]float64) {
	//distSelect := make([]float64, 10)
	for k, dst := range WoD {
		fmt.Printf("\nCluster<-`%v`->total(%v)\n\n", k, len(dst))
		fmt.Printf("WOD			= %.2f\n", dst[0:len(dst)])
		// fmt.Printf("WoC			= %.2f\n", WoC[0:10])
		// fmt.Printf("WoT			= %.2f\n", WoT[0:10])
		// for i, d := range dst {
		// 	//if i < 10 {
		// 	distSelect[i] = d + WoC[i] + WoT[i]
		// 	//}
		// }
		//fmt.Printf("Workers(probable)	= %.2f\n", distSelect)
	}
}

func getMax(gm []float64) float64 {
	max := gm[0]
	//min := gm[0]
	for _, value := range gm {
		if max < value {
			max = value
		}
		// if min > value {
		// 	min = value
		// }
	}
	return max
}

//ConsRandom provide next consecutive random value within the interval including min and max
func (dr *DefRange) ConsRandom(r *rand.Rand, N int) []int {
	//generate max/min between the given ranges like case is 1-5
	rn := make([]int, N)
	for i := 0; i < N; i++ {
		rn[i] = r.Intn(dr.max-dr.min+1) + dr.min
	}
	return rn
}

func now() int64 {
	return time.Now().UnixNano()
}
