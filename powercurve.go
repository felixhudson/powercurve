package main

import (
	"fmt"
	"github.com/felixhudson/peakdection"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data := read_tcx("input.tcx")
	result := calculate(data)
	json_data := Power_json(result)
	OutputHtml("index.html", json_data)
	printCSV(result)
	peaks, _ := peakdection.Findpeaks([]int{1, 5, 1, 5, 1, 5, 1, 9, 5, 1})
	fmt.Printf("peaks = %+v\n", peaks)
}

type power struct {
	Time  int     `json:"x"`
	Power float64 `json:"y"`
}

func read_tcx(name string) []power {
	byteinput, err := ioutil.ReadFile(name)
	//input := "<Trackpoint><Time>2017-08-06T07:22:21.000Z</Time><DistanceMeters>4.0</DistanceMeters><Cadence>21</Cadence><Extensions><ns2:TPX><ns2:Speed>1.337171</ns2:Speed><ns2:Watts>37</ns2:Watts></ns2:TPX></Extensions></Trackpoint>"

	input := string(byteinput)
	var start, end int
	var result = make([]power, 0)
	var val float64
	count := 0
	lines := strings.Split(input, "<Trackpoint>")
	for i := range lines {
		count++
		if len(lines[i]) == 0 {
			//fmt.Println("Read an empty line")
		} else {
			start = strings.Index(lines[i], "<ns2:Watts>") + 11
			end = strings.Index(lines[i], "</ns2:Watts>")
			if end-1 > 0 {
				val, err = strconv.ParseFloat(lines[i][start:end], 64)
				if err == nil {
					result = append(result, power{count, val})
				}
			}
		}

	}
	return result

}

func xmlgrep(data string, searchElement string) []string {
	//output any element that starts and ends with the search element

	//input := "<Trackpoint><Time>2017-08-06T07:22:21.000Z</Time><DistanceMeters>4.0</DistanceMeters><Cadence>21</Cadence><Extensions><ns2:TPX><ns2:Speed>1.337171</ns2:Speed><ns2:Watts>37</ns2:Watts></ns2:TPX></Extensions></Trackpoint>"
	// we will use a basic io.Reader and read 100chars at a time.

	return []string{"temp"}
}

func sum(data []power) (float64, int) {
	total := 0.0
	count := 0
	for _, v := range data {
		total += v.Power
		count++
	}
	return total, count

}
func calculate(data []power) []power {
	result := make([]power, 0)
	start := 0
	end := len(data) - 1
	count := end
	sum, count := sum(data)
	var avg float64
	for start != end {
		avg = float64(sum) / float64(count)
		//fmt.Printf("Best over %d is %f\n",count, avg)
		//fmt.Printf("start %d and %d\n", start,end)
		//fmt.Println(count,sum)
		result = append(result, power{count, avg})
		if data[start].Power >= data[end].Power {
			sum = sum - data[end].Power
			// move the end
			end--
		} else {
			sum = sum - data[start].Power
			start++
		}
		count--
	}
	return result
}

func printCSV(data []power) {
	fmt.Println("Sec,Power(watts)")
	for _, v := range data {
		fmt.Println(v.Time, ",", v.Power)
	}
}
