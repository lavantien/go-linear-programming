package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKnapsack(t *testing.T) {
	type param struct {
		objects  []Object
		capacity float64
	}
	type result struct {
		result float64
		index  []int
	}
	table := []struct {
		name  string
		input param
		want  result
	}{
		{
			name: "OK",
			input: param{
				objects: []Object{
					{Weight: 4, Value: 19},
					{Weight: 2, Value: 17},
					{Weight: 8, Value: 30},
					{Weight: 3, Value: 13},
					{Weight: 7, Value: 25},
					{Weight: 5, Value: 29},
					{Weight: 9, Value: 23},
					{Weight: 6, Value: 10},
				},
				capacity: 17,
			},
			want: result{
				result: 84.0,
				index:  []int{1, 3, 4, 5},
			},
		},
		{
			name: "Empty",
			input: param{
				objects:  []Object{},
				capacity: 17,
			},
			want: result{
				result: 0.0,
				index:  []int{},
			},
		},
	}
	for _, tc := range table {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gotResult, gotIndex := Knapsack(tc.input.objects, tc.input.capacity)
			if gotResult != tc.want.result {
				t.Errorf("Knapsack(%v, %v)\ngot\t%v, _,\nwant\t%v, _", tc.input.objects, tc.input.capacity, gotResult, tc.want.result)
			}
			if !cmp.Equal(gotIndex, tc.want.index) {
				t.Errorf("Knapsack(%v, %v)\ngot\t_, %v,\nwant\t_, %v", tc.input.objects, tc.input.capacity, gotIndex, tc.want.index)
			}
		})
	}
}

func TestNextBinary(t *testing.T) {
	type param struct {
		vars []float64
	}
	type result struct {
		result []float64
	}
	table := []struct {
		name  string
		input param
		want  result
	}{
		{
			name: "OK1",
			input: param{
				vars: []float64{1, 0, 1},
			},
			want: result{
				result: []float64{1, 1, 0},
			},
		},
		{
			name: "OK2",
			input: param{
				vars: []float64{1, 1, 0},
			},
			want: result{
				result: []float64{1, 1, 1},
			},
		},
		{
			name: "OK3",
			input: param{
				vars: []float64{0, 1, 1},
			},
			want: result{
				result: []float64{1, 0, 0},
			},
		},
		{
			name: "OK4",
			input: param{
				vars: []float64{1, 1, 1},
			},
			want: result{
				result: []float64{0, 0, 0},
			},
		},
	}
	for _, tc := range table {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			NextBinaryList(tc.input.vars)
			gotResult := tc.input.vars
			if !cmp.Equal(gotResult, tc.want.result) {
				t.Errorf("NextBinary(%v)\ngot\t%v,\nwant\t%v", tc.input.vars, gotResult, tc.want.result)
			}
		})
	}
}

func TestDotProduct(t *testing.T) {
	type param struct {
		a []float64
		b []float64
	}
	type result struct {
		result float64
	}
	table := []struct {
		name  string
		input param
		want  result
	}{
		{
			name: "OK1",
			input: param{
				a: []float64{1, 2, 3},
				b: []float64{4, 5, 6},
			},
			want: result{
				result: 32,
			},
		},
	}
	for _, tc := range table {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gotResult := DotProduct(tc.input.a, tc.input.b)
			if gotResult != tc.want.result {
				t.Errorf("DotProduct(%v, %v)\ngot\t%v,\nwant\t%v", tc.input.a, tc.input.b, gotResult, tc.want.result)
			}
		})
	}
}
