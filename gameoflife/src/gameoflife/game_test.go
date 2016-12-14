package main

import (
	"testing"
	"reflect"
)

type testField struct {
	f field
	tick int
}

func TestStart(t *testing.T) {

	tests := []testField{
		{
		f: field([][]string{
		[]string{"-","-","-"},
		[]string{"X","X","X"},
		[]string{"-","-","-"},
		},),
		tick: 3},
		{
		f: field([][]string{
		[]string{"-","-","-"},
		[]string{"X","X","X"},
		[]string{"-","-","-"},
		},),
		tick: 2},
		{
		f: field([][]string{
		[]string{"-","-","-"},
		[]string{"-","X","X"},
		[]string{"-","X","X"},
		},),
		tick: 3},
		{ 
		f:field([][]string{
		[]string{"-","-","-","-"},
		[]string{"-","X","X","X"},
		[]string{"X","X","X","-"},
		[]string{"-","-","-","-"},
		},),
		tick: 3},
		
	}
	
	results := []field{
		field([][]string{
		[]string{"-","X","-"},
		[]string{"-","X","-"},
		[]string{"-","X","-"},
		},),
		field([][]string{
		[]string{"-","-","-"},
		[]string{"X","X","X"},
		[]string{"-","-","-"},
		},),
		field([][]string{
		[]string{"-","-","-"},
		[]string{"-","X","X"},
		[]string{"-","X","X"},
		},),
		field([][]string{
		[]string{"-","-","X","-"},
		[]string{"X","-","-","X"},
		[]string{"X","-","-","X"},
		[]string{"-","X","-","-"},
		},),
	}
	
	for i , test := range tests {
		test.f.start(test.tick)
			if ! reflect.DeepEqual(test.f, results[i]) {
				t.Fatalf("Expected %s, but got %s", test.f, results[i])
			}
	}
}