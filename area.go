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
	return DF
}

func distFactorByCluster(grd []int, competences, weightTimes []float64) map[string][]float64 {
	dfbcData := make(map[string][]float64)
	N := len(grd)
	//fmt.Printf("Distance Factor						= %d\n", grd)
	dcw := make([]float64, N)
	dtemp := make([]float64, N)
	distBookKeep := make([]float64, N)
	totalDist := 0
	for _, d := range grd {
		totalDist += d
	}
	for k, v := range grd {
		SD := getDTCalculation(v, totalDist)
		s := 1 - SD
		distBookKeep[k] = s * w2
	}
	fmt.Printf("Distance						= %.3f\n", distBookKeep)
	for k := range grd {
		dcw[k] = distBookKeep[k] + competences[k] + weightTimes[k]
	}
	fmt.Printf("\n\nSummation of all factors				= %.3f\n", dcw)
	PT := 0.00
	for _, perFactorVal := range dcw {
		PT += perFactorVal
	}
	fmt.Printf("\nTotal of individual sum factors(PT)			= %.3f\n\n", PT)
	workerTotalFactor := 0.00
	for k, w := range dcw {
		dtemp[k] = w / PT
		fmt.Printf("P(E(%d))  						= %.3f\n", k, dtemp[k])
		workerTotalFactor += dtemp[k]
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("Total(WT)						= %.2f\n\n", workerTotalFactor)
	clusteredData(grd, dtemp)
	return dfbcData
}

func clusteredData(dists []int, dcw []float64) {
	dfbcData := make(map[int][]float64)
	dfData10 := make([]float64, 0)
	dfData20 := make([]float64, 0)
	dfData30 := make([]float64, 0)
	dfData40 := make([]float64, 0)
	dfData50 := make([]float64, 0)
	fmt.Printf("Random Distance Factor(df)				= %v\n", dists)
	for k, dst := range dists {
		switch {
		case dst >= 1 && dst <= 10:
			dfData10 = append(dfData10, dcw[k])
			dfbcData[10] = dfData10
		case dst > 10 && dst <= 20:
			dfData20 = append(dfData20, dcw[k])
			dfbcData[20] = dfData20
		case dst > 20 && dst <= 30:
			dfData30 = append(dfData30, dcw[k])
			dfbcData[30] = dfData30
		case dst > 30 && dst <= 40:
			dfData40 = append(dfData40, dcw[k])
			dfbcData[40] = dfData40
		case dst > 40 && dst <= 50:
			dfData50 = append(dfData50, dcw[k])
			dfbcData[50] = dfData50
		default:
		}
	}
	fmt.Printf("\n\n------------------------------------------------------cluster output-------------------------------------------\n\n")
	for k, dd := range dfbcData {
		sort.Float64s(dd)
		fmt.Printf("Area[`%d`]						= %.3f\n", k, dd)
	}
}

func getDTCalculation(dst, N int) float64 {
	var PRatio float64
	dstFloat := float64(dst) / float64(N)
	/*
		//TODO(Mujib): Keep it for future use
		switch {
		case dst >= 1 && dst <= 10:
			//fmt.Printf("DIST		=>%v, 10\n", dst)
			PRatio = dstFloat / 10.00
		case dst > 10 && dst <= 20:
			//fmt.Printf("DIST		=>%v, 20\n", dst)
			PRatio = dstFloat / 20.00
		case dst > 20 && dst <= 30:
			//fmt.Printf("DIST		=>%v, 30\n", dst)
			PRatio = dstFloat / 30.00
		case dst > 30 && dst <= 40:
			//fmt.Printf("DIST		=>%v, 40\n", dst)
			PRatio = dstFloat / 40.00
		case dst > 40 && dst <= 50:
			//fmt.Printf("DIST		=>%v, 50\n", dst)
			PRatio = dstFloat / 50.00
		default:
		}*/
	//fmt.Printf("NP * PRatio		=>%.3f\n", NP*PRatio)
	//return NP * PRatio
	PRatio = dstFloat / 50.00
	return PRatio
}
