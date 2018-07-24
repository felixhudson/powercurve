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
	value := result[len(result)-1].Power
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
	value := result[len(result)-1].Power
	if !comparefloat(expect, value) {
		t.Fatal("expecting", expect, " as peak power got", value)
	}
}
func Test_scanner(t *testing.T) {
	readTcx("input.tcx")
}
func Test_tcx(t *testing.T) {
	data := readTcx("input.tcx")
	result := calculate(data)
	if len(result) == 0 {
		t.Fatal("didnt read any lines ")
	}

}

func Test_split(t *testing.T) {
	input := "<Trackpoint><Time>2018-06-25T09:40:11.000Z</Time><DistanceMeters>0.0</DistanceMeters><Cadence>0</Cadence><Extensions><ns2:TPX><ns2:Speed>0.0</ns2:Speed><ns2:Watts>0</ns2:Watts></ns2:TPX></Extensions></Trackpoint><Trackpoint><Time>2018-06-25T09:40:12.000Z</Time><DistanceMeters>0.0</DistanceMeters><Cadence>0</Cadence><Extensions><ns2:TPX><ns2:Speed>0.0</ns2:Speed><ns2:Watts>0</ns2:Watts></ns2:TPX></Extensions></Trackpoint>"
	data := extractData([]byte(input))
	if len(data) != 2 {
		t.Fatal("didnt read any lines ")
	}
}
