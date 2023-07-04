package main

type Object struct {
	Weight float64
	Value  float64
}

func Knapsack(objects []Object, capacity float64) (float64, []int) {
	// looseVars := make([]float64, len(objects))
	return 0, nil
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
