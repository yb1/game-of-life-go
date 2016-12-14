package main

import (
	"testing"
	"reflect"
)

func TestStart(t *testing.T) {
	tests := []field{
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
		[]string{"-","-","-","-"},
		[]string{"-","X","X","X"},
		[]string{"X","X","X","-"},
		[]string{"-","-","-","-"},
		},),
		
	}
	
	results := []field{
		field([][]string{
		[]string{"-","X","-"},
		[]string{"-","X","-"},
		[]string{"-","X","-"},
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
		test.start(3)
		
			if ! reflect.DeepEqual(test, results[i]) {
				t.Fatalf("Expected %s, but got %s", test, results[i])
			}
	}
}