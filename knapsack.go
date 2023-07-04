package main

import "math"

type Object struct {
	Weight float64
	Value  float64
}

func Knapsack(objects []Object, capacity float64) (float64, []int) {
	mapping := func(key string) []float64 {
		var result []float64
		switch key {
		case "Weight":
			for i := range objects {
				result = append(result, objects[i].Weight)
			}
		case "Value":
			for i := range objects {
				result = append(result, objects[i].Value)
			}
		}
		return result
	}
	toIndexes := func(binaryList []float64) []int {
		var result []int
		for i := range binaryList {
			if binaryList[i] == 1.0 {
				result = append(result, i)
			}
		}
		return result
	}
	if len(objects) == 0 || len(objects) > 63 || capacity < 0 {
		return 0, []int{}
	}
	looseVars := make([]float64, len(objects))
	NextBinaryList(looseVars)
	n := int(math.Pow(2, float64(len(objects)))) - 1
	max := 0.0
	var indexes []int
	for i := 0; i < n; i++ {
		if DotProduct(looseVars, mapping("Weight")) <= capacity {
			values := DotProduct(looseVars, mapping("Value"))
			if values > max {
				max = values
				indexes = toIndexes(looseVars)
			}
		}
		NextBinaryList(looseVars)
	}
	return max, indexes
}

func NextBinaryList(vars []float64) {
	i := len(vars) - 1
	for i >= 0 && vars[i] == 1.0 {
		i--
	}
	if i == -1 {
		for j := i + 1; j < len(vars); j++ {
			vars[j] = 0.0
		}
		return
	}
	vars[i] = 1.0
	for j := i + 1; j < len(vars); j++ {
		vars[j] = 0.0
	}
}

func DotProduct(a []float64, b []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}
