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
	TotalProfileRead = 100
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
	fmt.Printf("Workers(CR * Weight(w1)) 				= %.3f\n", WoC)
	WoT := generateWaitingTime()
	fmt.Printf("WTF(WR * Weight(w3)) 					= %.3f\n", WoT)
	randDistance := generateDistance()
	distFactorByCluster(randDistance, WoC, WoT)
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
