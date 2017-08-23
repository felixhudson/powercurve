package main

import (
	"math"
	"testing"
)

func TestMain(t *testing.T) {
	//main()
}

func TestSum(t *testing.T) {
	input := []int{2, 4, 6, 8, 9, 9, 7, 5, 3, 1, 0}
	testdata := makedata(input)
	total, count := sum(testdata)
	if total != 54 {
		t.Fatal("cant add up")
	}
	if count != 11 {
		t.Fatal("cant count")
	}
}

func makedata(data []int) []power {
	var result = make([]power, 0)
	var foo power
	for k, v := range data {
		foo = power{k, float64(v)}
		result = append(result, foo)
	}
	return result

}
func comparefloat(a float64, b float64) bool {
	if math.Abs(a-b) < 0.00001 {
		return true
	}
	return false
}

func TestAlgorithm(t *testing.T) {
	input := []int{2, 4, 6, 8, 9, 9, 7, 5, 3, 1, 0}
	testdata := makedata(input)
	result := calculate(testdata)
	// expect to return the peak power output
	expect := 9.0
	value := result[len(result)-1].power
	if !comparefloat(expect, value) {
		t.Fatal("expecting", expect, "as peak power got", value)
	}

}
func Test_short(t *testing.T) {
	input := []int{2, 4}
	testdata := makedata(input)
	result := calculate(testdata)
	// expect to return the peak power output
	expect := 3.0
	value := result[len(result)-1].power
	if !comparefloat(expect, value) {
		t.Fatal("expecting", expect, " as peak power got", value)
	}
}
func Test_scanner(t *testing.T) {
	read_tcx("input.tcx")
}
func Test_tcx(t *testing.T) {
	data := read_tcx("input.tcx")
	result := calculate(data)
	if len(result) == 0 {
		t.Fatal("didnt read any lines ")
	}

}
