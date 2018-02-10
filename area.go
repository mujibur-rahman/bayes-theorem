package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//generateDistance() should give the clustered area output
func generateDistance() []int {
	r := DefRange{dMin, dMax}
	randSrc := rand.New(rand.NewSource(now()))
	DF := r.ConsRandom(randSrc, TotalProfileRead) //maximum distance is used
	dr := make([]int, TotalProfileRead)
	fmt.Printf("Random Distance Factor(df) 			= %v\n", DF)
	//Sorting the array from asc to desc
	sort.Ints(DF)
	fmt.Printf("Random Distance Factor(sorted) 			= %v\n", DF)
	//Clustering distance by every 10KM
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

//generate distFactorByCluster where grd <- represents the get Random Distance from random function
func distFactorByCluster(grd []int) map[string][]float64 {
	dfbcData := make(map[string][]float64)
	j := 0
	for i := 10; i <= dMax; i += 10 {
		//area should be clustered by value where distance might increase by 10
		//getMaxValueKey returns max value in each cluster
		mk := getMaxValueKey(grd, j+1, i)
		//fmt.Printf("MK -> %d | 	j = %v to %v\n", mk, j+1, i)
		//acd > Area cluster by distance
		//meaning, say area1 have 5 members found, so the acd will be index between lowest to highest
		//value of the area1 cluster
		//fmt.Printf("j=%d | mk=%d\n", j, mk)
		acd := grd[j : mk+1]   //here j = 0 because of array index starts from 0
		N := float64(len(acd)) //sumOfSlice(acd)
		//fmt.Printf("Area[%d] =>{\n", i)
		//fmt.Printf("	   %v = Total(%v)\n\t}\n", acd, N)
		//totalDurPerCluster represents D(T) for a1 cluster
		totalDurPerCluster := 0.00
		for _, v := range acd {
			/*
				a1[2,4,6,8,10] = max value 10 into cluster 1
				a2[12,14,16,18,20] = max value 20 into cluster 2
				...
				and so on,
				Cluster(a1) => D(T) = [NP(1/N) * PRatio(2/10)] + [NP(1/N) * PRatio(4/10)] + ....
				Cluster(a2) => D(T) = [NP(1/N) * PRatio(12/20)] + [NP(1/N) * PRatio(14/20)] + ....
				.....
				....
			*/
			NP := float64(1 / N)
			PRatio := float64(v) / float64(i) //here i = 10, for calculation i become float
			f := NP * PRatio
			totalDurPerCluster += f
		}
		//Area10's total duration found in totalDurPerCluster
		//Now calculate single persons duration factor and saved to map
		cdfp := make([]float64, len(acd))
		for k, v := range acd {
			/*
							[NP(1/N) * PRatio(2/10)]
				SP => 1 - 	-------------------------
								D(T)
			*/
			NP := float64(1 / N)
			PRatio := float64(v) / float64(i)
			f := NP * PRatio
			df := f / totalDurPerCluster
			cdfp[k] = (1 - df) * w2 //w2 is the weight factor
		}
		dfbcData[fmt.Sprintf("Area[%d]", i)] = cdfp
		//j = i
		j = mk + 1
	}
	return dfbcData
}
func getMaxValueKey(ds []int, minV, maxV int) int {
	maxKey := 0
	for i, v := range ds {
		if v >= minV && v <= maxV {
			//fmt.Printf("v= %v\n", ds[minV:maxV])
			maxKey = i
		}
	}
	//fmt.Printf("MaxKey=%d\n", maxKey)
	return maxKey
}

func sumOfSlice(as []int) float64 {
	sum := 0
	for _, v := range as {
		sum += v
	}
	return float64(sum)
}
