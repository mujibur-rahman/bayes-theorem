package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	fmt.Printf("PF(Cf * Weight(w1)) 				= %.2f\n", WoC)
	fmt.Println("---------------------------Competence factor---------------------------------------")
	genDis := generateDistance()
	WoD := distanceFactor(genDis)
	var distSelect []float64
	for k, v := range WoD {
		fmt.Printf("%v			= %.2f\n", k, v)
		if k == "Area[10]" { //default selecting nearest distance
			distSelect = make([]float64, len(v))
			for i, r := range v {
				distSelect[i] = r
			}
		}
	}
	fmt.Println("---------------------------Distance factor-----------------------------------------")
	WoT := generateWaitingTime()
	fmt.Printf("WTF(wtf * Weight(w3)) 				= %.2f\n", WoT)
	fmt.Println("---------------------------Waiting time factor--------------------------------------")
	fmt.Printf("DIST 				= %.2f\n", distSelect)
	maxCF := getMax(WoC)
	maxDF := getMax(distSelect)
	maxWF := getMax(WoT)
	fmt.Printf("maxCF 				= %.2f\n", maxCF)
	fmt.Printf("maxDF 				= %.2f\n", maxDF)
	fmt.Printf("maxWF 				= %.2f\n", maxWF)
	fmt.Printf("Result 				= %.2f\n", maxCF+maxDF+maxWF)

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

func generateWaitingTime() []float64 {
	dr := DefRange{wtMin, wtMax}
	//randSrc ensures that the number that is generated is random
	randSrc := rand.New(rand.NewSource(now()))
	WTF := dr.ConsRandom(randSrc, wtMax)
	pr := make([]float64, wtMax)
	fmt.Printf("Waiting time Factor 				= %v\n", WTF)
	for i, f := range WTF {
		p := float64(float64(f) / 10.0) //10 is the maximum number of jobs
		f := 1 - p
		pr[i] = f * w3
	}
	return pr
}

func distanceFactor(gd []int) map[string][]float64 {
	dfcMap := make(map[string][]float64)
	j := 0
	for i := 10; i <= dMax; i += 10 {
		//area should be clustered by value where distance might increase by 10
		mk := getMaxValueKey(gd, j+1, i)
		//fmt.Printf("MK -> %d | j== %v\n", mk, j)
		acd := gd[j:mk]
		sumACD := sumOfSlice(acd)
		//fmt.Printf("%v   == %v\n\n", acd, sumACD)
		dfc := make([]float64, len(acd))
		for k, v := range acd {
			//	fmt.Printf("V = %v\n", v)
			f := float64(float64(v) / sumACD)
			//	fmt.Printf("F = %v\n", f)
			mf := (1 - f)
			//fmt.Printf("MF = %v | key = %v\n", mf, k)
			dfc[k] = mf * w2
			dfcMap[fmt.Sprintf("Area[%d]", i)] = dfc
		}
		j = i
	}
	return dfcMap
}
func getMaxValueKey(ds []int, minV, maxV int) int {
	maxKey := 0
	for i, v := range ds {
		if v >= minV && v <= maxV {
			maxKey = i
		}
	}
	return maxKey
}

func sumOfSlice(as []int) float64 {
	sum := 0
	for _, v := range as {
		sum += v
	}
	return float64(sum)
}

//generateDistance() should give the clustered area output
func generateDistance() []int {
	r := DefRange{dMin, dMax}
	randSrc := rand.New(rand.NewSource(now()))
	DF := r.ConsRandom(randSrc, dMax) //maximum distance is used
	dr := make([]int, dMax)
	fmt.Printf("Random Distance Factor(df) 			= %v\n", DF)
	//Sorting the array from asc to desc
	sort.Ints(DF)
	fmt.Printf("Random Distance Factor(sorted) 			= %v\n", DF)
	for i, d := range DF {
		switch {
		case d <= 10:
			dr[i] = d
		case d > 10 && d <= 20:
			dr[i] = d
		case d > 20 && d <= 30:
			dr[i] = d
		case d > 30 && d <= 40:
			dr[i] = d
		case d > 40 && d <= 50:
			dr[i] = d
		default:
		}
	}
	return dr
}

//generateComp() should provide the random number between 1-5
func generateComp() []float64 {
	dr := DefRange{cMin, cMax}
	//randSrc ensures that the number that is generated is random
	randSrc := rand.New(rand.NewSource(now()))
	CF := dr.ConsRandom(randSrc, TotalProfileRead)
	pr := make([]float64, TotalProfileRead)
	fmt.Printf("Competence Factor 				= %v\n", CF)
	for i, f := range CF {
		p := float64(float64(f) / 5.0)
		pr[i] = p * w1
	}
	return pr
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
