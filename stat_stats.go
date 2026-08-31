package main

import (
	"maps"
	"math"
	"slices"
)

func calcSum(nums []float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += num
	}

	return sum
}

func calcProduct(nums []float64) float64 {
	product := 1.0

	for _, num := range nums {
		product *= num
	}

	return product
}

func calcQ1(nums []float64) float64 {
	var lowHalf []float64

	if len(nums)%2 == 0 {
		lowHalf = nums[:len(nums)/2]
	} else {
		lowHalf = nums[:len(nums)-1/2]
	}

	return calcMedian(lowHalf)
}

func calcQ3(nums []float64) float64 {
	var highHalf []float64

	if len(nums)%2 == 0 {
		highHalf = nums[len(nums)/2:]
	} else {
		highHalf = nums[(len(nums)+1)/2:]
	}

	return calcMedian(highHalf)
}

func calcArithMean(nums []float64) float64 {
	return calcSum(nums) / float64(len(nums))
}

func calcGeoMean(nums []float64) float64 {
	return math.Pow(calcProduct(nums), 1.0/float64(len(nums)))
}

func calcMedian(nums []float64) float64 {
	slices.Sort(nums)
	n := len(nums)

	if n%2 == 0 {
		return calcArithMean(nums[n/2-1 : n/2+1])
	}

	return nums[(n-1)/2]

}

func calcMode(nums []float64) []float64 {
	table := make(map[float64]int)
	var mode []float64

	for _, num := range nums {
		table[num] += 1
	}

	maxTimesAppeared := maxIterValue(maps.Values(table))

	for num := range maps.Keys(table) {
		if table[num] == maxTimesAppeared {
			mode = append(mode, num)
		}
	}

	if slices.Equal(mode, nums) {
		return nil
	}

	return mode
}

func calcIQR(nums []float64) float64 {
	return calcQ3(nums) - calcQ1(nums)
}

func calcVariance(nums []float64) float64 {
	var sum float64
	n := float64(len(nums))
	mean := calcArithMean(nums)

	coeff := 1 / (n - 1)

	for _, num := range nums {
		sum += math.Pow(num-mean, 2)
	}

	return coeff * sum
}

func calcStdDev(nums []float64) float64 {
	return math.Sqrt(calcVariance(nums))
}

func calcSkewness(nums []float64) float64 {
	var sum float64
	n := float64(len(nums))
	mean := calcArithMean(nums)
	stdDev := calcStdDev(nums)

	coeff := n / ((n - 1) * (n - 2))

	for _, num := range nums {
		sum += math.Pow((num-mean)/stdDev, 3)
	}

	return coeff * sum
}
